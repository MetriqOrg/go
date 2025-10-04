package orbitr

import (
	"net/url"
	"time"

	"github.com/metriqorg/go/ingest/ledgerbackend"

	"github.com/sirupsen/logrus"
	"github.com/stellar/throttled"
)

// Config is the configuration for orbitr.  It gets populated by the
// app's main function and is provided to NewApp.
type Config struct {
	DatabaseURL        string
	RoDatabaseURL      string
	HistoryArchiveURLs []string
	Port               uint
	AdminPort          uint

	EnableCaptiveCoreIngestion  bool
	EnableIngestionFiltering    bool
	CaptiveCoreBinaryPath       string
	RemoteCaptiveCoreURL        string
	CaptiveCoreConfigPath       string
	CaptiveCoreTomlParams       ledgerbackend.CaptiveCoreTomlParams
	CaptiveCoreToml             *ledgerbackend.CaptiveCoreToml
	CaptiveCoreStoragePath      string
	CaptiveCoreReuseStoragePath bool
	CaptiveCoreConfigUseDB      bool

	GravityDatabaseURL string
	GravityURL         string

	// MaxDBConnections has a priority over all 4 values below.
	MaxDBConnections            int
	OrbitRDBMaxOpenConnections int
	OrbitRDBMaxIdleConnections int

	SSEUpdateFrequency time.Duration
	ConnectionTimeout  time.Duration
	RateQuota          *throttled.RateQuota
	FriendbotURL       *url.URL
	LogLevel           logrus.Level
	LogFile            string

	// MaxPathLength is the maximum length of the path returned by `/paths` endpoint.
	MaxPathLength uint
	// MaxAssetsPerPathRequest is the maximum number of assets considered for `/paths/strict-send` and `/paths/strict-receive`
	MaxAssetsPerPathRequest int
	// DisablePoolPathFinding configures orbitr to run path finding without including liquidity pools
	// in the path finding search.
	DisablePoolPathFinding bool
	// DisablePathFinding configures orbitr without the path finding endpoint.
	DisablePathFinding bool
	// MaxPathFindingRequests is the maximum number of path finding requests orbitr will allow
	// in a 1-second period. A value of 0 disables the limit.
	MaxPathFindingRequests uint

	NetworkPassphrase string
	SentryDSN         string
	LogglyToken       string
	LogglyTag         string
	// TLSCert is a path to a certificate file to use for orbitr's TLS config
	TLSCert string
	// TLSKey is the path to a private key file to use for orbitr's TLS config
	TLSKey string
	// Ingest toggles whether this orbitr instance should run the data ingestion subsystem.
	Ingest bool
	// CursorName is the cursor used for ingesting from gravity.
	// Setting multiple cursors in different OrbitR instances allows multiple
	// OrbitRs to ingest from the same gravity instance without cursor
	// collisions.
	CursorName string
	// HistoryRetentionCount represents the minimum number of ledgers worth of
	// history data to retain in the orbitr database. For the purposes of
	// determining a "retention duration", each ledger roughly corresponds to 10
	// seconds of real time.
	HistoryRetentionCount uint
	// StaleThreshold represents the number of ledgers a history database may be
	// out-of-date by before orbitr begins to respond with an error to history
	// requests.
	StaleThreshold uint
	// SkipCursorUpdate causes the ingestor to skip reporting the "last imported
	// ledger" state to gravity.
	SkipCursorUpdate bool
	// IngestDisableStateVerification disables state verification
	// `System.verifyState()` when set to `true`.
	IngestDisableStateVerification bool
	// IngestStateVerificationCheckpointFrequency configures how often state verification is performed.
	// If IngestStateVerificationCheckpointFrequency is set to 1 state verification is run on every checkpoint,
	// If IngestStateVerificationCheckpointFrequency is set to 2 state verification is run on every second checkpoint,
	// etc...
	IngestStateVerificationCheckpointFrequency uint
	// IngestStateVerificationTimeout configures a timeout on the state verification routine.
	// If IngestStateVerificationTimeout is set to 0 the timeout is disabled.
	IngestStateVerificationTimeout time.Duration
	// IngestEnableExtendedLogLedgerStats enables extended ledger stats in
	// logging.
	IngestEnableExtendedLogLedgerStats bool
	// ApplyMigrations will apply pending migrations to the orbitr database
	// before starting the orbitr service
	ApplyMigrations bool
	// CheckpointFrequency establishes how many ledgers exist between checkpoints
	CheckpointFrequency uint32
	// BehindCloudflare determines if OrbitR instance is behind Cloudflare. In
	// such case http.Request.RemoteAddr will be replaced with Cloudflare header.
	BehindCloudflare bool
	// BehindAWSLoadBalancer determines if OrbitR instance is behind AWS load
	// balances like ELB or ALB. In such case http.Request.RemoteAddr will be
	// replaced with the last IP in X-Forwarded-For header.
	BehindAWSLoadBalancer bool
	// RoundingSlippageFilter excludes trades from /trade_aggregations with rounding slippage >x bps
	RoundingSlippageFilter int
	// Lantah Network: 'testnet' or 'pubnet'
	Network string
	// DisableTxSub disables transaction submission functionality for OrbitR.
	DisableTxSub bool
}
