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

package dbcmd

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strconv"
	"strings"

	"github.com/gravitational/trace"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"

	"github.com/gravitational/teleport/api/constants"
	"github.com/gravitational/teleport/lib/client"
	"github.com/gravitational/teleport/lib/client/db"
	"github.com/gravitational/teleport/lib/client/db/mysql"
	"github.com/gravitational/teleport/lib/client/db/opensearch"
	"github.com/gravitational/teleport/lib/client/db/postgres"
	"github.com/gravitational/teleport/lib/defaults"
	"github.com/gravitational/teleport/lib/tlsca"
	"github.com/gravitational/teleport/lib/utils"
)

const (
	// postgresBin is the Postgres client binary name.
	postgresBin = "psql"
	// cockroachBin is the Cockroach client binary name.
	cockroachBin = "cockroach"
	// mysqlBin is the MySQL client binary name.
	mysqlBin = "mysql"
	// mariadbBin is the MariaDB client binary name.
	mariadbBin = "mariadb"
	// mongoshBin is the Mongo Shell client binary name.
	mongoshBin = "mongosh"
	// mongoBin is the Mongo client binary name.
	mongoBin = "mongo"
	// redisBin is the Redis client binary name.
	redisBin = "redis-cli"
	// mssqlBin is the SQL Server client program name.
	mssqlBin = "mssql-cli"
	// sqlcmd is the SQL Server client program name.
	sqlcmdBin = "sqlcmd"
	// snowsqlBin is the Snowflake client program name.
	snowsqlBin = "snowsql"
	// cqlshBin is the Cassandra client program name.
	cqlshBin = "cqlsh"
	// curlBin is the program name for `curl`, which is used as Elasticsearch client if other options are unavailable.
	curlBin = "curl"
	// elasticsearchSQLBin is the Elasticsearch SQL client program name.
	elasticsearchSQLBin = "elasticsearch-sql-cli"
	// openSearchCLIBin is the OpenSearch CLI client program name.
	openSearchCLIBin = "opensearch-cli"
	// openSearchSQLBin is the OpenSearch SQL client program name.
	openSearchSQLBin = "opensearchsql"
	// awsBin is the aws CLI program name.
	awsBin = "aws"
	// oracleBin is the Oracle CLI program name.
	oracleBin = "sql"
)

// Execer is an abstraction of Go's exec module, as this one doesn't specify any interfaces.
// This interface exists only to enable mocking.
type Execer interface {
	// RunCommand runs a system command.
	RunCommand(name string, arg ...string) ([]byte, error)
	// LookPath returns a full path to a binary if this one is found in system PATH,
	// error otherwise.
	LookPath(file string) (string, error)
}

// SystemExecer implements execer interface by using Go exec module.
type SystemExecer struct{}

// RunCommand is a wrapper for exec.Command(...).Output()
func (s SystemExecer) RunCommand(name string, arg ...string) ([]byte, error) {
	return exec.Command(name, arg...).Output()
}

// LookPath is a wrapper for exec.LookPath(...)
func (s SystemExecer) LookPath(file string) (string, error) {
	return exec.LookPath(file)
}

// CLICommandBuilder holds data needed to build a CLI command from args passed to NewCmdBuilder.
// Any calls to the exec package within CLICommandBuilder methods that need to be mocked should
// use the exe field rather than calling the package directly.
type CLICommandBuilder struct {
	tc          *client.TeleportClient
	rootCluster string
	profile     *client.ProfileStatus
	db          *tlsca.RouteToDatabase
	host        string
	port        int
	options     connectionCommandOpts
	uid         utils.UID
}

func NewCmdBuilder(tc *client.TeleportClient, profile *client.ProfileStatus,
	db tlsca.RouteToDatabase, rootClusterName string, opts ...ConnectCommandFunc,
) *CLICommandBuilder {
	var options connectionCommandOpts
	for _, opt := range opts {
		opt(&options)
	}

	// In TLS routing mode a local proxy is started on demand so connect to it.
	host := options.localProxyHost
	port := options.localProxyPort
	if host == "" || port == 0 {
		host, port = tc.DatabaseProxyHostPort(db)
	}

	if options.log == nil {
		options.log = logrus.NewEntry(logrus.StandardLogger())
	}

	if options.exe == nil {
		options.exe = &SystemExecer{}
	}

	return &CLICommandBuilder{
		tc:          tc,
		profile:     profile,
		db:          &db,
		host:        host,
		port:        port,
		options:     options,
		rootCluster: rootClusterName,
		uid:         utils.NewRealUID(),
	}
}

