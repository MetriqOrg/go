package resourceadapter

import (
	"context"

	protocol "github.com/metriqorg/go/protocols/orbitr"
	"github.com/metriqorg/go/xdr"
)

func PopulateAsset(ctx context.Context, dest *protocol.Asset, asset xdr.Asset) error {
	return asset.Extract(&dest.Type, &dest.Code, &dest.Issuer)
}
