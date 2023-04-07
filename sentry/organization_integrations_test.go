package sentry

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrganizationIntegrationsService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/0/organizations/the-interstellar-jurisdiction/integrations/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"cursor": "100:-1:1", "provider_key": "github"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[
			{
				"id": "123456",
				"name": "octocat",
				"icon": "https://avatars.githubusercontent.com/u/583231?v=4",
				"domainName": "github.com/octocat",
				"accountType": "Organization",
				"scopes": ["read", "write"],
				"status": "active",
				"provider": {
					"key": "github",
					"slug": "github",
					"name": "GitHub",
					"canAdd": true,
					"canDisable": false,
					"features": [
						"codeowners",
						"commits",
						"issue-basic",
						"stacktrace-link"
					],
					"aspects": {}
				},
				"configOrganization": [],
				"configData": {},
				"externalId": "87654321",
				"organizationId": 2,
				"organizationIntegrationStatus": "active",
				"gracePeriodEnd": null
			}
		]`)
	})

	ctx := context.Background()
	integrations, _, err := client.OrganizationIntegrations.List(ctx, "the-interstellar-jurisdiction", &ListOrganizationIntegrationsParams{
		ListCursorParams: ListCursorParams{
			Cursor: "100:-1:1",
		},
		ProviderKey: "github",
	})
	assert.NoError(t, err)
	expected := []*OrganizationIntegration{
		{
			ID:          "123456",
			Name:        "octocat",
			Icon:        "https://avatars.githubusercontent.com/u/583231?v=4",
			DomainName:  "github.com/octocat",
			AccountType: "Organization",
			Scopes:      []string{"read", "write"},
			Status:      "active",
			Provider: OrganizationIntegrationProvider{
				Key:        "github",
				Slug:       "github",
				Name:       "GitHub",
				CanAdd:     true,
				CanDisable: false,
				Features: []string{
					"codeowners",
					"commits",
					"issue-basic",
					"stacktrace-link",
				},
			},
			ExternalId:                    "87654321",
			OrganizationId:                2,
			OrganizationIntegrationStatus: "active",
			GracePeriodEnd:                nil,
		},
	}
	assert.Equal(t, expected, integrations)
}

func TestOrganizationIntegrationsService_Get(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/0/organizations/the-interstellar-jurisdiction/integrations/123456/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
				"id": "123456",
				"name": "octocat",
				"icon": "https://avatars.githubusercontent.com/u/583231?v=4",
				"domainName": "github.com/octocat",
				"accountType": "Organization",
				"scopes": ["read", "write"],
				"status": "active",
				"provider": {
					"key": "github",
					"slug": "github",
					"name": "GitHub",
					"canAdd": true,
					"canDisable": false,
					"features": [
						"codeowners",
						"commits",
						"issue-basic",
						"stacktrace-link"
					],
					"aspects": {}
				},
				"configOrganization": [],
				"configData": {},
				"externalId": "87654321",
				"organizationId": 2,
				"organizationIntegrationStatus": "active",
				"gracePeriodEnd": null
			}`)
	})

	ctx := context.Background()
	integration, _, err := client.OrganizationIntegrations.Get(ctx, "the-interstellar-jurisdiction", "123456")
	assert.NoError(t, err)
	expected := OrganizationIntegration{
		ID:          "123456",
		Name:        "octocat",
		Icon:        "https://avatars.githubusercontent.com/u/583231?v=4",
		DomainName:  "github.com/octocat",
		AccountType: "Organization",
		Scopes:      []string{"read", "write"},
		Status:      "active",
		Provider: OrganizationIntegrationProvider{
			Key:        "github",
			Slug:       "github",
			Name:       "GitHub",
			CanAdd:     true,
			CanDisable: false,
			Features: []string{
				"codeowners",
				"commits",
				"issue-basic",
				"stacktrace-link",
			},
		},
		ExternalId:                    "87654321",
		OrganizationId:                2,
		OrganizationIntegrationStatus: "active",
		GracePeriodEnd:                nil,
	}
	assert.Equal(t, &expected, integration)
}
