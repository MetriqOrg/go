package resourceadapter

import (
	"context"
	"testing"

	"github.com/lantah/go/protocols/orbitr"
	"github.com/lantah/go/services/orbitr/internal/paths"
	"github.com/lantah/go/xdr"
	"github.com/stretchr/testify/assert"
)

func TestPopulatePath(t *testing.T) {
	native := xdr.MustNewNativeAsset()
	usdc := xdr.MustNewCreditAsset("USDC", "GC3C4AKRBQLHOJ45U4XG35ESVWRDECWO5XLDGYADO6DPR3L7KIDVUMML")
	bingo := xdr.MustNewCreditAsset("BINGO", "GBZ35ZJRIKJGYH5PBKLKOZ5L6EXCNTO7BKIL7DAVVDFQ2ODJEEHHJXIM")
	p := paths.Path{
		Path:              []string{bingo.String(), native.String()},
		Source:            native.String(),
		SourceAmount:      123,
		Destination:       usdc.String(),
		DestinationAmount: 345,
	}

	var dest orbitr.Path
	assert.NoError(t, PopulatePath(context.Background(), &dest, p))

	assert.Equal(t, orbitr.Path{
		SourceAssetType:        "native",
		SourceAssetCode:        "",
		SourceAssetIssuer:      "",
		SourceAmount:           "0.000123",
		DestinationAssetType:   "credit_alphanum4",
		DestinationAssetCode:   "USDC",
		DestinationAssetIssuer: "GC3C4AKRBQLHOJ45U4XG35ESVWRDECWO5XLDGYADO6DPR3L7KIDVUMML",
		DestinationAmount:      "0.000345",
		Path: []orbitr.Asset{
			{
				Type:   "credit_alphanum12",
				Code:   "BINGO",
				Issuer: "GBZ35ZJRIKJGYH5PBKLKOZ5L6EXCNTO7BKIL7DAVVDFQ2ODJEEHHJXIM",
			},
			{
				Type:   "native",
				Code:   "",
				Issuer: "",
			},
		},
	}, dest)
}
