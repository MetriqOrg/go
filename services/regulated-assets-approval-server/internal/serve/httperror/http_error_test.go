package httperror

import (
	"net/http"
	"testing"

	"github.com/metriqorg/go/clients/orbitrclient"
	hProtocol "github.com/metriqorg/go/protocols/orbitr"
	"github.com/metriqorg/go/support/errors"
	"github.com/metriqorg/go/support/render/problem"
	"github.com/stretchr/testify/require"
)

func TestParseOrbitRError(t *testing.T) {
	err := ParseOrbitRError(nil)
	require.Nil(t, err)

	err = ParseOrbitRError(errors.New("some error"))
	require.EqualError(t, err, "error submitting transaction: some error")

	orbitrError := orbitrclient.Error{
		Problem: problem.P{
			Type:   "bad_request",
			Title:  "Bad Request",
			Status: http.StatusBadRequest,
			Extras: map[string]interface{}{
				"result_codes": hProtocol.TransactionResultCodes{
					TransactionCode:      "tx_code_here",
					InnerTransactionCode: "",
					OperationCodes: []string{
						"op_success",
						"op_bad_auth",
					},
				},
			},
		},
	}
	err = ParseOrbitRError(orbitrError)
	require.EqualError(t, err, "error submitting transaction: problem: bad_request. full details: , &{TransactionCode:tx_code_here InnerTransactionCode: OperationCodes:[op_success op_bad_auth]}\n: orbitr error: \"Bad Request\" (tx_code_here, op_success, op_bad_auth) - check orbitr.Error.Problem for more information")
}
