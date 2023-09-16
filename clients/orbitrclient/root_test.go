package orbitrclient

import (
	"testing"

	"github.com/lantah/go/support/http/httptest"
	"github.com/stretchr/testify/assert"
)

func TestRoot(t *testing.T) {
	hmock := httptest.NewClient()
	client := &Client{
		OrbitRURL: "https://localhost/",
		HTTP:       hmock,
	}

	// happy path
	hmock.On(
		"GET",
		"https://localhost/",
	).ReturnString(200, rootResponse)

	root, err := client.Root()
	if assert.NoError(t, err) {
		assert.Equal(t, root.OrbitRVersion, "0.17.6-unstable-bc999a67d0b2413d8abd76153a56733c7d517484")
		assert.Equal(t, root.GravityVersion, "gravity 11.0.0 (236f831521b6724c0ae63906416faa997ef27e19)")
		assert.Equal(t, root.OrbitRSequence, int32(84959))
		assert.Equal(t, root.NetworkPassphrase, "Test Lantah Network ; 2023")
	}

	// failure response
	hmock.On(
		"GET",
		"https://localhost/",
	).ReturnString(404, notFoundResponse)

	_, err = client.Root()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "orbitr error")
		orbitrError, ok := err.(*Error)
		assert.Equal(t, ok, true)
		assert.Equal(t, orbitrError.Problem.Title, "Resource Missing")
	}

	// connection error
	hmock.On(
		"GET",
		"https://localhost/",
	).ReturnError("http.Client error")

	_, err = client.Root()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "http.Client error")
		_, ok := err.(*Error)
		assert.Equal(t, ok, false)
	}
}

var rootResponse = `{
  "_links": {
    "account": {
      "href": "https://orbitr-testnet.lantah.network/accounts/{account_id}",
      "templated": true
    },
    "account_transactions": {
      "href": "https://orbitr-testnet.lantah.network/accounts/{account_id}/transactions{?cursor,limit,order}",
      "templated": true
    },
    "assets": {
      "href": "https://orbitr-testnet.lantah.network/assets{?asset_code,asset_issuer,cursor,limit,order}",
      "templated": true
    },
    "friendbot": {
      "href": "https://friendbot.stellar.org/{?addr}",
      "templated": true
    },
    "metrics": {
      "href": "https://orbitr-testnet.lantah.network/metrics"
    },
    "order_book": {
      "href": "https://orbitr-testnet.lantah.network/order_book{?selling_asset_type,selling_asset_code,selling_asset_issuer,buying_asset_type,buying_asset_code,buying_asset_issuer,limit}",
      "templated": true
    },
    "self": {
      "href": "https://orbitr-testnet.lantah.network/"
    },
    "transaction": {
      "href": "https://orbitr-testnet.lantah.network/transactions/{hash}",
      "templated": true
    },
    "transactions": {
      "href": "https://orbitr-testnet.lantah.network/transactions{?cursor,limit,order}",
      "templated": true
    }
  },
  "orbitr_version": "0.17.6-unstable-bc999a67d0b2413d8abd76153a56733c7d517484",
  "core_version": "gravity 11.0.0 (236f831521b6724c0ae63906416faa997ef27e19)",
  "history_latest_ledger": 84959,
  "history_elder_ledger": 1,
  "core_latest_ledger": 84959,
  "network_passphrase": "Test Lantah Network ; 2023",
  "current_protocol_version": 10,
  "core_supported_protocol_version": 11
}`
