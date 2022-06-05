package sentry

import (
	"fmt"
	"net/http"
	"testing"

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
			"id": "2",
			"slug": "the-interstellar-jurisdiction",
			"status": {
				"id": "active",
				"name": "active"
			},
			"name": "The Interstellar Jurisdiction",
			"dateCreated": "2022-06-05T17:31:31.170029Z",
			"isEarlyAdopter": false,
			"require2FA": false,
			"requireEmailVerification": false,
			"avatar": {
				"avatarType": "letter_avatar",
				"avatarUuid": null
			},
			"features": [
				"release-health-return-metrics",
				"slack-overage-notifications",
				"symbol-sources",
				"discover-frontend-use-events-endpoint",
				"dashboard-grid-layout",
				"performance-view",
				"open-membership",
				"integrations-stacktrace-link",
				"performance-frontend-use-events-endpoint",
				"performance-dry-run-mep",
				"auto-start-free-trial",
				"event-attachments",
				"new-widget-builder-experience-design",
				"metrics-extraction",
				"shared-issues",
				"performance-suspect-spans-view",
				"dashboards-template",
				"advanced-search",
				"performance-autogroup-sibling-spans",
				"widget-library",
				"performance-span-histogram-view",
				"performance-ops-breakdown",
				"intl-sales-tax",
				"crash-rate-alerts",
				"widget-viewer-modal",
				"invite-members-rate-limits",
				"onboarding",
				"images-loaded-v2",
				"new-weekly-report",
				"unified-span-view",
				"org-subdomains",
				"ondemand-budgets",
				"alert-crash-free-metrics",
				"custom-event-title",
				"mobile-app",
				"minute-resolution-sessions"
			],
			"experiments": {
				"TargetedOnboardingIntegrationSelectExperiment": 0,
				"TargetedOnboardingMobileRedirectExperiment": "hide"
			},
			"quota": {
				"maxRate": null,
				"maxRateInterval": 60,
				"accountLimit": 0,
				"projectLimit": 100
			},
			"isDefault": false,
			"defaultRole": "member",
			"availableRoles": [
				{
					"id": "billing",
					"name": "Billing"
				},
				{
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
			"openMembership": true,
			"allowSharedIssues": true,
			"enhancedPrivacy": false,
			"dataScrubber": false,
			"dataScrubberDefaults": false,
			"sensitiveFields": [],
			"safeFields": [],
			"storeCrashReports": 0,
			"attachmentsRole": "member",
			"debugFilesRole": "admin",
			"eventsMemberAdmin": true,
			"alertsMemberWrite": true,
			"scrubIPAddresses": false,
			"scrapeJavaScript": true,
			"allowJoinRequests": true,
			"relayPiiConfig": null,
			"trustedRelays": [],
			"access": [
				"org:write",
				"team:admin",
				"alerts:write",
				"project:releases",
				"member:admin",
				"org:admin",
				"project:read",
				"project:write",
				"alerts:read",
				"org:integrations",
				"event:admin",
				"project:admin",
				"member:write",
				"member:read",
				"org:billing",
				"team:write",
				"event:write",
				"event:read",
				"org:read",
				"team:read"
			],
			"role": "owner",
			"pendingAccessRequests": 0,
			"onboardingTasks": []
		}`)
	})

	client := NewClient(httpClient, nil, "")
	organization, _, err := client.Organizations.Get("the-interstellar-jurisdiction")
	assert.NoError(t, err)
	expected := &DetailedOrganization{
		ID:   "2",
		Slug: "the-interstellar-jurisdiction",
		Status: OrganizationStatus{
			ID:   "active",
			Name: "active",
		},
		Name:                     "The Interstellar Jurisdiction",
		DateCreated:              mustParseTime("2022-06-05T17:31:31.170029Z"),
		IsEarlyAdopter:           false,
		Require2FA:               false,
		RequireEmailVerification: false,
		Avatar: Avatar{
			Type: "letter_avatar",
		},
		Features: []string{
			"release-health-return-metrics",
			"slack-overage-notifications",
			"symbol-sources",
			"discover-frontend-use-events-endpoint",
			"dashboard-grid-layout",
			"performance-view",
			"open-membership",
			"integrations-stacktrace-link",
			"performance-frontend-use-events-endpoint",
			"performance-dry-run-mep",
			"auto-start-free-trial",
			"event-attachments",
			"new-widget-builder-experience-design",
			"metrics-extraction",
			"shared-issues",
			"performance-suspect-spans-view",
			"dashboards-template",
			"advanced-search",
			"performance-autogroup-sibling-spans",
			"widget-library",
			"performance-span-histogram-view",
			"performance-ops-breakdown",
			"intl-sales-tax",
			"crash-rate-alerts",
			"widget-viewer-modal",
			"invite-members-rate-limits",
			"onboarding",
			"images-loaded-v2",
			"new-weekly-report",
			"unified-span-view",
			"org-subdomains",
			"ondemand-budgets",
			"alert-crash-free-metrics",
			"custom-event-title",
			"mobile-app",
			"minute-resolution-sessions",
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
				ID:   "billing",
				Name: "Billing",
			},
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
		AllowSharedIssues:    true,
		EnhancedPrivacy:      false,
		DataScrubber:         false,
		DataScrubberDefaults: false,
		SensitiveFields:      []string{},
		SafeFields:           []string{},
		StoreCrashReports:    0,
		AttachmentsRole:      "member",
		DebugFilesRole:       "admin",
		EventsMemberAdmin:    true,
		AlertsMemberWrite:    true,
		ScrubIPAddresses:     false,
		ScrapeJavaScript:     true,
		AllowJoinRequests:    true,
		RelayPiiConfig:       nil,
		Access: []string{
			"org:write",
			"team:admin",
			"alerts:write",
			"project:releases",
			"member:admin",
			"org:admin",
			"project:read",
			"project:write",
			"alerts:read",
			"org:integrations",
			"event:admin",
			"project:admin",
			"member:write",
			"member:read",
			"org:billing",
			"team:write",
			"event:write",
			"event:read",
			"org:read",
			"team:read",
		},
		Role:                  "owner",
		PendingAccessRequests: 0,
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
