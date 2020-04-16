package sentry

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeamMemberService_Get(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/0/organizations/the-interstellar-jurisdiction/members/1/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Accept", "application/json")
		fmt.Fprint(w, `{
				"id": "1",
				"name": "user1@example.com",
				"teams": ["a", "b"]
			}`,
		)
	})

	client := NewClient(httpClient, nil, "")
	member, _, err := client.TeamMembers.Get("the-interstellar-jurisdiction", "1")
	assert.NoError(t, err)

	expected := TeamMember{
		ID:    "1",
		Name:  "user1@example.com",
		Teams: []string{"a", "b"},
	}
	assert.Equal(t, expected, member)
}

func TestTeamMemberService_Update(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/0/organizations/the-interstellar-jurisdiction/members/1/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "PUT", r)
		assertPostJSON(t, map[string]interface{}{
			"teams": []interface{}{"a", "b", "c"},
		}, r)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Accept", "application/json")
		fmt.Fprint(w, `{
				"id": "1",
				"name": "user1@example.com",
				"teams": ["a", "b", "c"]
			}`,
		)
	})

	params := &UpdateTeamMemberParams{
		Teams: []string{"a", "b", "c"},
	}

	client := NewClient(httpClient, nil, "")
	member, _, err := client.TeamMembers.Update("the-interstellar-jurisdiction", "1", params)
	assert.NoError(t, err)

	expected := TeamMember{
		ID:    "1",
		Name:  "user1@example.com",
		Teams: []string{"a", "b", "c"},
	}
	assert.Equal(t, expected, member)
}
