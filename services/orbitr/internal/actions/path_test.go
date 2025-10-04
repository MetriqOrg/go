package actions

import (
	"context"
	"net/http"
	"testing"

	orbitrContext "github.com/metriqorg/go/services/orbitr/internal/context"
	"github.com/metriqorg/go/services/orbitr/internal/db2/history"
	"github.com/metriqorg/go/services/orbitr/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestAssetsForAddressRequiresTransaction(t *testing.T) {
	tt := test.Start(t)
	defer tt.Finish()
	test.ResetOrbitRDB(t, tt.OrbitRDB)
	q := &history.Q{tt.OrbitRSession()}

	r := &http.Request{}
	ctx := context.WithValue(
		r.Context(),
		&orbitrContext.SessionContextKey,
		q,
	)

	_, _, err := assetsForAddress(r.WithContext(ctx), "GCATOZ7YJV2FANQQLX47TIV6P7VMPJCEEJGQGR6X7TONPKBN3UCLKEIS")
	assert.EqualError(t, err, "cannot be called outside of a transaction")

	assert.NoError(t, q.Begin(ctx))
	defer q.Rollback()

	_, _, err = assetsForAddress(r.WithContext(ctx), "GCATOZ7YJV2FANQQLX47TIV6P7VMPJCEEJGQGR6X7TONPKBN3UCLKEIS")
	assert.EqualError(t, err, "should only be called in a repeatable read transaction")
}
