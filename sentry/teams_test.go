package sentry

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeamService_List(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/0/organizations/the-interstellar-jurisdiction/teams/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[
			{
				"avatar": {
					"avatarType": "letter_avatar",
					"avatarUuid": null
				},
				"dateCreated": "2018-11-06T21:20:08.115Z",
				"hasAccess": true,
				"id": "3",
				"isMember": true,
				"isPending": false,
				"memberCount": 1,
				"name": "Ancient Gabelers",
				"projects": [],
				"slug": "ancient-gabelers"
			},
			{
				"avatar": {
					"avatarType": "letter_avatar",
					"avatarUuid": null
				},
				"dateCreated": "2018-11-06T21:19:55.114Z",
				"hasAccess": true,
				"id": "2",
				"isMember": true,
				"isPending": false,
				"memberCount": 1,
				"name": "Powerful Abolitionist",
				"projects": [
					{
						"avatar": {
							"avatarType": "letter_avatar",
							"avatarUuid": null
						},
						"color": "#bf5b3f",
						"dateCreated": "2018-11-06T21:19:58.536Z",
						"features": [
							"releases",
							"sample-events",
							"minidump",
							"servicehooks",
							"rate-limits",
							"data-forwarding"
						],
						"firstEvent": null,
						"hasAccess": true,
						"id": "3",
						"isBookmarked": false,
						"isInternal": false,
						"isMember": true,
						"isPublic": false,
						"name": "Prime Mover",
						"platform": null,
						"slug": "prime-mover",
						"status": "active"
					},
					{
						"avatar": {
							"avatarType": "letter_avatar",
							"avatarUuid": null
						},
						"color": "#3fbf7f",
						"dateCreated": "2018-11-06T21:19:55.121Z",
						"features": [
							"releases",
							"sample-events",
							"minidump",
							"servicehooks",
							"rate-limits",
							"data-forwarding"
						],
						"firstEvent": null,
						"hasAccess": true,
						"id": "2",
						"isBookmarked": false,
						"isInternal": false,
						"isMember": true,
						"isPublic": false,
						"name": "Pump Station",
						"platform": null,
						"slug": "pump-station",
						"status": "active"
					},
					{
						"avatar": {
							"avatarType": "letter_avatar",
							"avatarUuid": null
						},
						"color": "#bf6e3f",
						"dateCreated": "2018-11-06T21:20:08.064Z",
						"features": [
							"servicehooks",
							"sample-events",
							"data-forwarding",
							"rate-limits",
							"minidump"
						],
						"firstEvent": null,
						"hasAccess": true,
						"id": "4",
						"isBookmarked": false,
						"isInternal": false,
						"isMember": true,
						"isPublic": false,
						"name": "The Spoiled Yoghurt",
						"platform": null,
						"slug": "the-spoiled-yoghurt",
						"status": "active"
					}
				],
				"slug": "powerful-abolitionist"
			}
		]`)
	})

	client := NewClient(httpClient, nil, "")
	teams, _, err := client.Teams.List("the-interstellar-jurisdiction")
	assert.NoError(t, err)

	expected := []Team{
		{
			ID:          "3",
			Slug:        "ancient-gabelers",
			Name:        "Ancient Gabelers",
			DateCreated: mustParseTime("2017-07-18T19:29:46.305Z"),
			HasAccess:   true,
			IsPending:   false,
			IsMember:    false,
		},
		{
			ID:          "2",
			Slug:        "powerful-abolitionist",
			Name:        "Powerful Abolitionist",
			DateCreated: mustParseTime("2017-07-18T19:29:24.743Z"),
			HasAccess:   true,
			IsPending:   false,
			IsMember:    false,
		},
	}
	assert.Equal(t, expected, teams)
}

func TestTeamService_Get(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/0/teams/the-interstellar-jurisdiction/powerful-abolitionist/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
			"slug": "powerful-abolitionist",
			"name": "Powerful Abolitionist",
			"hasAccess": true,
			"isPending": false,
			"dateCreated": "2017-07-18T19:29:24.743Z",
			"isMember": false,
			"organization": {
				"name": "The Interstellar Jurisdiction",
				"slug": "the-interstellar-jurisdiction",
				"avatar": {
					"avatarUuid": null,
					"avatarType": "letter_avatar"
				},
				"dateCreated": "2017-07-18T19:29:24.565Z",
				"id": "2",
				"isEarlyAdopter": false
			},
			"id": "2"
		}`)
	})

	client := NewClient(httpClient, nil, "")
	team, _, err := client.Teams.Get("the-interstellar-jurisdiction", "powerful-abolitionist")
	assert.NoError(t, err)

	expected := &Team{
		ID:          "2",
		Slug:        "powerful-abolitionist",
		Name:        "Powerful Abolitionist",
		DateCreated: mustParseTime("2017-07-18T19:29:24.743Z"),
		HasAccess:   true,
		IsPending:   false,
		IsMember:    false,
	}
	assert.Equal(t, expected, team)
}

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

	expected := &Team{
		ID:          "3",
		Slug:        "ancient-gabelers",
		Name:        "Ancient Gabelers",
		DateCreated: mustParseTime("2017-07-18T19:29:46.305Z"),
		HasAccess:   true,
		IsPending:   false,
		IsMember:    false,
	}
	assert.Equal(t, expected, team)
}

func TestTeamService_Update(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/0/teams/the-interstellar-jurisdiction/the-obese-philosophers/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "PUT", r)
		assertPostJSON(t, map[string]interface{}{
			"name": "The Inflated Philosophers",
		}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
			"slug": "the-obese-philosophers",
			"name": "The Inflated Philosophers",
			"hasAccess": true,
			"isPending": false,
			"dateCreated": "2017-07-18T19:30:14.736Z",
			"isMember": false,
			"id": "4"
		}`)
	})

	client := NewClient(httpClient, nil, "")
	params := &UpdateTeamParams{
		Name: "The Inflated Philosophers",
	}
	team, _, err := client.Teams.Update("the-interstellar-jurisdiction", "the-obese-philosophers", params)
	assert.NoError(t, err)
	expected := &Team{
		ID:          "4",
		Slug:        "the-obese-philosophers",
		Name:        "The Inflated Philosophers",
		DateCreated: mustParseTime("2017-07-18T19:30:14.736Z"),
		HasAccess:   true,
		IsPending:   false,
		IsMember:    false,
	}
	assert.Equal(t, expected, team)
}

func TestTeamService_Delete(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/0/teams/the-interstellar-jurisdiction/the-obese-philosophers/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "DELETE", r)
	})

	client := NewClient(httpClient, nil, "")
	_, err := client.Teams.Delete("the-interstellar-jurisdiction", "the-obese-philosophers")
	assert.NoError(t, err)

}