// GetConnectCommand returns a command that can connect the user directly to the given database
// using an appropriate CLI database client. It takes into account cluster configuration, binaries
// available on the system and in some cases it even connects to the database to check which exact
// version of the database the user is running.
//
// Underneath it uses exec.Command, so the resulting command will always be expanded to its absolute
// path if exec.LookPath was able to find the given binary on user's system.
//
// If CLICommandBuilder's options.tolerateMissingCLIClient is set to true, GetConnectCommand
// shouldn't return an error if it cannot locate a client binary. Check WithTolerateMissingCLIClient
// docs for more details.
func (c *CLICommandBuilder) GetConnectCommand() (*exec.Cmd, error) {
	switch c.db.Protocol {
	case defaults.ProtocolPostgres:
		return c.getPostgresCommand(), nil

	case defaults.ProtocolCockroachDB:
		return c.getCockroachCommand(), nil

	case defaults.ProtocolMySQL:
		return c.getMySQLCommand()

	case defaults.ProtocolMongoDB:
		return c.getMongoCommand(), nil

	case defaults.ProtocolRedis:
		return c.getRedisCommand(), nil

	case defaults.ProtocolSQLServer:
		return c.getSQLServerCommand(), nil

	case defaults.ProtocolSnowflake:
		return c.getSnowflakeCommand(), nil

	case defaults.ProtocolCassandra:
		return c.getCassandraCommand()

	case defaults.ProtocolElasticsearch:
		return c.getElasticsearchCommand()

	case defaults.ProtocolOpenSearch:
		return c.getOpenSearchCommand()

	case defaults.ProtocolDynamoDB:
		return c.getDynamoDBCommand()

	case defaults.ProtocolOracle:
		return c.getOracleCommand()
	}

	return nil, trace.BadParameter("unsupported database protocol: %v", c.db)
}

// CommandAlternative represents alternative command along with description.
type CommandAlternative struct {
	Description string
	Command     *exec.Cmd
}

// GetConnectCommandAlternatives returns optional connection commands for protocols that offer multiple options.
// Otherwise, it falls back to GetConnectCommand.
// The keys in the returned map are command descriptions suitable for display to the end user.
func (c *CLICommandBuilder) GetConnectCommandAlternatives() ([]CommandAlternative, error) {

	switch c.db.Protocol {
	case defaults.ProtocolElasticsearch:
		return c.getElasticsearchAlternativeCommands(), nil
	case defaults.ProtocolOpenSearch:
		return c.getOpenSearchAlternativeCommands(), nil
	}

	cmd, err := c.GetConnectCommand()
	if err != nil {
		return nil, trace.Wrap(err)
	}

	return []CommandAlternative{{Description: "default command", Command: cmd}}, nil
}

func (c *CLICommandBuilder) getPostgresCommand() *exec.Cmd {
	return exec.Command(postgresBin, c.getPostgresConnString())
}

func (c *CLICommandBuilder) getCockroachCommand() *exec.Cmd {
	// If cockroach CLI client is not available, fallback to psql.
	if _, err := c.options.exe.LookPath(cockroachBin); err != nil {
		c.options.log.Debugf("Couldn't find %q client in PATH, falling back to %q: %v.",
			cockroachBin, postgresBin, err)
		return c.getPostgresCommand()
	}
	return exec.Command(cockroachBin, "sql", "--url", c.getPostgresConnString())
}

// getPostgresConnString returns the connection string for postgres.
func (c *CLICommandBuilder) getPostgresConnString() string {
	return postgres.GetConnString(
		db.New(c.tc, *c.db, *c.profile, c.rootCluster, c.host, c.port),
		c.options.noTLS,
		c.options.printFormat,
	)
}

