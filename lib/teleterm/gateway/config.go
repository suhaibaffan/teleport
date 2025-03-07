/*
Copyright 2021 Gravitational, Inc.

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

package gateway

import (
	"context"
	"crypto/x509"
	"net"
	"runtime"

	"github.com/google/uuid"
	"github.com/gravitational/trace"
	"github.com/jonboulle/clockwork"
	"github.com/sirupsen/logrus"

	"github.com/gravitational/teleport/api/constants"
	"github.com/gravitational/teleport/lib/defaults"
	"github.com/gravitational/teleport/lib/srv/alpnproxy"
	"github.com/gravitational/teleport/lib/teleterm/api/uri"
	"github.com/gravitational/teleport/lib/tlsca"
)

// Config describes gateway configuration
type Config struct {
	// URI is the gateway URI
	URI uri.ResourceURI
	// TargetName is the remote resource name
	TargetName string
	// TargetURI is the remote resource URI
	TargetURI uri.ResourceURI
	// TargetUser is the target user name
	TargetUser string
	// TargetGroups is a list of target groups
	TargetGroups []string
	// TargetSubresourceName points at a subresource of the remote resource, for example a database
	// name on a database server. It is used only for generating the CLI command.
	TargetSubresourceName string

	// Port is the gateway port
	LocalPort string
	// LocalAddress is the local address
	LocalAddress string
	// Protocol is the gateway protocol
	Protocol string
	// CertPath specifies the path to the user certificate that the local proxy
	// uses to connect to the Teleport Proxy. The path may depend on the type
	// and the parameters of the gateway.
	CertPath string
	// KeyPath specifies the path to the private key of the cert specified in
	// the CertPath. This is usually the private key of the user profile.
	KeyPath string
	// Insecure
	Insecure bool
	// ClusterName is the Teleport cluster name.
	ClusterName string
	// Username is the username of the profile.
	Username string
	// WebProxyAddr
	WebProxyAddr string
	// Log is a component logger
	Log *logrus.Entry
	// CLICommandProvider returns a CLI command for the gateway
	CLICommandProvider CLICommandProvider
	// TCPPortAllocator creates listeners on the given ports. This interface lets us avoid occupying
	// hardcoded ports in tests.
	TCPPortAllocator TCPPortAllocator
	// Clock is used by Gateway.localProxy to check cert expiration.
	Clock clockwork.Clock
	// OnExpiredCert is called when a new downstream connection is accepted by the
	// gateway but cannot be proxied because the cert used by the gateway has expired.
	//
	// Handling of the connection is blocked until OnExpiredCert returns.
	OnExpiredCert OnExpiredCertFunc
	// TLSRoutingConnUpgradeRequired indicates that ALPN connection upgrades
	// are required for making TLS routing requests.
	TLSRoutingConnUpgradeRequired bool
	// RootClusterCACertPoolFunc is callback function to fetch Root cluster CAs
	// when ALPN connection upgrade is required.
	RootClusterCACertPoolFunc alpnproxy.GetClusterCACertPoolFunc
	// ProfileDir specifies the tsh home dir of the user profile.
	ProfileDir string
}

// OnExpiredCertFunc is the type of a function that is called when a new downstream connection is
// accepted by the gateway but cannot be proxied because the cert used by the gateway has expired.
//
// Handling of the connection is blocked until the function returns.
type OnExpiredCertFunc func(context.Context, Gateway) error

// CheckAndSetDefaults checks and sets the defaults
func (c *Config) CheckAndSetDefaults() error {
	if c.URI.String() == "" {
		c.URI = uri.NewGatewayURI(uuid.NewString())
	}

	if c.LocalAddress == "" {
		c.LocalAddress = "localhost"
		// SQL Server Management Studio won't connect to localhost:12345, so use 127.0.0.1:12345 instead.
		if runtime.GOOS == constants.WindowsOS && c.Protocol == defaults.ProtocolSQLServer {
			c.LocalAddress = "127.0.0.1"
		}
	}

	if c.LocalPort == "" {
		c.LocalPort = "0"
	}

	if c.Log == nil {
		c.Log = logrus.NewEntry(logrus.StandardLogger())
	}

	if c.TargetName == "" {
		return trace.BadParameter("missing target name")
	}

	if c.TargetURI.String() == "" {
		return trace.BadParameter("missing target URI")
	}

	if c.CLICommandProvider == nil {
		return trace.BadParameter("missing CLICommandProvider")
	}

	if c.TCPPortAllocator == nil {
		c.TCPPortAllocator = NetTCPPortAllocator{}
	}

	if c.Clock == nil {
		c.Clock = clockwork.NewRealClock()
	}

	if c.RootClusterCACertPoolFunc == nil {
		if !c.Insecure {
			return trace.BadParameter("missing RootClusterCACertPoolFunc")
		}
		c.RootClusterCACertPoolFunc = func(_ context.Context) (*x509.CertPool, error) {
			return x509.NewCertPool(), nil
		}
	}

	c.Log = c.Log.WithFields(logrus.Fields{
		"resource": c.TargetURI.String(),
		"gateway":  c.URI.String(),
	})
	return nil
}

// RouteToDatabase returns tlsca.RouteToDatabase based on the config of the gateway.
//
// The tlsca.RouteToDatabase.Database field is skipped, as it's an optional field and gateways can
// change their Config.TargetSubresourceName at any moment.
func (c *Config) RouteToDatabase() tlsca.RouteToDatabase {
	return tlsca.RouteToDatabase{
		ServiceName: c.TargetName,
		Protocol:    c.Protocol,
		Username:    c.TargetUser,
	}
}

func (c *Config) makeListener() (net.Listener, error) {
	listener, err := c.TCPPortAllocator.Listen(c.LocalAddress, c.LocalPort)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	// retrieve automatically assigned port number
	_, port, err := net.SplitHostPort(listener.Addr().String())
	if err != nil {
		return nil, trace.Wrap(err)
	}

	c.LocalPort = port
	return listener, nil
}
