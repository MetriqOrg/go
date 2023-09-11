package gravity

const (
	// LiveState represents the state value returned by gravity when a
	// ledger entry is live
	LiveState = "live"

	// DeadState represents the state value returned by gravity when a
	// ledger entry is dead
	DeadState = "dead"
)

// GetLedgerEntryResponse is the response from Gravity for the getLedgerEntry endpoint
type GetLedgerEntryResponse struct {
	State  string `json:"state"`
	Entry  string `json:"entry"`
	Ledger int64  `json:"ledger"`
}
