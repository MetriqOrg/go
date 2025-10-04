// Package reap contains the history reaping subsystem for orbitr.  This system
// is designed to remove data from the history database such that it does not
// grow indefinitely.  The system can be configured with a number of ledgers to
// maintain at a minimum.
package reap

import (
	"context"

	"github.com/metriqorg/go/services/orbitr/internal/db2/history"
	"github.com/metriqorg/go/services/orbitr/internal/ledger"
	"github.com/metriqorg/go/support/db"
)

// System represents the history reaping subsystem of orbitr.
type System struct {
	HistoryQ       *history.Q
	RetentionCount uint
	ledgerState    *ledger.State
	ctx            context.Context
	cancel         context.CancelFunc
}

// New initializes the reaper, causing it to begin polling the gravity
// database for now ledgers and ingesting data into the orbitr database.
func New(retention uint, dbSession db.SessionInterface, ledgerState *ledger.State) *System {
	ctx, cancel := context.WithCancel(context.Background())

	r := &System{
		HistoryQ:       &history.Q{dbSession.Clone()},
		RetentionCount: retention,
		ledgerState:    ledgerState,
		ctx:            ctx,
		cancel:         cancel,
	}

	return r
}
