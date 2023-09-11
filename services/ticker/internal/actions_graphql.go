package ticker

import (
	"github.com/lantah/go/services/ticker/internal/gql"
	"github.com/lantah/go/services/ticker/internal/tickerdb"
	hlog "github.com/lantah/go/support/log"
)

func StartGraphQLServer(s *tickerdb.TickerSession, l *hlog.Entry, port string) {
	graphql := gql.New(s, l)

	graphql.Serve(port)
}
