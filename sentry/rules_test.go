package sentry

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
				  "name": "An issue is first seen",
				  "value": 500,
				  "interval": "1h"
				}
			  ],
			  "id": "12345",
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
	})
	client := NewClient(httpClient, nil, "")
	rules, _, err := client.Rules.List("the-interstellar-jurisdiction", "pump-station")
	require.NoError(t, err)

	environment := "production"
	expected := []Rule{
		{
			ID:          "12345",
			ActionMatch: "any",
			Environment: &environment,
			Frequency:   30,
			Name:        "Notify errors",
			Conditions: []ConditionType{
				{
					"id":       "sentry.rules.conditions.first_seen_event.FirstSeenEventCondition",
					"name":     "An issue is first seen",
					"value":    float64(500),
					"interval": "1h",
				},
			},
			Actions: []ActionType{
				{
					"id":         "sentry.integrations.slack.notify_action.SlackNotifyServiceAction",
					"name":       "Send a notification to the Dummy Slack workspace to #dummy-channel and show tags [environment] in notification",
					"tags":       "environment",
					"channel_id": "XX00X0X0X",
					"channel":    "#dummy-channel",
					"workspace":  "1234",
				},
			},
			Created: mustParseTime("2019-08-24T18:12:16.321Z"),
		},
	}
	require.Equal(t, expected, rules)

}

