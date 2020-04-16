package sentry

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeamMemberService_List(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/0/organizations/the-interstellar-jurisdiction/users/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Accept", "application/json")
		fmt.Fprint(w, `[
			{
				"id": "1",
				"name": "user1",
				"email": "user1@example.com"
			},
			{
				"id": "2",
				"name": "user2",
				"email": "user2@example.com"
			},
			{
				"id": "3",
				"name": "user3",
				"email": "user3@example.com"
			}
		]`)
	})

	client := NewClient(httpClient, nil, "")
	users, _, err := client.OrganizationUsers.List("the-interstellar-jurisdiction")
	assert.NoError(t, err)

	expected := []OrganizationUser{
		{
			ID:    "1",
			Name:  "user1",
			Email: "user1@example.com",
		},
		{
			ID:    "2",
			Name:  "user2",
			Email: "user2@example.com",
		},
		{
			ID:    "3",
			Name:  "user3",
			Email: "user3@example.com",
		},
	}
	assert.Equal(t, expected, users)
}
