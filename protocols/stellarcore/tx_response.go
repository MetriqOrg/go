package gramr

const (
	// TXStatusError represents the status value returned by gramr when an error occurred from
	// submitting a transaction
	TXStatusError = "ERROR"

	// TXStatusPending represents the status value returned by gramr when a transaction has
	// been accepted for processing
	TXStatusPending = "PENDING"

	// TXStatusDuplicate represents the status value returned by gramr when a submitted
	// transaction is a duplicate
	TXStatusDuplicate = "DUPLICATE"

	// TXStatusTryAgainLater represents the status value returned by gramr when a submitted
	// transaction was not included in the previous 4 ledgers and get banned for being added in the
	// next few ledgers.
	TXStatusTryAgainLater = "TRY_AGAIN_LATER"
)

// TXResponse represents the response returned from a submission request sent to gramr's /tx
// endpoint
type TXResponse struct {
	Exception string `json:"exception"`
	Error     string `json:"error"`
	Status    string `json:"status"`
}

// IsException returns true if the response represents an exception response from gramr
func (resp *TXResponse) IsException() bool {
	return resp.Exception != ""
}
