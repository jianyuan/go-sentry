package sentry

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrganizationCodeMappingsService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/0/organizations/the-interstellar-jurisdiction/code-mappings/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"cursor": "100:-1:1", "integrationId": "123456"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[
			{
				"id": "54321",
				"projectId": "7654321",
				"projectSlug": "spoon-knife",
				"repoId": "456123",
				"repoName": "octocat/Spoon-Knife",
				"integrationId": "123456",
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
				"stackRoot": "/",
				"sourceRoot": "src/",
				"defaultBranch": "main"
			}
		]`)
	})

	ctx := context.Background()
	integrations, _, err := client.OrganizationCodeMappings.List(ctx, "the-interstellar-jurisdiction", &ListOrganizationCodeMappingsParams{
		ListCursorParams: ListCursorParams{
			Cursor: "100:-1:1",
		},
		IntegrationId: "123456",
	})
	assert.NoError(t, err)
	expected := []*OrganizationCodeMapping{
		{
			ID:            "54321",
			ProjectId:     "7654321",
			ProjectSlug:   "spoon-knife",
			RepoId:        "456123",
			RepoName:      "octocat/Spoon-Knife",
			IntegrationId: "123456",
			Provider: &OrganizationIntegrationProvider{
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
			StackRoot:     "/",
			SourceRoot:    "src/",
			DefaultBranch: "main",
		},
	}
	assert.Equal(t, expected, integrations)
}
