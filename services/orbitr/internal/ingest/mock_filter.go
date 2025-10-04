package ingest

import (
	"context"

	"github.com/metriqorg/go/services/orbitr/internal/db2/history"
	"github.com/metriqorg/go/services/orbitr/internal/ingest/processors"
	"github.com/stretchr/testify/mock"
)

type MockFilters struct {
	mock.Mock
}

func (m *MockFilters) GetFilters(filterQ history.QFilter, ctx context.Context) []processors.LedgerTransactionFilterer {
	return []processors.LedgerTransactionFilterer{}
}
