package sentry

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTeamService_Create(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/0/organizations/the-interstellar-jurisdiction/teams/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertPostJSON(t, map[string]interface{}{
			"name": "Ancient Gabelers",
		}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
			"slug": "ancient-gabelers",
			"name": "Ancient Gabelers",
			"hasAccess": true,
			"isPending": false,
			"dateCreated": "2017-07-18T19:29:46.305Z",
			"isMember": false,
			"id": "3"
		}`)
	})

	client := NewClient(httpClient, nil, "")
	params := &CreateTeamParams{
		Name: "Ancient Gabelers",
	}
	team, _, err := client.Teams.Create("the-interstellar-jurisdiction", params)
	assert.NoError(t, err)

	expectedDateCreated, _ := time.Parse(time.RFC3339, "2017-07-18T19:29:46.305Z")
	expected := &Team{
		ID:          "3",
		Slug:        "ancient-gabelers",
		Name:        "Ancient Gabelers",
		DateCreated: expectedDateCreated,
		HasAccess:   true,
		IsPending:   false,
		IsMember:    false,
	}
	assert.Equal(t, expected, team)
}
