package filters

import (
	"testing"

	"github.com/metriqorg/go/services/orbitr/internal/db2/history"
	"github.com/metriqorg/go/services/orbitr/internal/test"
)

func TestItGetsFilters(t *testing.T) {
	tt := test.Start(t)
	defer tt.Finish()
	test.ResetOrbitRDB(t, tt.OrbitRDB)
	q := &history.Q{tt.OrbitRSession()}

	filtersService := NewFilters()

	ingestFilters := filtersService.GetFilters(q, tt.Ctx)

	// should be total of filters implemented in the system
	tt.Assert.Len(ingestFilters, 2)
}
