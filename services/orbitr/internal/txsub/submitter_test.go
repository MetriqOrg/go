package txsub

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"

	"github.com/metriqorg/go/services/orbitr/internal/test"
)

func TestDefaultSubmitter(t *testing.T) {
	ctx := test.Context()
	// submits to the configured gravity instance correctly
	server := test.NewStaticMockServer(`{
		"status": "PENDING",
		"error": null
		}`)
	defer server.Close()

	s := NewDefaultSubmitter(http.DefaultClient, server.URL)
	sr := s.Submit(ctx, "hello")
	assert.Nil(t, sr.Err)
	assert.True(t, sr.Duration > 0)
	assert.Equal(t, "hello", server.LastRequest.URL.Query().Get("blob"))

	// Succeeds when gravity gives the DUPLICATE response.
	server = test.NewStaticMockServer(`{
				"status": "DUPLICATE",
				"error": null
				}`)
	defer server.Close()

	s = NewDefaultSubmitter(http.DefaultClient, server.URL)
	sr = s.Submit(ctx, "hello")
	assert.Nil(t, sr.Err)

	// Errors when the gravity url is empty

	s = NewDefaultSubmitter(http.DefaultClient, "")
	sr = s.Submit(ctx, "hello")
	assert.NotNil(t, sr.Err)

	//errors when the gravity url is not parseable

	s = NewDefaultSubmitter(http.DefaultClient, "http://Not a url")
	sr = s.Submit(ctx, "hello")
	assert.NotNil(t, sr.Err)

	// errors when the gravity url is not reachable
	s = NewDefaultSubmitter(http.DefaultClient, "http://127.0.0.1:65535")
	sr = s.Submit(ctx, "hello")
	assert.NotNil(t, sr.Err)

	// errors when the gravity returns an unparseable response
	server = test.NewStaticMockServer(`{`)
	defer server.Close()

	s = NewDefaultSubmitter(http.DefaultClient, server.URL)
	sr = s.Submit(ctx, "hello")
	assert.NotNil(t, sr.Err)

	// errors when the gravity returns an exception response
	server = test.NewStaticMockServer(`{"exception": "Invalid XDR"}`)
	defer server.Close()

	s = NewDefaultSubmitter(http.DefaultClient, server.URL)
	sr = s.Submit(ctx, "hello")
	assert.NotNil(t, sr.Err)
	assert.Contains(t, sr.Err.Error(), "Invalid XDR")

	// errors when the gravity returns an unrecognized status
	server = test.NewStaticMockServer(`{"status": "NOTREAL"}`)
	defer server.Close()

	s = NewDefaultSubmitter(http.DefaultClient, server.URL)
	sr = s.Submit(ctx, "hello")
	assert.NotNil(t, sr.Err)
	assert.Contains(t, sr.Err.Error(), "NOTREAL")

	// errors when the gravity returns an error response
	server = test.NewStaticMockServer(`{"status": "ERROR", "error": "1234"}`)
	defer server.Close()

	s = NewDefaultSubmitter(http.DefaultClient, server.URL)
	sr = s.Submit(ctx, "hello")
	assert.IsType(t, &FailedTransactionError{}, sr.Err)
	ferr := sr.Err.(*FailedTransactionError)
	assert.Equal(t, "1234", ferr.ResultXDR)
}
