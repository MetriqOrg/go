package serve

import (
	"net/http"
	"testing"
	"time"

	"github.com/metriqorg/go/clients/orbitrclient"
	"github.com/stretchr/testify/require"
)

func TestOrbitRClient(t *testing.T) {
	opts := Options{OrbitRURL: "my-orbitr.domain.com"}
	orbitrClientInterface := opts.orbitrClient()

	orbitrClient, ok := orbitrClientInterface.(*orbitrclient.Client)
	require.True(t, ok)
	require.Equal(t, "my-orbitr.domain.com", orbitrClient.OrbitRURL)

	httpClient, ok := orbitrClient.HTTP.(*http.Client)
	require.True(t, ok)
	require.Equal(t, http.Client{Timeout: 30 * time.Second}, *httpClient)
}
