package gramr

const (
	// LiveState represents the state value returned by gramr when a
	// ledger entry is live
	LiveState = "live"

	// DeadState represents the state value returned by gramr when a
	// ledger entry is dead
	DeadState = "dead"
)

// GetLedgerEntryResponse is the response from Gramr for the getLedgerEntry endpoint
type GetLedgerEntryResponse struct {
	State  string `json:"state"`
	Entry  string `json:"entry"`
	Ledger int64  `json:"ledger"`
}
