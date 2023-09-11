package actions

import (
	"net/http"
	"net/url"

	"github.com/lantah/go/protocols/orbitr"
	"github.com/lantah/go/services/orbitr/internal/ledger"
	"github.com/lantah/go/services/orbitr/internal/resourceadapter"
)

type GetRootHandler struct {
	LedgerState *ledger.State
	CoreStateGetter
	NetworkPassphrase string
	FriendbotURL      *url.URL
	OrbitRVersion    string
}

func (handler GetRootHandler) GetResource(w HeaderWriter, r *http.Request) (interface{}, error) {
	var res orbitr.Root
	templates := map[string]string{
		"accounts":           AccountsQuery{}.URITemplate(),
		"claimableBalances":  ClaimableBalancesQuery{}.URITemplate(),
		"liquidityPools":     LiquidityPoolsQuery{}.URITemplate(),
		"offers":             OffersQuery{}.URITemplate(),
		"strictReceivePaths": StrictReceivePathsQuery{}.URITemplate(),
		"strictSendPaths":    FindFixedPathsQuery{}.URITemplate(),
	}
	coreState := handler.GetCoreState()
	resourceadapter.PopulateRoot(
		r.Context(),
		&res,
		handler.LedgerState.CurrentStatus(),
		handler.OrbitRVersion,
		coreState.CoreVersion,
		handler.NetworkPassphrase,
		coreState.CurrentProtocolVersion,
		coreState.CoreSupportedProtocolVersion,
		handler.FriendbotURL,
		templates,
	)
	return res, nil
}
