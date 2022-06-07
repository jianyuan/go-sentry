package sentry

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"encoding/json"

	"github.com/stretchr/testify/assert"
)

func setup() (client *Client, mux *http.ServeMux, serverURL string, teardown func()) {
	mux = http.NewServeMux()
	server := httptest.NewServer(mux)
	client = NewClient(nil)
	url, _ := url.Parse(server.URL + "/api/")
	client.BaseURL = url
	return client, mux, server.URL, server.Close
}

// RewriteTransport rewrites https requests to http to avoid TLS cert issues
// during testing.
type RewriteTransport struct {
	Transport http.RoundTripper
}

// RoundTrip rewrites the request scheme to http and calls through to the
// composed RoundTripper or if it is nil, to the http.DefaultTransport.
func (t *RewriteTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.URL.Scheme = "http"
	if t.Transport == nil {
		return http.DefaultTransport.RoundTrip(req)
	}
	return t.Transport.RoundTrip(req)
}

func assertMethod(t *testing.T, expectedMethod string, req *http.Request) {
	assert.Equal(t, expectedMethod, req.Method)
}

// assertQuery tests that the Request has the expected url query key/val pairs
func assertQuery(t *testing.T, expected map[string]string, req *http.Request) {
	queryValues := req.URL.Query()
	expectedValues := url.Values{}
	for key, value := range expected {
		expectedValues.Add(key, value)
	}
	assert.Equal(t, expectedValues, queryValues)
}

// assertPostJSON tests that the Request has the expected JSON in its Body
func assertPostJSON(t *testing.T, expected interface{}, req *http.Request) {
	var actual interface{}

	d := json.NewDecoder(req.Body)
	d.UseNumber()

	err := d.Decode(&actual)
	assert.NoError(t, err)
	assert.EqualValues(t, expected, actual)
}

// assertPostJSON tests that the request has the expected values in its body.
func assertPostJSONValue(t *testing.T, expected interface{}, req *http.Request) {
	var actual interface{}

	d := json.NewDecoder(req.Body)
	d.UseNumber()

	err := d.Decode(&actual)
	assert.NoError(t, err)
	assert.ObjectsAreEqualValues(expected, actual)
}

func mustParseTime(value string) time.Time {
	t, err := time.Parse(time.RFC3339, value)
	if err != nil {
		panic(fmt.Sprintf("mustParseTime: %s", err))
	}
	return t
}

func TestNewClient(t *testing.T) {
	c := NewClient(nil)

	assert.Equal(t, "https://sentry.io/api/", c.BaseURL.String())
}

func TestNewOnPremiseClient(t *testing.T) {
	testCases := []struct {
		baseURL string
	}{
		{"https://example.com"},
		{"https://example.com/"},
		{"https://example.com/api"},
		{"https://example.com/api/"},
	}
	for _, tc := range testCases {
		t.Run(tc.baseURL, func(t *testing.T) {
			c, err := NewOnPremiseClient(tc.baseURL, nil)

			assert.NoError(t, err)
			assert.Equal(t, "https://example.com/api/", c.BaseURL.String())
		})
	}

}
