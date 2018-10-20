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
	expected := []Organization{
		{
			ID:             "2",
			Slug:           "the-interstellar-jurisdiction",
			Name:           "The Interstellar Jurisdiction",
			DateCreated:    mustParseTime("2017-07-17T14:10:36.141Z"),
			IsEarlyAdopter: false,
			Avatar: Avatar{
				UUID: nil,
				Type: "letter_avatar",
			},
		},
	}
	assert.Equal(t, expected, organizations)
}

func TestOrganizationService_Get(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/0/organizations/the-interstellar-jurisdiction/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
			"access": [],
			"allowSharedIssues": true,
			"availableRoles": [{
					"id": "member",
					"name": "Member"
				},
				{
					"id": "admin",
					"name": "Admin"
				},
				{
					"id": "manager",
					"name": "Manager"
				},
				{
					"id": "owner",
					"name": "Owner"
				}
			],
			"avatar": {
				"avatarType": "letter_avatar",
				"avatarUuid": null
			},
			"dataScrubber": false,
			"dataScrubberDefaults": false,
			"dateCreated": "2018-09-20T15:47:52.908Z",
			"defaultRole": "member",
			"enhancedPrivacy": false,
			"experiments": {},
			"features": [
				"sso",
				"api-keys",
				"github-apps",
				"repos",
				"new-issue-ui",
				"github-enterprise",
				"bitbucket-integration",
				"jira-integration",
				"vsts-integration",
				"suggested-commits",
				"new-teams",
				"open-membership",
				"shared-issues"
			],
			"id": "2",
			"isDefault": false,
			"isEarlyAdopter": false,
			"name": "The Interstellar Jurisdiction",
			"onboardingTasks": [],
			"openMembership": true,
			"pendingAccessRequests": 0,
			"projects": [{
					"dateCreated": "2018-09-20T15:47:56.723Z",
					"firstEvent": null,
					"hasAccess": true,
					"id": "3",
					"isBookmarked": false,
					"isMember": false,
					"latestDeploys": null,
					"name": "Prime Mover",
					"platform": null,
					"platforms": [],
					"slug": "prime-mover",
					"team": {
						"id": "2",
						"name": "Powerful Abolitionist",
						"slug": "powerful-abolitionist"
					},
					"teams": [{
						"id": "2",
						"name": "Powerful Abolitionist",
						"slug": "powerful-abolitionist"
					}]
				},
				{
					"dateCreated": "2018-09-20T15:47:52.926Z",
					"firstEvent": null,
					"hasAccess": true,
					"id": "2",
					"isBookmarked": false,
					"isMember": false,
					"latestDeploys": null,
					"name": "Pump Station",
					"platform": null,
					"platforms": [],
					"slug": "pump-station",
					"team": {
						"id": "2",
						"name": "Powerful Abolitionist",
						"slug": "powerful-abolitionist"
					},
					"teams": [{
						"id": "2",
						"name": "Powerful Abolitionist",
						"slug": "powerful-abolitionist"
					}]
				},
				{
					"dateCreated": "2018-09-20T15:48:07.592Z",
					"firstEvent": null,
					"hasAccess": true,
					"id": "4",
					"isBookmarked": false,
					"isMember": false,
					"latestDeploys": null,
					"name": "The Spoiled Yoghurt",
					"platform": null,
					"platforms": [],
					"slug": "the-spoiled-yoghurt",
					"team": {
						"id": "2",
						"name": "Powerful Abolitionist",
						"slug": "powerful-abolitionist"
					},
					"teams": [{
						"id": "2",
						"name": "Powerful Abolitionist",
						"slug": "powerful-abolitionist"
					}]
				}
			],
			"quota": {
				"accountLimit": 0,
				"maxRate": 0,
				"maxRateInterval": 60,
				"projectLimit": 100
			},
			"require2FA": false,
			"safeFields": [],
			"scrapeJavaScript": true,
			"scrubIPAddresses": false,
			"sensitiveFields": [],
			"slug": "the-interstellar-jurisdiction",
			"status": {
				"id": "active",
				"name": "active"
			},
			"storeCrashReports": false,
			"teams": [{
					"avatar": {
						"avatarType": "letter_avatar",
						"avatarUuid": null
					},
					"dateCreated": "2018-09-20T15:48:07.803Z",
					"hasAccess": true,
					"id": "3",
					"isMember": false,
					"isPending": false,
					"name": "Ancient Gabelers",
					"slug": "ancient-gabelers"
				},
				{
					"avatar": {
						"avatarType": "letter_avatar",
						"avatarUuid": null
					},
					"dateCreated": "2018-09-20T15:47:52.922Z",
					"hasAccess": true,
					"id": "2",
					"isMember": false,
					"isPending": false,
					"name": "Powerful Abolitionist",
					"slug": "powerful-abolitionist"
				}
			]
		}`)
	})

	client := NewClient(httpClient, nil, "")
	organization, _, err := client.Organizations.Get("the-interstellar-jurisdiction")
	assert.NoError(t, err)
	expected := &Organization{
		ID:   "2",
		Slug: "the-interstellar-jurisdiction",
		Status: OrganizationStatus{
			ID:   "active",
			Name: "active",
		},
		Name:           "The Interstellar Jurisdiction",
		DateCreated:    mustParseTime("2018-09-20T15:47:52.908Z"),
		IsEarlyAdopter: false,
		Avatar: Avatar{
			Type: "letter_avatar",
		},

		Quota: OrganizationQuota{
			MaxRate:         0,
			MaxRateInterval: 60,
			AccountLimit:    0,
			ProjectLimit:    100,
		},

		IsDefault:   false,
		DefaultRole: "member",
		AvailableRoles: []OrganizationAvailableRole{
			{
				ID:   "member",
				Name: "Member",
			},
			{
				ID:   "admin",
				Name: "Admin",
			},
			{
				ID:   "manager",
				Name: "Manager",
			},
			{
				ID:   "owner",
				Name: "Owner",
			},
		},
		OpenMembership:       true,
		Require2FA:           false,
		AllowSharedIssues:    true,
		EnhancedPrivacy:      false,
		DataScrubber:         false,
		DataScrubberDefaults: false,
		SensitiveFields:      []string{},
		SafeFields:           []string{},
		ScrubIPAddresses:     false,

		Access: []string{},
		Features: []string{
			"sso",
			"api-keys",
			"github-apps",
			"repos",
			"new-issue-ui",
			"github-enterprise",
			"bitbucket-integration",
			"jira-integration",
			"vsts-integration",
			"suggested-commits",
			"new-teams",
			"open-membership",
			"shared-issues",
		},
		PendingAccessRequests: 0,

		AccountRateLimit: 0,
		ProjectRateLimit: 0,

		Teams: []Team{
			{
				ID:          "3",
				Slug:        "ancient-gabelers",
				Name:        "Ancient Gabelers",
				DateCreated: mustParseTime("2018-09-20T15:48:07.803Z"),
				IsMember:    false,
				HasAccess:   true,
				IsPending:   false,
				Avatar: Avatar{
					Type: "letter_avatar",
				},
			},
			{
				ID:          "2",
				Slug:        "powerful-abolitionist",
				Name:        "Powerful Abolitionist",
				DateCreated: mustParseTime("2018-09-20T15:47:52.922Z"),
				IsMember:    false,
				HasAccess:   true,
				IsPending:   false,
				Avatar: Avatar{
					Type: "letter_avatar",
				},
			},
		},
		Projects: []ProjectSummary{
			{
				ID:           "3",
				Name:         "Prime Mover",
				Slug:         "prime-mover",
				IsBookmarked: false,
				IsMember:     false,
				HasAccess:    true,
				DateCreated:  mustParseTime("2018-09-20T15:47:56.723Z"),
				FirstEvent:   time.Time{},
				Platform:     nil,
				Platforms:    []string{},
				Team: &ProjectSummaryTeam{
					ID:   "2",
					Name: "Powerful Abolitionist",
					Slug: "powerful-abolitionist",
				},
				Teams: []ProjectSummaryTeam{
					{
						ID:   "2",
						Name: "Powerful Abolitionist",
						Slug: "powerful-abolitionist",
					},
				},
			},
			{
				ID:           "2",
				Name:         "Pump Station",
				Slug:         "pump-station",
				IsBookmarked: false,
				IsMember:     false,
				HasAccess:    true,
				DateCreated:  mustParseTime("2018-09-20T15:47:52.926Z"),
				FirstEvent:   time.Time{},
				Platform:     nil,
				Platforms:    []string{},
				Team: &ProjectSummaryTeam{
					ID:   "2",
					Name: "Powerful Abolitionist",
					Slug: "powerful-abolitionist",
				},
				Teams: []ProjectSummaryTeam{
					{
						ID:   "2",
						Name: "Powerful Abolitionist",
						Slug: "powerful-abolitionist",
					},
				},
			},
			{
				ID:           "4",
				Name:         "The Spoiled Yoghurt",
				Slug:         "the-spoiled-yoghurt",
				IsBookmarked: false,
				IsMember:     false,
				HasAccess:    true,
				DateCreated:  mustParseTime("2018-09-20T15:48:07.592Z"),
				FirstEvent:   time.Time{},
				Platform:     nil,
				Platforms:    []string{},
				Team: &ProjectSummaryTeam{
					ID:   "2",
					Name: "Powerful Abolitionist",
					Slug: "powerful-abolitionist",
				},
				Teams: []ProjectSummaryTeam{
					{
						ID:   "2",
						Name: "Powerful Abolitionist",
						Slug: "powerful-abolitionist",
					},
				},
			},
		},
	}
	assert.Equal(t, expected, organization)
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

func TestOrganizationService_Create_AgreeTerms(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/0/organizations/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertPostJSON(t, map[string]interface{}{
			"name":       "The Interstellar Jurisdiction",
			"slug":       "the-interstellar-jurisdiction",
			"agreeTerms": true,
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
		Name:       "The Interstellar Jurisdiction",
		Slug:       "the-interstellar-jurisdiction",
		AgreeTerms: Bool(true),
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
