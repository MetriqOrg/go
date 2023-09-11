package internal

import (
	"github.com/lantah/go/clients/orbitrclient"
	"github.com/lantah/go/support/errors"
)

// Account implements the `txnbuild.Account` interface.
type Account struct {
	AccountID string
	Sequence  int64
}

// GetAccountID returns the Account ID.
func (a Account) GetAccountID() string {
	return a.AccountID
}

// IncrementSequenceNumber increments the internal record of the
// account's sequence number by 1.
func (a Account) IncrementSequenceNumber() (int64, error) {
	a.Sequence++
	return a.Sequence, nil
}

func (a Account) GetSequenceNumber() (int64, error) {
	return a.Sequence, nil
}

// RefreshSequenceNumber gets an Account's correct in-memory sequence number from OrbitR.
func (a *Account) RefreshSequenceNumber(hclient orbitrclient.ClientInterface) error {
	accountRequest := orbitrclient.AccountRequest{AccountID: a.GetAccountID()}
	accountDetail, err := hclient.AccountDetail(accountRequest)
	if err != nil {
		return errors.Wrap(err, "getting account detail")
	}
	a.Sequence = accountDetail.Sequence
	return nil
}