// getMySQLCommonCmdOpts returns common command line arguments for mysql and mariadb.
// Currently, the common options are: user, database, host, port and protocol.
func (c *CLICommandBuilder) getMySQLCommonCmdOpts() []string {
	args := make([]string, 0)
	if c.db.Username != "" {
		args = append(args, "--user", c.db.Username)
	}
	if c.db.Database != "" {
		args = append(args, "--database", c.db.Database)
	}

	if c.options.localProxyPort != 0 {
		args = append(args, "--port", strconv.Itoa(c.options.localProxyPort))
		args = append(args, "--host", c.options.localProxyHost)
		// MySQL CLI treats localhost as a special value and tries to use Unix Domain Socket for connection
		// To enforce TCP connection protocol needs to be explicitly specified.
		if c.options.localProxyHost == "localhost" {
			args = append(args, "--protocol", "TCP")
		}
	}

	return args
}

// getMariaDBArgs returns arguments unique for mysql cmd shipped by MariaDB and mariadb cmd. Common options for mysql
// between Oracle and MariaDB version are covered by getMySQLCommonCmdOpts().
func (c *CLICommandBuilder) getMariaDBArgs() []string {
	args := c.getMySQLCommonCmdOpts()

	if c.options.noTLS {
		return args
	}

	// Some options used in the MySQL options file are not compatible with the
	// "mariadb" client. Thus instead of using `--defaults-group-suffix=`,
	// specify the proxy host and port directly as parameters. When
	// localProxyPort is specified, the --port and --host flags are set by
	// getMySQLCommonCmdOpts.
	if c.options.localProxyPort == 0 {
		host, port := c.tc.MySQLProxyHostPort()
		args = append(args, "--port", strconv.Itoa(port))
		args = append(args, "--host", host)
	}

	sslCertPath := c.profile.DatabaseCertPathForCluster(c.tc.SiteName, c.db.ServiceName)

	args = append(args, []string{"--ssl-key", c.profile.KeyPath()}...)
	args = append(args, []string{"--ssl-ca", c.profile.CACertPathForCluster(c.rootCluster)}...)
	args = append(args, []string{"--ssl-cert", sslCertPath}...)

	// Flag below verifies "Common Name" check on the certificate provided by the server.
	// This option is disabled by default.
	if !c.tc.InsecureSkipVerify {
		args = append(args, "--ssl-verify-server-cert")
	}

	return args
}

// getMySQLOracleCommand returns arguments unique for mysql cmd shipped by Oracle. Common options between
// Oracle and MariaDB version are covered by getMySQLCommonCmdOpts().
func (c *CLICommandBuilder) getMySQLOracleCommand() (*exec.Cmd, error) {
	args := c.getMySQLCommonCmdOpts()

	if c.options.noTLS {
		return exec.Command(mysqlBin, args...), nil
	}

	// defaults-group-suffix must be first.
	groupSuffix := []string{fmt.Sprintf("--defaults-group-suffix=_%v-%v", c.tc.SiteName, c.db.ServiceName)}
	args = append(groupSuffix, args...)

	if runtime.GOOS == constants.WindowsOS {
		// We save configuration to ~/.my.cnf, but on Windows that file is not read,
		// see tables 4.1 and 4.2 on https://dev.mysql.com/doc/refman/8.0/en/option-files.html.
		// We instruct mysql client to use use that file with --defaults-extra-file.
		configPath, err := mysql.DefaultConfigPath(c.tc.HomePath)
		if err != nil {
			return nil, trace.Wrap(err)
		}
		extraFile := []string{fmt.Sprintf("--defaults-extra-file=%v", configPath)}
		args = append(extraFile, args...)
	}

	// override the ssl-mode from a config file is --insecure flag is provided to 'tsh db connect'.
	if c.tc.InsecureSkipVerify {
		args = append(args, fmt.Sprintf("--ssl-mode=%s", mysql.MySQLSSLModeVerifyCA))
	}

	return exec.Command(mysqlBin, args...), nil
}

