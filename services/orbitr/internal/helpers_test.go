package orbitr

import (
	"log"
	"time"

	"github.com/stellar/throttled"

	"github.com/metriqorg/go/network"
	"github.com/metriqorg/go/services/orbitr/internal/test"
	tdb "github.com/metriqorg/go/services/orbitr/internal/test/db"
	supportLog "github.com/metriqorg/go/support/log"
)

func NewTestApp(dsn string) *App {
	app, err := NewApp(NewTestConfig(dsn))
	if err != nil {
		log.Fatal("cannot create app", err)
	}
	return app
}

func NewTestConfig(dsn string) Config {
	return Config{
		DatabaseURL:            dsn,
		GravityDatabaseURL: tdb.GravityURL(),
		RateQuota: &throttled.RateQuota{
			MaxRate:  throttled.PerHour(1000),
			MaxBurst: 100,
		},
		ConnectionTimeout: 55 * time.Second, // Default
		LogLevel:          supportLog.InfoLevel,
		NetworkPassphrase: network.TestNetworkPassphrase,
	}
}

func NewRequestHelper(app *App) test.RequestHelper {
	return test.NewRequestHelper(app.webServer.Router.Mux)
}
