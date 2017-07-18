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
		assertQuery(t, map[string]string{"cursor": "1500300636142:0:1"}, r)
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

func TestOrganizationService_Create(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/0/organizations/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertPostJSON(t, map[string]interface{}{
			"name": "The Interstellar Jurisdiction",
			"slug": "the-interstellar-jurisdiction",
		}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
			"name": "The Interstellar Jurisdiction",
			"slug": "the-interstellar-jurisdiction",
			"id": "2"
		}`)
	})

	client := NewClient(httpClient, nil, "")
	params := &CreateOrganizationParams{
		Name: "The Interstellar Jurisdiction",
		Slug: "the-interstellar-jurisdiction",
	}
	organization, _, err := client.Organizations.Create(params)
	assert.NoError(t, err)
	expected := &Organization{
		ID:   "2",
		Name: "The Interstellar Jurisdiction",
		Slug: "the-interstellar-jurisdiction",
	}
	assert.Equal(t, expected, organization)
}

func TestOrganizationService_Update(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/0/organizations/badly-misnamed/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "PUT", r)
		assertPostJSON(t, map[string]interface{}{
			"name": "Impeccably Designated",
			"slug": "impeccably-designated",
		}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
			"name": "Impeccably Designated",
			"slug": "impeccably-designated",
			"id": "2"
		}`)
	})

	client := NewClient(httpClient, nil, "")
	params := &UpdateOrganizationParams{
		Name: "Impeccably Designated",
		Slug: "impeccably-designated",
	}
	organization, _, err := client.Organizations.Update("badly-misnamed", params)
	assert.NoError(t, err)
	expected := &Organization{
		ID:   "2",
		Name: "Impeccably Designated",
		Slug: "impeccably-designated",
	}
	assert.Equal(t, expected, organization)
}

func TestOrganizationService_Delete(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/0/organizations/the-interstellar-jurisdiction/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "DELETE", r)
	})

	client := NewClient(httpClient, nil, "")
	_, err := client.Organizations.Delete("the-interstellar-jurisdiction")
	assert.NoError(t, err)
}