// getMySQLCommand returns mariadb command if the binary is on the path. Otherwise,
// mysql command is returned. Both mysql versions (MariaDB and Oracle) are supported.
func (c *CLICommandBuilder) getMySQLCommand() (*exec.Cmd, error) {
	// Check if mariadb client is available. Prefer it over mysql client even if connecting to MySQL server.
	if c.isMariaDBBinAvailable() {
		args := c.getMariaDBArgs()
		return exec.Command(mariadbBin, args...), nil
	}

	// Check for mysql binary. In case the caller doesn't tolerate a missing CLI client, return with
	// error as mysql and mariadb are missing. There is nothing else we can do here.
	if !c.isMySQLBinAvailable() {
		if c.options.tolerateMissingCLIClient {
			return c.getMySQLOracleCommand()
		}

		return nil, trace.NotFound("neither %q nor %q CLI clients were found, please make sure an appropriate CLI client is available in $PATH", mysqlBin, mariadbBin)
	}

	// Check which flavor is installed. Otherwise, we don't know which ssl flag to use.
	// At the moment of writing mysql binary shipped by Oracle and MariaDB accept different ssl parameters and have the same name.
	mySQLMariaDBFlavor, err := c.isMySQLBinMariaDBFlavor()
	if mySQLMariaDBFlavor && err == nil {
		args := c.getMariaDBArgs()
		return exec.Command(mysqlBin, args...), nil
	}

	// Either we failed to check the flavor or binary comes from Oracle. Regardless return mysql/Oracle command.
	return c.getMySQLOracleCommand()
}

func (c *CLICommandBuilder) isBinAvailable(file string) bool {
	_, err := c.options.exe.LookPath(file)
	return err == nil
}

// isMariaDBBinAvailable returns true if "mariadb" binary is found in the system PATH.
func (c *CLICommandBuilder) isMariaDBBinAvailable() bool {
	return c.isBinAvailable(mariadbBin)
}

// isMySQLBinAvailable returns true if "mysql" binary is found in the system PATH.
func (c *CLICommandBuilder) isMySQLBinAvailable() bool {
	return c.isBinAvailable(mysqlBin)
}

// isMongoshBinAvailable returns true if "mongosh" binary is found in the system PATH.
func (c *CLICommandBuilder) isMongoshBinAvailable() bool {
	return c.isBinAvailable(mongoshBin)
}

// isElasticsearchSQLBinAvailable returns true if "elasticsearch-sql-cli" binary is found in the system PATH.
func (c *CLICommandBuilder) isElasticsearchSQLBinAvailable() bool {
	return c.isBinAvailable(elasticsearchSQLBin)
}

// isOpenSearchCLIBinAvailable returns true if "opensearch-cli" binary is found in the system PATH.
func (c *CLICommandBuilder) isOpenSearchCLIBinAvailable() bool {
	return c.isBinAvailable(openSearchCLIBin)
}

// isOpenSearchCLIBinAvailable returns true if "opensearchsql" binary is found in the system PATH.
func (c *CLICommandBuilder) isOpenSearchSQLBinAvailable() bool {
	return c.isBinAvailable(openSearchSQLBin)
}

// isMySQLBinMariaDBFlavor checks if mysql binary comes from Oracle or MariaDB.
// true is returned when binary comes from MariaDB, false when from Oracle.
func (c *CLICommandBuilder) isMySQLBinMariaDBFlavor() (bool, error) {
	// Check if mysql comes from Oracle or MariaDB
	mysqlVer, err := c.options.exe.RunCommand(mysqlBin, "--version")
	if err != nil {
		// Looks like incorrect mysql installation.
		return false, trace.Wrap(err)
	}

	// Check which flavor is installed. Otherwise, we don't know which ssl flag to use.
	// Example output:
	// Oracle:
	// mysql  Ver 8.0.27-0ubuntu0.20.04.1 for Linux on x86_64 ((Ubuntu))
	// MariaDB:
	// mysql  Ver 15.1 Distrib 10.3.32-MariaDB, for debian-linux-gnu (x86_64) using readline 5.2
	return strings.Contains(strings.ToLower(string(mysqlVer)), "mariadb"), nil
}

// isSqlcmdAvailable returns true if "sqlcmd" binary is fouind in the system
// PATH.
func (c *CLICommandBuilder) isSqlcmdAvailable() bool {
	return c.isBinAvailable(sqlcmdBin)
}

func (c *CLICommandBuilder) shouldUseMongoshBin() bool {
	// Use "mongosh" if available.
	// If not, use legacy "mongo" if available.
	// If both are not available, pick "mongosh" in print out.
	return c.isMongoshBinAvailable() || !c.isBinAvailable(mongoBin)
}