func TestRulesService_Create(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/0/projects/the-interstellar-jurisdiction/pump-station/rules/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertPostJSONValue(t, map[string]interface{}{
			"actionMatch": "all",
			"environment": "production",
			"frequency":   30,
			"name":        "Notify errors",
			"conditions": []map[string]interface{}{
				{
					"interval": "1h",
					"name":     "The issue is seen more than 10 times in 1h",
					"value":    10,
					"id":       "sentry.rules.conditions.event_frequency.EventFrequencyCondition",
				},
			},
			"actions": []map[string]interface{}{
				{
					"id":         "sentry.integrations.slack.notify_action.SlackNotifyServiceAction",
					"name":       "Send a notification to the Dummy Slack workspace to #dummy-channel and show tags [environment] in notification",
					"tags":       "environment",
					"channel":    "#dummy-channel",
					"channel_id": "XX00X0X0X",
					"workspace":  "1234",
				},
			},
		}, r)

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
			"id": "123456",
			"actionMatch": "all",
			"environment": "production",
			"frequency": 30,
			"name": "Notify errors",
			"conditions": [
				{
					"interval": "1h",
					"name": "The issue is seen more than 10 times in 1h",
					"value": 10,
					"id": "sentry.rules.conditions.event_frequency.EventFrequencyCondition"
				}
			],
			"actions": [
				{
					"id": "sentry.integrations.slack.notify_action.SlackNotifyServiceAction",
					"name": "Send a notification to the Dummy Slack workspace to #dummy-channel and show tags [environment] in notification",
					"tags": "environment",
					"channel_id": "XX00X0X0X",
					"workspace": "1234",
					"channel": "#dummy-channel"
				}
			],
			"dateCreated": "2019-08-24T18:12:16.321Z"
		}`)
	})

	params := &CreateRuleParams{
		ActionMatch: "all",
		Environment: "production",
		Frequency:   30,
		Name:        "Notify errors",
		Conditions: []ConditionType{
			{
				"interval": "1h",
				"name":     "The issue is seen more than 10 times in 1h",
				"value":    float64(10),
				"id":       "sentry.rules.conditions.event_frequency.EventFrequencyCondition",
			},
		},
		Actions: []ActionType{
			{
				"id":         "sentry.integrations.slack.notify_action.SlackNotifyServiceAction",
				"name":       "Send a notification to the Dummy Slack workspace to #dummy-channel and show tags [environment] in notification",
				"tags":       "environment",
				"channel_id": "XX00X0X0X",
				"workspace":  "1234",
				"channel":    "#dummy-channel",
			},
		},
	}

	client := NewClient(httpClient, nil, "")
	rules, _, err := client.Rules.Create("the-interstellar-jurisdiction", "pump-station", params)
	require.NoError(t, err)

	environment := "production"
	expected := &Rule{
		ID:          "123456",
		ActionMatch: "all",
		Environment: &environment,
		Frequency:   30,
		Name:        "Notify errors",
		Conditions: []ConditionType{
			{
				"interval": "1h",
				"name":     "The issue is seen more than 10 times in 1h",
				"value":    float64(10),
				"id":       "sentry.rules.conditions.event_frequency.EventFrequencyCondition",
			},
		},
		Actions: []ActionType{
			{
				"id":         "sentry.integrations.slack.notify_action.SlackNotifyServiceAction",
				"name":       "Send a notification to the Dummy Slack workspace to #dummy-channel and show tags [environment] in notification",
				"tags":       "environment",
				"channel_id": "XX00X0X0X",
				"channel":    "#dummy-channel",
				"workspace":  "1234",
			},
		},
		Created: mustParseTime("2019-08-24T18:12:16.321Z"),
	}
	require.Equal(t, expected, rules)

}

func TestRulesService_Create_Async_Task(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/0/projects/the-interstellar-jurisdiction/pump-station/rule-task/fakeuuid/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
			"status": "success",
			"error": null,
			"rule": {
				"id": "123456",
				"actionMatch": "all",
				"environment": "production",
				"frequency": 30,
				"name": "Notify errors",
				"conditions": [
					{
						"interval": "1h",
						"name": "The issue is seen more than 10 times in 1h",
						"value": 10,
						"id": "sentry.rules.conditions.event_frequency.EventFrequencyCondition"
					}
				],
				"actions": [
					{
						"id": "sentry.integrations.slack.notify_action.SlackNotifyServiceAction",
						"name": "Send a notification to the Dummy Slack workspace to #dummy-channel and show tags [environment] in notification",
						"tags": "environment",
						"channel_id": "XX00X0X0X",
						"workspace": "1234",
						"channel": "#dummy-channel"
					}
				],
				"dateCreated": "2019-08-24T18:12:16.321Z"
			}
		}`)
	})
	mux.HandleFunc("/api/0/projects/the-interstellar-jurisdiction/pump-station/rules/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertPostJSONValue(t, map[string]interface{}{
			"actionMatch": "all",
			"environment": "production",
			"frequency":   30,
			"name":        "Notify errors",
			"conditions": []map[string]interface{}{
				{
					"interval": "1h",
					"name":     "The issue is seen more than 10 times in 1h",
					"value":    10,
					"id":       "sentry.rules.conditions.event_frequency.EventFrequencyCondition",
				},
			},
			"actions": []map[string]interface{}{
				{
					"id":         "sentry.integrations.slack.notify_action.SlackNotifyServiceAction",
					"name":       "Send a notification to the Dummy Slack workspace to #dummy-channel and show tags [environment] in notification",
					"tags":       "environment",
					"channel":    "#dummy-channel",
					"channel_id": "XX00X0X0X",
					"workspace":  "1234",
				},
			},
		}, r)

		w.WriteHeader(http.StatusAccepted)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"uuid": "fakeuuid"}`)

	})

	params := &CreateRuleParams{
		ActionMatch: "all",
		Environment: "production",
		Frequency:   30,
		Name:        "Notify errors",
		Conditions: []ConditionType{
			{
				"interval": "1h",
				"name":     "The issue is seen more than 10 times in 1h",
				"value":    float64(10),
				"id":       "sentry.rules.conditions.event_frequency.EventFrequencyCondition",
			},
		},
		Actions: []ActionType{
			{
				"id":         "sentry.integrations.slack.notify_action.SlackNotifyServiceAction",
				"name":       "Send a notification to the Dummy Slack workspace to #dummy-channel and show tags [environment] in notification",
				"tags":       "environment",
				"channel_id": "XX00X0X0X",
				"workspace":  "1234",
				"channel":    "#dummy-channel",
			},
		},
	}

	client := NewClient(httpClient, nil, "")
	rule, _, err := client.Rules.Create("the-interstellar-jurisdiction", "pump-station", params)
	require.NoError(t, err)

	environment := "production"
	expected := &Rule{
		ID:          "123456",
		ActionMatch: "all",
		Environment: &environment,
		Frequency:   30,
		Name:        "Notify errors",
		Conditions: []ConditionType{
			{
				"interval": "1h",
				"name":     "The issue is seen more than 10 times in 1h",
				"value":    float64(10),
				"id":       "sentry.rules.conditions.event_frequency.EventFrequencyCondition",
			},
		},
		Actions: []ActionType{
			{
				"id":         "sentry.integrations.slack.notify_action.SlackNotifyServiceAction",
				"name":       "Send a notification to the Dummy Slack workspace to #dummy-channel and show tags [environment] in notification",
				"tags":       "environment",
				"channel_id": "XX00X0X0X",
				"channel":    "#dummy-channel",
				"workspace":  "1234",
			},
		},
		Created: mustParseTime("2019-08-24T18:12:16.321Z"),
	}
	require.Equal(t, expected, rule)

}

func TestRulesService_Update(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	environment := "staging"
	params := &Rule{
		ID:          "12345",
		ActionMatch: "all",
		FilterMatch: "any",
		Environment: &environment,
		Frequency:   30,
		Name:        "Notify errors",
		Conditions: []ConditionType{
			{
				"id":       "sentry.rules.conditions.event_frequency.EventFrequencyCondition",
				"value":    500,
				"interval": "1h",
			},
		},
		Actions: []ActionType{
			{
				"id":         "sentry.integrations.slack.notify_action.SlackNotifyServiceAction",
				"name":       "Send a notification to the Dummy Slack workspace to #dummy-channel and show tags [environment] in notification",
				"tags":       "environment",
				"channel_id": "XX00X0X0X",
				"channel":    "#dummy-channel",
				"workspace":  "1234",
			},
		},
		Filters: []FilterType{
			{
				"id":    "sentry.rules.filters.issue_occurrences.IssueOccurrencesFilter",
				"name":  "The issue has happened at least 4 times",
				"value": 4,
			},
			{
				"attribute": "message",
				"id":        "sentry.rules.filters.event_attribute.EventAttributeFilter",
				"match":     "eq",
				"name":      "The event's message value equals test",
				"value":     "test",
			},
		},
		Created: mustParseTime("2019-08-24T18:12:16.321Z"),
	}

	mux.HandleFunc("/api/0/projects/the-interstellar-jurisdiction/pump-station/rules/12345/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "PUT", r)
		assertPostJSONValue(t, map[string]interface{}{
			"id":          "12345",
			"actionMatch": "all",
			"filterMatch": "any",
			"environment": "staging",
			"frequency":   json.Number("30"),
			"name":        "Notify errors",
			"dateCreated": "2019-08-24T18:12:16.321Z",
			"conditions": []map[string]interface{}{
				{
					"id":       "sentry.rules.conditions.event_frequency.EventFrequencyCondition",
					"value":    json.Number("500"),
					"interval": "1h",
				},
			},
			"actions": []map[string]interface{}{
				{
					"id":         "sentry.integrations.slack.notify_action.SlackNotifyServiceAction",
					"name":       "Send a notification to the Dummy Slack workspace to #dummy-channel and show tags [environment] in notification",
					"tags":       "environment",
					"channel":    "#dummy-channel",
					"channel_id": "XX00X0X0X",
					"workspace":  "1234",
				},
			},
			"filters": []map[string]interface{}{
				{
					"id":    "sentry.rules.filters.issue_occurrences.IssueOccurrencesFilter",
					"name":  "The issue has happened at least 4 times",
					"value": json.Number("4"),
				},
				{
					"attribute": "message",
					"id":        "sentry.rules.filters.event_attribute.EventAttributeFilter",
					"match":     "eq",
					"name":      "The event's message value equals test",
					"value":     "test",
				},
			},
		}, r)

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
			"environment": "staging",
			"actionMatch": "any",
			"frequency": 30,
			"name": "Notify errors",
			"conditions": [
				{
					"id": "sentry.rules.conditions.first_seen_event.FirstSeenEventCondition",
					"name": "An issue is first seen"
				}
			],
			"id": "12345",
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
		}`)
	})

	client := NewClient(httpClient, nil, "")
	rules, _, err := client.Rules.Update("the-interstellar-jurisdiction", "pump-station", "12345", params)
	assert.NoError(t, err)

	expected := &Rule{
		ID:          "12345",
		ActionMatch: "any",
		Environment: &environment,
		Frequency:   30,
		Name:        "Notify errors",
		Conditions: []ConditionType{
			{
				"id":   "sentry.rules.conditions.first_seen_event.FirstSeenEventCondition",
				"name": "An issue is first seen",
			},
		},
		Actions: []ActionType{
			{
				"id":         "sentry.integrations.slack.notify_action.SlackNotifyServiceAction",
				"name":       "Send a notification to the Dummy Slack workspace to #dummy-channel and show tags [environment] in notification",
				"tags":       "environment",
				"channel_id": "XX00X0X0X",
				"channel":    "#dummy-channel",
				"workspace":  "1234",
			},
		},
		Created: mustParseTime("2019-08-24T18:12:16.321Z"),
	}
	require.Equal(t, expected, rules)

}

func TestRuleService_Delete(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/0/projects/the-interstellar-jurisdiction/pump-station/rules/12345/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "DELETE", r)
	})

	client := NewClient(httpClient, nil, "")
	_, err := client.Rules.Delete("the-interstellar-jurisdiction", "pump-station", "12345")
	require.NoError(t, err)
}
