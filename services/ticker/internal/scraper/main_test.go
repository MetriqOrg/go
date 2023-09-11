package scraper

import (
	"testing"
	"time"

	orbitrclient "github.com/lantah/go/clients/orbitrclient"
	hProtocol "github.com/lantah/go/protocols/orbitr"
	"github.com/lantah/go/support/errors"
	"github.com/lantah/go/support/log"
	"github.com/stretchr/testify/assert"
)

func Test_ScraperConfig_FetchAllTrades_doesntCrashWhenReceivesAnError(t *testing.T) {
	orbitrClient := &orbitrclient.MockClient{}
	orbitrClient.
		On("Trades", orbitrclient.TradeRequest{Limit: 200, Order: orbitrclient.OrderDesc}).
		Return(hProtocol.TradesPage{}, errors.New("something went wrong"))

	sc := ScraperConfig{
		Logger: log.DefaultLogger,
		Client: orbitrClient,
	}

	trades, err := sc.FetchAllTrades(time.Now(), 0)
	assert.EqualError(t, err, "something went wrong")
	assert.Empty(t, trades)
}

func Test_ScraperConfig_FetchAllTrades_doesntCrashWhenReceivesEmptyList(t *testing.T) {
	orbitrClient := &orbitrclient.MockClient{}
	orbitrClient.
		On("Trades", orbitrclient.TradeRequest{Limit: 200, Order: orbitrclient.OrderDesc}).
		Return(hProtocol.TradesPage{}, nil)

	sc := ScraperConfig{
		Logger: log.DefaultLogger,
		Client: orbitrClient,
	}

	trades, err := sc.FetchAllTrades(time.Now(), 0)
	assert.NoError(t, err)
	assert.Empty(t, trades)
}