func (c *CLICommandBuilder) getMongoCommand() *exec.Cmd {
	// look for `mongosh`
	useMongosh := c.shouldUseMongoshBin()

	var args []string

	if !c.options.noTLS {
		// Starting with Mongo 4.2 there is an updated set of flags.
		// We are using them with `mongosh` as otherwise warnings will get displayed.
		type tlsFlags struct {
			tls            string
			tlsCertKeyFile string
			tlsCAFile      string
		}

		var flags tlsFlags

		if useMongosh {
			flags = tlsFlags{tls: "--tls", tlsCertKeyFile: "--tlsCertificateKeyFile", tlsCAFile: "--tlsCAFile"}
		} else {
			flags = tlsFlags{tls: "--ssl", tlsCertKeyFile: "--sslPEMKeyFile", tlsCAFile: "--sslCAFile"}
		}

		args = append(args,
			flags.tls,
			flags.tlsCertKeyFile,
			c.profile.DatabaseCertPathForCluster(c.tc.SiteName, c.db.ServiceName))

		if c.options.caPath != "" {
			// caPath is set only if mongo connects to the Teleport Proxy via ALPN SNI Local Proxy
			// and connection is terminated by proxy identity certificate.
			args = append(args, []string{flags.tlsCAFile, c.options.caPath}...)
		} else {
			// mongosh does not load system CAs by default which will cause issues if
			// the proxy presents a certificate signed by a non-recognized authority
			// which your system trusts (e.g. mkcert).
			if useMongosh {
				args = append(args, "--tlsUseSystemCA")
			}
		}
	}

	// Add the address at the end. Address contains host, port, database name,
	// and other options like server selection timeout.
	args = append(args, c.getMongoAddress())

	// use `mongosh` if available
	if useMongosh {
		return exec.Command(mongoshBin, args...)
	}

	// fall back to `mongo` if `mongosh` isn't found
	return exec.Command(mongoBin, args...)
}

func (c *CLICommandBuilder) getMongoAddress() string {
	query := make(url.Values)

	// Use the same default server selection timeout (5s) that the backend
	// engine is using. The environment variable serves as a hidden option to
	// force a different timeout for debugging purpose or extreme situations.
	serverSelectionTimeoutMS := "5000"
	if envValue := os.Getenv(envVarMongoServerSelectionTimeoutMS); envValue != "" {
		c.options.log.Infof("Using environment variable %s=%s.", envVarMongoServerSelectionTimeoutMS, envValue)
		serverSelectionTimeoutMS = envValue
	}
	query.Set("serverSelectionTimeoutMS", serverSelectionTimeoutMS)

	address := url.URL{
		Scheme:   connstring.SchemeMongoDB,
		Host:     fmt.Sprintf("%s:%d", c.host, c.port),
		RawQuery: query.Encode(),
		Path:     fmt.Sprintf("/%s", c.db.Database),
	}

	// Quote the address for printing as the address contains "?".
	if c.options.printFormat {
		return fmt.Sprintf(`"%s"`, address.String())
	}
	return address.String()
}

// getRedisCommand returns redis-cli commands used by 'tsh db connect' when connecting to a Redis instance.
func (c *CLICommandBuilder) getRedisCommand() *exec.Cmd {
	// TODO(jakub): Add "-3" when Teleport adds support for Redis RESP3 protocol.
	args := []string{
		"-h", c.host,
		"-p", strconv.Itoa(c.port),
	}

	if !c.options.noTLS {
		args = append(args,
			"--tls",
			"--key", c.profile.KeyPath(),
			"--cert", c.profile.DatabaseCertPathForCluster(c.tc.SiteName, c.db.ServiceName))

		if c.tc.InsecureSkipVerify {
			args = append(args, "--insecure")
		}

		if c.options.caPath != "" {
			args = append(args, []string{"--cacert", c.options.caPath}...)
		}

		// Set SNI when sending to remote web proxy.
		if c.options.localProxyHost == "" {
			args = append(args, []string{"--sni", c.tc.WebProxyHost()}...)
		}
	}

	// append database number if provided
	if c.db.Database != "" {
		args = append(args, []string{"-n", c.db.Database}...)
	}

	return exec.Command(redisBin, args...)
}

