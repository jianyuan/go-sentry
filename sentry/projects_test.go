package sentry

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectService_List(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/0/projects/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[
    	{
    		"status": "active",
    		"slug": "the-spoiled-yoghurt",
    		"defaultEnvironment": null,
    		"features": [
    			"data-forwarding",
    			"rate-limits"
    		],
    		"color": "#bf6e3f",
    		"isPublic": false,
    		"dateCreated": "2017-07-18T19:29:44.996Z",
    		"platforms": [],
    		"callSign": "THE-SPOILED-YOGHURT",
    		"firstEvent": null,
    		"processingIssues": 0,
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
    		"isBookmarked": false,
    		"callSignReviewed": false,
    		"id": "4",
    		"name": "The Spoiled Yoghurt"
    	},
    	{
    		"status": "active",
    		"slug": "prime-mover",
    		"defaultEnvironment": null,
    		"features": [
    			"data-forwarding",
    			"rate-limits",
    			"releases"
    		],
    		"color": "#bf5b3f",
    		"isPublic": false,
    		"dateCreated": "2017-07-18T19:29:30.063Z",
    		"platforms": [],
    		"callSign": "PRIME-MOVER",
    		"firstEvent": null,
    		"processingIssues": 0,
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
    		"isBookmarked": false,
    		"callSignReviewed": false,
    		"id": "3",
    		"name": "Prime Mover"
    	},
    	{
    		"status": "active",
    		"slug": "pump-station",
    		"defaultEnvironment": null,
    		"features": [
    			"data-forwarding",
    			"rate-limits",
    			"releases"
    		],
    		"color": "#3fbf7f",
    		"isPublic": false,
    		"dateCreated": "2017-07-18T19:29:24.793Z",
    		"platforms": [],
    		"callSign": "PUMP-STATION",
    		"firstEvent": null,
    		"processingIssues": 0,
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
    		"isBookmarked": false,
    		"callSignReviewed": false,
    		"id": "2",
    		"name": "Pump Station"
    	}
    ]`)
	})

	client := NewClient(httpClient, nil, "")
	projects, _, err := client.Projects.List()
	assert.NoError(t, err)

	expected := []Project{
		{
			ID:           "4",
			Slug:         "the-spoiled-yoghurt",
			Name:         "The Spoiled Yoghurt",
			DateCreated:  mustParseTime("2017-07-18T19:29:44.996Z"),
			IsPublic:     false,
			IsBookmarked: false,
			CallSign:     "THE-SPOILED-YOGHURT",
			Color:        "#bf6e3f",
			Features: []string{
				"data-forwarding",
				"rate-limits",
			},
			Status: "active",
		},
		{
			ID:           "3",
			Slug:         "prime-mover",
			Name:         "Prime Mover",
			DateCreated:  mustParseTime("2017-07-18T19:29:30.063Z"),
			IsPublic:     false,
			IsBookmarked: false,
			CallSign:     "PRIME-MOVER",
			Color:        "#bf5b3f",
			Features: []string{
				"data-forwarding",
				"rate-limits",
				"releases",
			},
			Status: "active",
		},
		{
			ID:           "2",
			Slug:         "pump-station",
			Name:         "Pump Station",
			DateCreated:  mustParseTime("2017-07-18T19:29:24.793Z"),
			IsPublic:     false,
			IsBookmarked: false,
			CallSign:     "PUMP-STATION",
			Color:        "#3fbf7f",
			Features: []string{
				"data-forwarding",
				"rate-limits",
				"releases",
			},
			Status: "active",
		},
	}
	assert.Equal(t, expected, projects)
}
