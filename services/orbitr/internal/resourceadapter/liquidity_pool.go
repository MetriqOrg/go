package resourceadapter

import (
	"context"
	"fmt"

	"github.com/lantah/go/amount"
	protocol "github.com/lantah/go/protocols/orbitr"
	orbitrContext "github.com/lantah/go/services/orbitr/internal/context"
	"github.com/lantah/go/services/orbitr/internal/db2/history"
	"github.com/lantah/go/support/errors"
	"github.com/lantah/go/support/render/hal"
	"github.com/lantah/go/xdr"
)

// PopulateLiquidityPool fills out the resource's fields
func PopulateLiquidityPool(
	ctx context.Context,
	dest *protocol.LiquidityPool,
	liquidityPool history.LiquidityPool,
	ledger *history.Ledger,
) error {
	dest.ID = liquidityPool.PoolID
	dest.FeeBP = liquidityPool.Fee
	typ, ok := xdr.LiquidityPoolTypeToString[liquidityPool.Type]
	if !ok {
		return errors.Errorf("unknown liquidity pool type: %d", liquidityPool.Type)
	}
	dest.Type = typ
	dest.TotalTrustlines = liquidityPool.TrustlineCount
	dest.TotalShares = amount.StringFromInt64(int64(liquidityPool.ShareCount))
	for _, reserve := range liquidityPool.AssetReserves {
		dest.Reserves = append(dest.Reserves, protocol.LiquidityPoolReserve{
			Asset:  reserve.Asset.StringCanonical(),
			Amount: amount.StringFromInt64(int64(reserve.Reserve)),
		})
	}

	dest.LastModifiedLedger = liquidityPool.LastModifiedLedger

	if ledger != nil {
		dest.LastModifiedTime = &ledger.ClosedAt
	}

	lb := hal.LinkBuilder{Base: orbitrContext.BaseURL(ctx)}
	self := fmt.Sprintf("/liquidity_pools/%s", dest.ID)
	dest.Links.Self = lb.Link(self)
	dest.PT = dest.ID
	dest.Links.Transactions = lb.PagedLink(self, "transactions")
	dest.Links.Operations = lb.PagedLink(self, "operations")
	return nil
}