// getSQLServerCommand returns a command to connect to SQL Server.
// mssql-cli and sqlcmd commands have the same argument names.
func (c *CLICommandBuilder) getSQLServerCommand() *exec.Cmd {
	args := []string{
		// Host and port must be comma-separated.
		"-S", fmt.Sprintf("%v,%v", c.host, c.port),
		"-U", c.db.Username,
		// Password is required by the client but doesn't matter as we're
		// connecting to local proxy.
		"-P", c.uid.New(),
	}

	if c.db.Database != "" {
		args = append(args, "-d", c.db.Database)
	}

	if c.isSqlcmdAvailable() {
		return exec.Command(sqlcmdBin, args...)
	}

	return exec.Command(mssqlBin, args...)
}

func (c *CLICommandBuilder) getSnowflakeCommand() *exec.Cmd {
	args := []string{
		"-a", "teleport", // Account name doesn't matter as it will be overridden in the backend anyway.
		"-u", c.db.Username,
		"-h", c.host,
		"-p", strconv.Itoa(c.port),
	}

	if c.db.Database != "" {
		args = append(args, "-w", c.db.Database)
	}

	cmd := exec.Command(snowsqlBin, args...)
	cmd.Env = append(cmd.Env, fmt.Sprintf("SNOWSQL_PWD=%s", c.uid.New()))

	return cmd
}

func (c *CLICommandBuilder) getCassandraCommand() (*exec.Cmd, error) {
	args := []string{
		"-u", c.db.Username,
		c.host, strconv.Itoa(c.port),
	}
	if c.options.password != "" {
		args = append(args, []string{"-p", c.options.password}...)
	}
	return exec.Command(cqlshBin, args...), nil
}

// getElasticsearchCommand returns a command to connect to Elasticsearch. We support `elasticsearch-sql-cli`, but only in non-TLS scenario.
func (c *CLICommandBuilder) getElasticsearchCommand() (*exec.Cmd, error) {
	if c.options.noTLS {
		return exec.Command(elasticsearchSQLBin, fmt.Sprintf("http://%v:%v/", c.host, c.port)), nil
	}
	return nil, trace.BadParameter("%v interactive command is only supported in --tunnel mode.", elasticsearchSQLBin)
}

// getOpenSearchCommand returns a command to connect to OpenSearch.
func (c *CLICommandBuilder) getOpenSearchCommand() (*exec.Cmd, error) {
	if c.options.tolerateMissingCLIClient == false && c.isOpenSearchSQLBinAvailable() == false {
		return nil, trace.NotFound("%q not found, please make sure it is available in $PATH", openSearchSQLBin)
	}

	if c.options.noTLS {
		args := []string{fmt.Sprintf("http://%v:%v", c.host, c.port)}
		return exec.Command(openSearchSQLBin, args...), nil
	}

	return nil, trace.BadParameter("%v interactive command is only supported in --tunnel mode.", openSearchSQLBin)
}

func (c *CLICommandBuilder) getOpenSearchCLICommand() (*exec.Cmd, error) {
	cfg := opensearch.ConfigNoTLS(c.host, c.port)
	if !c.options.noTLS {
		cfg = opensearch.ConfigTLS(c.host, c.port, c.options.caPath, c.profile.DatabaseCertPathForCluster(c.tc.SiteName, c.db.ServiceName), c.profile.KeyPath())
	}

	baseDir := path.Join(c.profile.Dir, c.profile.Cluster, c.db.ServiceName)
	tempCfg, err := opensearch.WriteConfig(baseDir, cfg)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	args := []string{"--profile", opensearch.ProfileName, "--config", tempCfg, "curl", "get", "--path", "/"}

	return exec.Command(openSearchCLIBin, args...), nil
}

