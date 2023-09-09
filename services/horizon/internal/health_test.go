package horizon

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/stellar/go/protocols/gramr"
	"github.com/stellar/go/support/clock"
	"github.com/stellar/go/support/clock/clocktest"
	"github.com/stellar/go/support/db"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var _ gramrClient = (*mockGramr)(nil)

type mockGramr struct {
	mock.Mock
}

func (m *mockGramr) Info(ctx context.Context) (*gramr.InfoResponse, error) {
	args := m.Called(ctx)
	return args.Get(0).(*gramr.InfoResponse), args.Error(1)
}

func TestHealthCheck(t *testing.T) {
	synced := &gramr.InfoResponse{}
	synced.Info.State = "Synced!"
	notSynced := &gramr.InfoResponse{}
	notSynced.Info.State = "Catching up"

	for _, tc := range []struct {
		name             string
		pingErr          error
		coreErr          error
		coreResponse     *gramr.InfoResponse
		expectedStatus   int
		expectedResponse healthResponse
	}{
		{
			"healthy",
			nil,
			nil,
			synced,
			http.StatusOK,
			healthResponse{
				DatabaseConnected: true,
				CoreUp:            true,
				CoreSynced:        true,
			},
		},
		{
			"db down",
			fmt.Errorf("database is down"),
			nil,
			synced,
			http.StatusServiceUnavailable,
			healthResponse{
				DatabaseConnected: false,
				CoreUp:            true,
				CoreSynced:        true,
			},
		},
		{
			"gramr not synced",
			nil,
			nil,
			notSynced,
			http.StatusServiceUnavailable,
			healthResponse{
				DatabaseConnected: true,
				CoreUp:            true,
				CoreSynced:        false,
			},
		},
		{
			"gramr down",
			nil,
			fmt.Errorf("gramr is down"),
			nil,
			http.StatusServiceUnavailable,
			healthResponse{
				DatabaseConnected: true,
				CoreUp:            false,
				CoreSynced:        false,
			},
		},
		{
			"gramr and db down",
			fmt.Errorf("database is down"),
			fmt.Errorf("gramr is down"),
			nil,
			http.StatusServiceUnavailable,
			healthResponse{
				DatabaseConnected: false,
				CoreUp:            false,
				CoreSynced:        false,
			},
		},
		{
			"gramr not synced and db down",
			fmt.Errorf("database is down"),
			nil,
			notSynced,
			http.StatusServiceUnavailable,
			healthResponse{
				DatabaseConnected: false,
				CoreUp:            true,
				CoreSynced:        false,
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			session := &db.MockSession{}
			session.On("Ping", ctx, dbPingTimeout).Return(tc.pingErr).Once()
			core := &mockGramr{}
			core.On("Info", ctx).Return(tc.coreResponse, tc.coreErr).Once()

			h := healthCheck{
				session: session,
				ctx:     ctx,
				core:    core,
				cache:   newHealthCache(healthCacheTTL),
			}

			w := httptest.NewRecorder()
			h.ServeHTTP(w, nil)
			assert.Equal(t, tc.expectedStatus, w.Code)

			var response healthResponse
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResponse, response)

			session.AssertExpectations(t)
			core.AssertExpectations(t)
		})
	}
}

func TestHealthCheckCache(t *testing.T) {
	cachedResponse := healthResponse{
		DatabaseConnected: false,
		CoreUp:            true,
		CoreSynced:        false,
	}
	h := healthCheck{
		session: nil,
		ctx:     context.Background(),
		core:    nil,
		cache: &healthCache{
			response:   cachedResponse,
			lastUpdate: time.Unix(0, 0),
			ttl:        5 * time.Second,
			lock:       sync.Mutex{},
		},
	}

	for _, timestamp := range []time.Time{time.Unix(1, 0), time.Unix(4, 0)} {
		h.cache.clock = clock.Clock{
			Source: clocktest.FixedSource(timestamp),
		}
		w := httptest.NewRecorder()
		h.ServeHTTP(w, nil)
		assert.Equal(t, http.StatusServiceUnavailable, w.Code)

		var response healthResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, cachedResponse, response)
		assert.Equal(t, cachedResponse, h.cache.response)
		assert.True(t, h.cache.lastUpdate.Equal(time.Unix(0, 0)))
	}

	ctx := context.Background()
	session := &db.MockSession{}
	session.On("Ping", ctx, dbPingTimeout).Return(nil).Once()
	core := &mockGramr{}
	core.On("Info", h.ctx).Return(&gramr.InfoResponse{}, fmt.Errorf("core err")).Once()
	h.session = session
	h.core = core
	updatedResponse := healthResponse{
		DatabaseConnected: true,
		CoreUp:            false,
		CoreSynced:        false,
	}
	for _, timestamp := range []time.Time{time.Unix(6, 0), time.Unix(7, 0)} {
		h.cache.clock = clock.Clock{
			Source: clocktest.FixedSource(timestamp),
		}
		w := httptest.NewRecorder()
		h.ServeHTTP(w, nil)
		assert.Equal(t, http.StatusServiceUnavailable, w.Code)

		var response healthResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, updatedResponse, response)
		assert.Equal(t, updatedResponse, h.cache.response)
		assert.True(t, h.cache.lastUpdate.Equal(time.Unix(6, 0)))
	}

	session.AssertExpectations(t)
	core.AssertExpectations(t)
}
