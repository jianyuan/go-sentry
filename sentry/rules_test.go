package sentry

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRulesService_List(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/0/projects/the-interstellar-jurisdiction/pump-station/rules/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[
			{
			  "environment": "production",
			  "actionMatch": "any",
			  "frequency": 30,
			  "name": "Notify errors",
			  "conditions": [
				{
				  "id": "sentry.rules.conditions.first_seen_event.FirstSeenEventCondition",
				  "name": "An issue is first seen"
				}
			  ],
			  "id": "123456",
			  "actions": [
				{
				  "name": "Send a notification to the Dummy Slack workspace to #dummy-channel and show tags [environment] in notification",
				  "tags": "environment",
				  "channel_id": "XX00X0X0X",
				  "workspace": "1234",
				  "id": "sentry.integrations.slack.notify_action.SlackNotifyServiceAction",
				  "channel": "#dummy-channel"
				}
			  ],
			  "dateCreated": "2019-08-24T18:12:16.321Z"
			}
		]`)

		client := NewClient(httpClient, nil, "")
		rules, _, err := client.Rules.List("the-interstellar-jurisdiction", "pump-station")
		assert.NoError(t, err)

		expected := []Rule{
			{
				ID: "123456",
				ActionMatch: "all",
				Environment: "production",
				Frequency: 30,
				Name: "Notify errors",
				Conditions: []RuleCondition{
					{
						ID: "sentry.rules.conditions.first_seen_event.FirstSeenEventCondition",
						Name: "An issue is first seen",
					},
				},
				Actions: []RuleAction{
					{
						ID: "sentry.integrations.slack.notify_action.SlackNotifyServiceAction",
						Name: "Send a notification to the Dummy Slack workspace to #dummy-channel and show tags [environment] in notification",
						Tags: "environment",
						ChannelID: "XX00X0X0X",
						Channel: "#dummy-channel",
						Workspace: "1234",
					},
				},
				Created: mustParseTime("2019-08-24T18:12:16.321Z"),
			},
		}
		assert.Equal(t, expected, rules)
	})
}