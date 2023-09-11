package main

import (
	"net/http"
	"testing"

	"github.com/lantah/go/clients/orbitrclient"
	"github.com/lantah/go/keypair"
	"github.com/lantah/go/protocols/orbitr"
	"github.com/lantah/go/services/friendbot/internal"
	"github.com/lantah/go/support/render/problem"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInitFriendbot_createMinionAccounts_success(t *testing.T) {

	randSecretKey := "SDLNA2YUQSFIWVEB57M6D3OOCJHFVCVQZJ33LPA656KJESVRK5DQUZOH"
	botKP, err := keypair.Parse(randSecretKey)
	assert.NoError(t, err)

	botKeypair := botKP.(*keypair.Full)
	botAccountID := botKeypair.Address()
	botAccountMock := orbitr.Account{
		AccountID: botAccountID,
		Sequence:  1,
	}
	botAccount := internal.Account{AccountID: botAccountID, Sequence: 1}

	orbitrClientMock := orbitrclient.MockClient{}
	orbitrClientMock.
		On("AccountDetail", orbitrclient.AccountRequest{
			AccountID: botAccountID,
		}).
		Return(botAccountMock, nil)
	orbitrClientMock.
		On("SubmitTransactionXDR", mock.Anything).
		Return(orbitr.Transaction{}, nil)

	numMinion := 1000
	minionBatchSize := 50
	submitTxRetriesAllowed := 5
	createdMinions, err := createMinionAccounts(botAccount, botKeypair, "Test Lantah Network ; 2023", "10000", "101", numMinion, minionBatchSize, submitTxRetriesAllowed, 1000, &orbitrClientMock)
	assert.NoError(t, err)

	assert.Equal(t, 1000, len(createdMinions))
}

func TestInitFriendbot_createMinionAccounts_timeoutError(t *testing.T) {
	randSecretKey := "SDLNA2YUQSFIWVEB57M6D3OOCJHFVCVQZJ33LPA656KJESVRK5DQUZOH"
	botKP, err := keypair.Parse(randSecretKey)
	assert.NoError(t, err)

	botKeypair := botKP.(*keypair.Full)
	botAccountID := botKeypair.Address()
	botAccountMock := orbitr.Account{
		AccountID: botAccountID,
		Sequence:  1,
	}
	botAccount := internal.Account{AccountID: botAccountID, Sequence: 1}

	orbitrClientMock := orbitrclient.MockClient{}
	orbitrClientMock.
		On("AccountDetail", orbitrclient.AccountRequest{
			AccountID: botAccountID,
		}).
		Return(botAccountMock, nil)

	// Successful on first 3 calls only, and then a timeout error occurs
	orbitrClientMock.
		On("SubmitTransactionXDR", mock.Anything).
		Return(orbitr.Transaction{}, nil).Times(3)
	hError := &orbitrclient.Error{
		Problem: problem.P{
			Type:   "timeout",
			Title:  "Timeout",
			Status: http.StatusGatewayTimeout,
		},
	}
	orbitrClientMock.
		On("SubmitTransactionXDR", mock.Anything).
		Return(orbitr.Transaction{}, hError)

	numMinion := 1000
	minionBatchSize := 50
	submitTxRetriesAllowed := 5
	createdMinions, err := createMinionAccounts(botAccount, botKeypair, "Test Lantah Network ; 2023", "10000", "101", numMinion, minionBatchSize, submitTxRetriesAllowed, 1000, &orbitrClientMock)
	assert.Equal(t, 150, len(createdMinions))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "after retrying 5 times: submitting create accounts tx:")
}
