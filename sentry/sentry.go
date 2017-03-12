package sentry

import (
	"net/http"

	"github.com/dghubble/sling"
)

type Client struct {
	sling         *sling.Sling
	Organizations *OrganizationsService
}

// NewClient returns a new Sentry API client.
// If a nil httpClient is given, the http.DefaultClient will be used.
func NewClient(httpClient *http.Client, token string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	base := sling.New().Base("https://sentry.io/api/0/").Client(httpClient)

	if token != "" {
		base.Add("Authorization", "Bearer "+token)
	}

	c := &Client{
		sling:         base,
		Organizations: newOrganizationsService(base.New()),
	}
	return c
}
