package ticker

import (
	"github.com/metriqorg/go/services/ticker/internal/gql"
	"github.com/metriqorg/go/services/ticker/internal/tickerdb"
	hlog "github.com/metriqorg/go/support/log"
)

func StartGraphQLServer(s *tickerdb.TickerSession, l *hlog.Entry, port string) {
	graphql := gql.New(s, l)

	graphql.Serve(port)
}