func (c *CLICommandBuilder) getDynamoDBCommand() (*exec.Cmd, error) {
	// we can't guess at what the user wants to do, so this command is for print purposes only,
	// and it only works with a local proxy tunnel.
	if !c.options.printFormat || !c.options.noTLS || c.options.localProxyHost == "" || c.options.localProxyPort == 0 {
		svc := "<db>"
		if c.db != nil && c.db.ServiceName != "" {
			svc = c.db.ServiceName
		}
		return nil, trace.BadParameter("DynamoDB requires a local proxy tunnel. Use `tsh proxy db --tunnel %v`", svc)
	}
	args := []string{
		"--endpoint", fmt.Sprintf("http://%v:%v/", c.options.localProxyHost, c.options.localProxyPort),
		"[dynamodb|dynamodbstreams|dax]",
		"<command>",
	}
	return exec.Command(awsBin, args...), nil
}

type jdbcOracleThinConnection struct {
	host     string
	port     int
	db       string
	tnsAdmin string
}

func (j *jdbcOracleThinConnection) ConnString() string {
	return fmt.Sprintf(`jdbc:oracle:thin:@tcps://%s:%d/%s?TNS_ADMIN=%s`, j.host, j.port, j.db, j.tnsAdmin)
}

func (c *CLICommandBuilder) getOracleCommand() (*exec.Cmd, error) {
	cs := jdbcOracleThinConnection{
		host:     c.host,
		port:     c.port,
		db:       c.db.Database,
		tnsAdmin: c.profile.OracleWalletDir(c.profile.Cluster, c.db.ServiceName),
	}
	// Quote the address for printing as the address contains "?".
	connString := cs.ConnString()
	if c.options.printFormat {
		connString = fmt.Sprintf(`'%s'`, connString)
	}
	args := []string{
		"-L", // dont retry
		connString,
	}
	return exec.Command(oracleBin, args...), nil
}

func (c *CLICommandBuilder) getElasticsearchAlternativeCommands() []CommandAlternative {
	var commands []CommandAlternative
	if c.isElasticsearchSQLBinAvailable() {
		if cmd, err := c.getElasticsearchCommand(); err == nil {
			commands = append(commands, CommandAlternative{Description: "interactive SQL connection", Command: cmd})
		}
	}

	var curlCommand *exec.Cmd
	if c.options.noTLS {
		curlCommand = exec.Command(curlBin, fmt.Sprintf("http://%v:%v/", c.host, c.port))
	} else {
		args := []string{
			fmt.Sprintf("https://%v:%v/", c.host, c.port),
			"--key", c.profile.KeyPath(),
			"--cert", c.profile.DatabaseCertPathForCluster(c.tc.SiteName, c.db.ServiceName),
		}

		if c.tc.InsecureSkipVerify {
			args = append(args, "--insecure")
		}

		if c.options.caPath != "" {
			args = append(args, []string{"--cacert", c.options.caPath}...)
		}

		// Force HTTP 1.1 when connecting to remote web proxy. Otherwise, HTTP2 can
		// be negotiated which breaks the engine.
		if c.options.localProxyHost == "" {
			args = append(args, "--http1.1")
		}

		curlCommand = exec.Command(curlBin, args...)
	}
	commands = append(commands, CommandAlternative{Description: "run single request with curl", Command: curlCommand})

	return commands
}

func (c *CLICommandBuilder) getOpenSearchAlternativeCommands() []CommandAlternative {
	var commands []CommandAlternative
	if c.isOpenSearchSQLBinAvailable() {
		if cmd, err := c.getOpenSearchCommand(); err == nil {
			commands = append(commands, CommandAlternative{Description: "start interactive session with opensearchsql", Command: cmd})
		}
	}

	if c.isOpenSearchCLIBinAvailable() {
		if cmd, err := c.getOpenSearchCLICommand(); err == nil {
			commands = append(commands, CommandAlternative{Description: "run request with opensearch-cli", Command: cmd})
		}
	}

	var curlCommand *exec.Cmd
	if c.options.noTLS {
		curlCommand = exec.Command(curlBin, fmt.Sprintf("http://%v:%v/", c.host, c.port))
	} else {
		args := []string{
			fmt.Sprintf("https://%v:%v/", c.host, c.port),
			"--key", c.profile.KeyPath(),
			"--cert", c.profile.DatabaseCertPathForCluster(c.tc.SiteName, c.db.ServiceName),
		}

		if c.tc.InsecureSkipVerify {
			args = append(args, "--insecure")
		}

		if c.options.caPath != "" {
			args = append(args, []string{"--cacert", c.options.caPath}...)
		}

		// Force HTTP 1.1 when connecting to remote web proxy. Otherwise, HTTP2 can
		// be negotiated which breaks the engine.
		if c.options.localProxyHost == "" {
			args = append(args, "--http1.1")
		}

		curlCommand = exec.Command(curlBin, args...)
	}
	commands = append(commands, CommandAlternative{Description: "run request with curl", Command: curlCommand})

	return commands
}

