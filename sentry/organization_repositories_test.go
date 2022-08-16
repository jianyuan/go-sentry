package sentry

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrganizationRepositoriesService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/0/organizations/the-interstellar-jurisdiction/repos/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"cursor": "100:-1:1", "status": "", "query": "foo"}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[
			{
				"id": "456123",
				"name": "octocat/Spoon-Knife",
				"url": "https://github.com/octocat/Spoon-Knife",
				"provider": {
					"id": "integrations:github",
					"name": "GitHub"
				},
				"status": "active",
				"dateCreated": "2022-08-15T06:31:49.817916Z",
				"integrationId": "123456",
				"externalSlug": "aht4davchml6srhh6mvthluoscl2lzmi"
			}
		]`)
	})

	ctx := context.Background()
	repos, _, err := client.OrganizationRepositories.List(ctx, "the-interstellar-jurisdiction", &ListOrganizationRepositoriesParams{
		ListCursorParams: ListCursorParams{
			Cursor: "100:-1:1",
		},
		Query: "foo",
	})
	assert.NoError(t, err)
	expected := []*OrganizationRepository{
		{
			ID:   "456123",
			Name: "octocat/Spoon-Knife",
			Url:  "https://github.com/octocat/Spoon-Knife",
			Provider: OrganizationRepositoryProvider{
				ID:   "integrations:github",
				Name: "GitHub",
			},
			Status:        "active",
			DateCreated:   mustParseTime("2022-08-15T06:31:49.817916Z"),
			IntegrationId: "123456",
			ExternalSlug:  "aht4davchml6srhh6mvthluoscl2lzmi",
		},
	}
	assert.Equal(t, expected, repos)
}
