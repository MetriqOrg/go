package ingest

import (
	"testing"

	"github.com/metriqorg/go/xdr"
	"github.com/stretchr/testify/assert"
)

func TestGenesisLeaderStateReader(t *testing.T) {
	change := GenesisChange("Public Global Lantah Network ; 2023")
	assert.Equal(t, xdr.LedgerEntryTypeAccount, change.Type)
	assert.Equal(t, xdr.Uint32(1), change.Post.LastModifiedLedgerSeq)
	account := change.Post.Data.MustAccount()
	assert.Equal(t, "GB2YTHGV3PDS4OJR3B6ZBYTNNO42DKXNFH3WACMTKFYSXDYXCV3LIS4R", account.AccountId.Address())
	assert.Equal(t, xdr.SequenceNumber(0), account.SeqNum)
	assert.Equal(t, xdr.Int64(1000000000000000000), account.Balance)
	assert.Equal(t, xdr.Thresholds{1, 0, 0, 0}, account.Thresholds)
}