type connectionCommandOpts struct {
	localProxyPort           int
	localProxyHost           string
	caPath                   string
	noTLS                    bool
	printFormat              bool
	tolerateMissingCLIClient bool
	log                      *logrus.Entry
	exe                      Execer
	password                 string
}

// ConnectCommandFunc is a type for functions returned by the "With*" functions in this package.
// A function of type ConnectCommandFunc changes connectionCommandOpts of CLICommandBuilder based on
// the arguments passed to a "With*" function.
type ConnectCommandFunc func(*connectionCommandOpts)

// WithLocalProxy makes CLICommandBuilder pass appropriate args to the CLI database clients that
// will let them connect to a database through a local proxy.
// In most cases it means using the passed host and port as the address, but some database clients
// require additional flags in those scenarios.
func WithLocalProxy(host string, port int, caPath string) ConnectCommandFunc {
	return func(opts *connectionCommandOpts) {
		opts.localProxyPort = port
		opts.localProxyHost = host
		opts.caPath = caPath
	}
}

// WithNoTLS is the connect command option that makes the command connect
// without TLS.
//
// It is used when connecting through the local proxy that was started in
// mutual TLS mode (i.e. with a client certificate).
func WithNoTLS() ConnectCommandFunc {
	return func(opts *connectionCommandOpts) {
		opts.noTLS = true
	}
}

// WithPassword is the command option that allows to set the database password
// that will be used for database CLI.
func WithPassword(pass string) ConnectCommandFunc {
	return func(opts *connectionCommandOpts) {
		opts.password = pass
	}
}

// WithPrintFormat is the connect command option that hints the command will be
// printed instead of being executed.
//
// For example, when enabled, a quote will be used for Postgres and MongoDB
// connection strings to avoid "&" getting interpreted by the shell.
//
// WithPrintFormat is known to be used for the following situations:
// - tsh db config --format cmd <database>
// - tsh proxy db --tunnel <database>
// - Teleport Connect where the command is put into a terminal.
//
// WithPrintFormat should NOT be used when the exec.Cmd gets executed by the
// client application.
func WithPrintFormat() ConnectCommandFunc {
	return func(opts *connectionCommandOpts) {
		opts.printFormat = true
	}
}

// WithLogger is the connect command option that allows the caller to pass a logger that will be
// used by CLICommandBuilder.
func WithLogger(log *logrus.Entry) ConnectCommandFunc {
	return func(opts *connectionCommandOpts) {
		opts.log = log
	}
}

// WithTolerateMissingCLIClient is the connect command option that makes CLICommandBuilder not
// return an error in case a specific binary couldn't be found in the system. Instead, it should
// return the command with just a base version of the binary name, without an absolute path.
//
// In general CLICommandBuilder doesn't return an error in that scenario as it uses exec.Command
// underneath. However, there are some specific situations where we need to execute some
// binaries before returning the final command.
//
// The flag is mostly for scenarios where the caller doesn't care that the final command might not
// work.
func WithTolerateMissingCLIClient() ConnectCommandFunc {
	return func(opts *connectionCommandOpts) {
		opts.tolerateMissingCLIClient = true
	}
}

// WithExecer allows to provide a different Execer than the default SystemExecer. Useful in contexts
// where there's a place that wants to use dbcmd with the ability to mock out SystemExecer in tests.
func WithExecer(exe Execer) ConnectCommandFunc {
	return func(opts *connectionCommandOpts) {
		opts.exe = exe
	}
}

const (
	// envVarMongoServerSelectionTimeoutMS is the environment variable that
	// controls the server selection timeout used for MongoDB clients.
	envVarMongoServerSelectionTimeoutMS = "TELEPORT_MONGO_SERVER_SELECTION_TIMEOUT_MS"
)
