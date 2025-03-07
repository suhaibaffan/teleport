/*
Copyright 2022 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package web

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gravitational/trace"
	"github.com/jonboulle/clockwork"
	"github.com/stretchr/testify/require"

	"github.com/gravitational/teleport/api/constants"
	"github.com/gravitational/teleport/api/utils/pingconn"
	"github.com/gravitational/teleport/lib/utils"
)

func TestWriteUpgradeResponse(t *testing.T) {
	var buf bytes.Buffer
	require.NoError(t, writeUpgradeResponse(&buf, "custom"))

	resp, err := http.ReadResponse(bufio.NewReader(&buf), nil)
	require.NoError(t, err)

	// Always drain/close the body.
	io.Copy(io.Discard, resp.Body)
	_ = resp.Body.Close()

	require.Equal(t, resp.StatusCode, http.StatusSwitchingProtocols)
	require.Equal(t, "custom", resp.Header.Get("Upgrade"))
}

func TestHandlerConnectionUpgrade(t *testing.T) {
	t.Parallel()

	expectedPayload := "hello@"
	expectedIP := "1.2.3.4"
	alpnHandler := func(_ context.Context, conn net.Conn) error {
		// Handles connection asynchronously to verify web handler waits until
		// connection is closed.
		go func() {
			defer conn.Close()

			clientIP, err := utils.ClientIPFromConn(conn)
			require.NoError(t, err)
			require.Equal(t, expectedIP, clientIP)

			n, err := conn.Write([]byte(expectedPayload))
			require.NoError(t, err)
			require.Equal(t, len(expectedPayload), n)
		}()
		return nil
	}

	// Cherry picked some attributes to create a Handler to test only the
	// connection upgrade portion.
	h := &Handler{
		cfg: Config{
			ALPNHandler: alpnHandler,
		},
		log:   newPackageLogger(),
		clock: clockwork.NewRealClock(),
	}

	t.Run("unsupported type", func(t *testing.T) {
		r, err := http.NewRequest("GET", "http://localhost/webapi/connectionupgrade", nil)
		require.NoError(t, err)
		r.Header.Add("Upgrade", "unsupported-protocol")

		_, err = h.connectionUpgrade(httptest.NewRecorder(), r, nil)
		require.True(t, trace.IsNotFound(err))
	})

	t.Run("upgraded to ALPN", func(t *testing.T) {
		serverConn, clientConn := net.Pipe()
		defer serverConn.Close()
		defer clientConn.Close()

		sendConnUpgradeRequest(t, h, constants.WebAPIConnUpgradeTypeALPN, serverConn, clientConn, expectedIP)

		// Verify clientConn receives data sent by Config.ALPNHandler.
		receive, err := bufio.NewReader(clientConn).ReadString(byte('@'))
		require.NoError(t, err)
		require.Equal(t, expectedPayload, receive)
	})

	t.Run("upgraded to ALPN with Ping", func(t *testing.T) {
		serverConn, clientConn := net.Pipe()
		defer serverConn.Close()
		defer clientConn.Close()

		sendConnUpgradeRequest(t, h, constants.WebAPIConnUpgradeTypeALPNPing, serverConn, clientConn, expectedIP)

		// Verify ping-wrapped clientConn receives data sent by Config.ALPNHandler.
		receive, err := bufio.NewReader(pingconn.New(clientConn)).ReadString(byte('@'))
		require.NoError(t, err)
		require.Equal(t, expectedPayload, receive)
	})
}

func sendConnUpgradeRequest(t *testing.T, h *Handler, upgradeType string, serverConn, clientConn net.Conn, xForwardedFor string) {
	t.Helper()

	r, err := http.NewRequest("GET", "http://localhost/webapi/connectionupgrade", nil)
	require.NoError(t, err)
	r.Header.Add("Upgrade", upgradeType)
	r.Header.Add("X-Forwarded-For", xForwardedFor)

	// serverConn will be hijacked.
	w := newResponseWriterHijacker(nil, serverConn)
	require.NoError(t, err)

	go func() {
		// Use XForwardedFor middleware to set IPs.
		var err error
		connUpgradeHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, err = h.connectionUpgrade(w, r, nil)
		})
		NewXForwardedForMiddleware(connUpgradeHandler).ServeHTTP(w, r)

		require.NoError(t, err)
	}()

	// Verify clientConn receives http.StatusSwitchingProtocols.
	resp, err := http.ReadResponse(bufio.NewReader(clientConn), r)
	require.NoError(t, err)

	// Always drain/close the body.
	io.Copy(io.Discard, resp.Body)
	_ = resp.Body.Close()

	require.Equal(t, upgradeType, resp.Header.Get(constants.WebAPIConnUpgradeHeader))
	require.Equal(t, constants.WebAPIConnUpgradeConnectionType, resp.Header.Get(constants.WebAPIConnUpgradeConnectionHeader))
	require.Equal(t, http.StatusSwitchingProtocols, resp.StatusCode)
}

// responseWriterHijacker is a mock http.ResponseWriter that also serves a
// net.Conn for http.Hijacker.
type responseWriterHijacker struct {
	http.ResponseWriter
	conn net.Conn
}

func newResponseWriterHijacker(w http.ResponseWriter, conn net.Conn) http.ResponseWriter {
	if w == nil {
		w = httptest.NewRecorder()
	}
	return &responseWriterHijacker{
		ResponseWriter: w,
		conn:           conn,
	}
}

func (h *responseWriterHijacker) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return h.conn, nil, nil
}
