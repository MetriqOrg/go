package resourceadapter

import (
	"context"

	protocol "github.com/lantah/go/protocols/orbitr"
	"github.com/lantah/go/xdr"
)

func PopulateAsset(ctx context.Context, dest *protocol.Asset, asset xdr.Asset) error {
	return asset.Extract(&dest.Type, &dest.Code, &dest.Issuer)
}
