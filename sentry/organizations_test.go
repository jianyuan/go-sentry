package sentry

import (
	"fmt"
	"net/http"
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
)

func TestOrganizationService_List(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/0/organizations/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[
  {
    "name": "The Interstellar Jurisdiction",
    "slug": "the-interstellar-jurisdiction",
    "avatar": {
      "avatarUuid": null,
      "avatarType": "letter_avatar"
    },
    "dateCreated": "2017-07-17T14:10:36.141Z",
    "id": "2",
    "isEarlyAdopter": false
  }
]`)
	})

	client := NewClient(httpClient, nil, "")
	organizations, _, err := client.Organizations.List(&ListOrganizationParams{
		Cursor: "1500300636142:0:1",
	})
	assert.NoError(t, err)

	expectedDateCreated, _ := time.Parse(time.RFC3339, "2017-07-17T14:10:36.141Z")
	expected := []Organization{
		{
			ID:             "2",
			Slug:           "the-interstellar-jurisdiction",
			Name:           "The Interstellar Jurisdiction",
			DateCreated:    expectedDateCreated,
			IsEarlyAdopter: false,
			Avatar: &OrganizationAvatar{
				UUID: nil,
				Type: "letter_avatar",
			},
		},
	}
	assert.Equal(t, expected, organizations)
}
