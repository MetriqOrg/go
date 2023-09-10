package orbitrclient

import (
	"github.com/lantah/go/support/errors"
	"net/http"
)

// BuildURL returns the url for getting fee stats about a running orbitr instance
func (fr feeStatsRequest) BuildURL() (endpoint string, err error) {
	endpoint = fr.endpoint
	if endpoint == "" {
		err = errors.New("invalid request: too few parameters")
	}

	return
}

// HTTPRequest returns the http request for the fee stats endpoint
func (fr feeStatsRequest) HTTPRequest(orbitrURL string) (*http.Request, error) {
	endpoint, err := fr.BuildURL()
	if err != nil {
		return nil, err
	}

	return http.NewRequest("GET", orbitrURL+endpoint, nil)
}
