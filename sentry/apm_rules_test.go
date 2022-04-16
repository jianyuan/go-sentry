package sentry

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAPMRuleService_List(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/0/projects/the-interstellar-jurisdiction/pump-station/alert-rules/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `[
			{
				"id": "12345",
				"name": "pump-station-alert",
				"environment": "production",
				"dataset": "transactions",
				"query": "http.url:http://service/unreadmessages",
				"aggregate": "p50(transaction.duration)",
				"thresholdType": 0,
				"resolveThreshold": 100.0,
				"timeWindow": 5.0,
				"triggers": [
					{
						"id": "6789",
						"alertRuleId": "12345",
						"label": "critical",
						"thresholdType": 0,
						"alertThreshold": 55501.0,
						"resolveThreshold": 100.0,
						"dateCreated": "2022-04-07T16:46:48.607583Z",
						"actions": [
							{
								"id": "12345",
								"alertRuleTriggerId": "12345",
								"type": "slack",
								"targetType": "specific",
								"targetIdentifier": "#apm-alerts",
								"inputChannelId": "C038NF00X4F",
								"integrationId": 123,
								"sentryAppId": null,
								"dateCreated": "2022-04-07T16:46:49.154638Z",
								"desc": "Send a Slack notification to #apm-alerts"
							}
						]
					}
				],
				"projects": [
					"pump-station"
				],
				"owner": "pump-station:12345",
				"dateCreated": "2022-04-07T16:46:48.569571Z"
			}
		]`)
	})

	client := NewClient(httpClient, nil, "")
	apmRules, _, err := client.APMRules.List("the-interstellar-jurisdiction", "pump-station")
	require.NoError(t, err)

	environment := "production"
	expected := []APMRule{
		{
			ID:               "12345",
			Name:             "pump-station-alert",
			Environment:      &environment,
			DataSet:          "transactions",
			Query:            "http.url:http://service/unreadmessages",
			Aggregate:        "p50(transaction.duration)",
			ThresholdType:    int(0),
			ResolveThreshold: float64(100.0),
			TimeWindow:       float64(5.0),
			Triggers: []Trigger{
				{
					"id":               "6789",
					"alertRuleId":      "12345",
					"label":            "critical",
					"thresholdType":    float64(0),
					"alertThreshold":   float64(55501.0),
					"resolveThreshold": float64(100.0),
					"dateCreated":      "2022-04-07T16:46:48.607583Z",
					"actions": []interface{}{map[string]interface{}{
						"id":                 "12345",
						"alertRuleTriggerId": "12345",
						"type":               "slack",
						"targetType":         "specific",
						"targetIdentifier":   "#apm-alerts",
						"inputChannelId":     "C038NF00X4F",
						"integrationId":      float64(123),
						"sentryAppId":        interface{}(nil),
						"dateCreated":        "2022-04-07T16:46:49.154638Z",
						"desc":               "Send a Slack notification to #apm-alerts",
					},
					},
				},
			},
			Projects: []string{"pump-station"},
			Owner:    "pump-station:12345",
			Created:  mustParseTime("2022-04-07T16:46:48.569571Z"),
		},
	}
	require.Equal(t, expected, apmRules)
}

func TestAPMRuleService_Create(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/0/projects/the-interstellar-jurisdiction/pump-station/alert-rules/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `
			{
				"id": "12345",
				"name": "pump-station-alert",
				"environment": "production",
				"dataset": "transactions",
				"query": "http.url:http://service/unreadmessages",
				"aggregate": "p50(transaction.duration)",
				"timeWindow": 10,
				"thresholdType": 0,
				"resolveThreshold": 0,
				"triggers": [
				  {
					"actions": [
					  {
						"alertRuleTriggerId": "56789",
						"dateCreated": "2022-04-15T15:06:01.087054Z",
						"desc": "Send a Slack notification to #apm-alerts",
						"id": "12389",
						"inputChannelId": "C0XXXFKLXXX",
						"integrationId": 111,
						"sentryAppId": null,
						"targetIdentifier": "#apm-alerts",
						"targetType": "specific",
						"type": "slack"
					  }
					],
					"alertRuleId": "12345",
					"alertThreshold": 10000,
					"dateCreated": "2022-04-15T15:06:01.079598Z",
					"id": "56789",
					"label": "critical",
					"resolveThreshold": 0,
					"thresholdType": 0
				  }
				],
				"projects": [
				  "pump-station"
				],
				"owner": "pump-station:12345",
				"dateCreated": "2022-04-15T15:06:01.05618Z"
			}
		`)
	})

	client := NewClient(httpClient, nil, "")
	environment := "production"
	params := CreateAPMRuleParams{
		Name:             "pump-station-alert",
		Environment:      &environment,
		DataSet:          "transactions",
		Query:            "http.url:http://service/unreadmessages",
		Aggregate:        "p50(transaction.duration)",
		TimeWindow:       10.0,
		ThresholdType:    0,
		ResolveThreshold: 0,
		Triggers: []Trigger{{
			"actions": []interface{}{map[string]interface{}{
				"type":             "slack",
				"targetType":       "specific",
				"targetIdentifier": "#apm-alerts",
				"inputChannelId":   "C0XXXFKLXXX",
				"integrationId":    111,
			},
			},
			"alertThreshold":   10000,
			"label":            "critical",
			"resolveThreshold": 0,
			"thresholdType":    0,
		}},
		Projects: []string{"pump-station"},
		Owner:    "pump-station:12345",
	}
	apmRule, _, err := client.APMRules.Create("the-interstellar-jurisdiction", "pump-station", &params)
	require.NoError(t, err)

	expected := &APMRule{
		ID:               "12345",
		Name:             "pump-station-alert",
		Environment:      &environment,
		DataSet:          "transactions",
		Query:            "http.url:http://service/unreadmessages",
		Aggregate:        "p50(transaction.duration)",
		ThresholdType:    int(0),
		ResolveThreshold: float64(0),
		TimeWindow:       float64(10.0),
		Triggers: []Trigger{
			{
				"id":               "56789",
				"alertRuleId":      "12345",
				"label":            "critical",
				"thresholdType":    float64(0),
				"alertThreshold":   float64(10000),
				"resolveThreshold": float64(0),
				"dateCreated":      "2022-04-15T15:06:01.079598Z",
				"actions": []interface{}{map[string]interface{}{
					"id":                 "12389",
					"alertRuleTriggerId": "56789",
					"type":               "slack",
					"targetType":         "specific",
					"targetIdentifier":   "#apm-alerts",
					"inputChannelId":     "C0XXXFKLXXX",
					"integrationId":      float64(111),
					"sentryAppId":        interface{}(nil),
					"dateCreated":        "2022-04-15T15:06:01.087054Z",
					"desc":               "Send a Slack notification to #apm-alerts",
				},
				},
			},
		},
		Projects: []string{"pump-station"},
		Owner:    "pump-station:12345",
		Created:  mustParseTime("2022-04-15T15:06:01.05618Z"),
	}

	require.Equal(t, expected, apmRule)
}

func TestAPMRuleService_Update(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	environment := "production"
	params := &APMRule{
		ID:               "12345",
		Name:             "pump-station-alert",
		Environment:      &environment,
		DataSet:          "transactions",
		Query:            "http.url:http://service/unreadmessages",
		Aggregate:        "p50(transaction.duration)",
		TimeWindow:       10,
		ThresholdType:    0,
		ResolveThreshold: 0,
		Triggers: []Trigger{{
			"actions":          []interface{}{map[string]interface{}{}},
			"alertRuleId":      "12345",
			"alertThreshold":   10000,
			"dateCreated":      "2022-04-15T15:06:01.079598Z",
			"id":               "56789",
			"label":            "critical",
			"resolveThreshold": 0,
			"thresholdType":    0,
		}},
		Owner:   "pump-station:12345",
		Created: mustParseTime("2022-04-15T15:06:01.079598Z"),
	}

	mux.HandleFunc("/api/0/projects/the-interstellar-jurisdiction/pump-station/alert-rules/12345/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "PUT", r)
		assertPostJSON(t, map[string]interface{}{
			"id":               "12345",
			"name":             "pump-station-alert",
			"environment":      environment,
			"dataset":          "transactions",
			"query":            "http.url:http://service/unreadmessages",
			"aggregate":        "p50(transaction.duration)",
			"timeWindow":       json.Number("10"),
			"thresholdType":    json.Number("0"),
			"resolveThreshold": json.Number("0"),
			"triggers": []interface{}{
				map[string]interface{}{
					"actions":          []interface{}{map[string]interface{}{}},
					"alertRuleId":      "12345",
					"alertThreshold":   json.Number("10000"),
					"dateCreated":      "2022-04-15T15:06:01.079598Z",
					"id":               "56789",
					"label":            "critical",
					"resolveThreshold": json.Number("0"),
					"thresholdType":    json.Number("0"),
				},
			},
			"projects":    interface{}(nil),
			"owner":       "pump-station:12345",
			"dateCreated": "2022-04-15T15:06:01.079598Z",
		}, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `
			{
				"id": "12345",
				"name": "pump-station-alert",
				"environment": "production",
				"dataset": "transactions",
				"query": "http.url:http://service/unreadmessages",
				"aggregate": "p50(transaction.duration)",
				"timeWindow": 10,
				"thresholdType": 0,
				"resolveThreshold": 0,
				"triggers": [
				  {
					"actions": [
						{
							"id":                 "12389",
							"alertRuleTriggerId": "56789",
							"type":               "slack",
							"targetType":         "specific",
							"targetIdentifier":   "#apm-alerts",
							"inputChannelId":     "C0XXXFKLXXX",
							"integrationId":      111,
							"sentryAppId":        null,
							"dateCreated":        "2022-04-15T15:06:01.087054Z",
							"desc":               "Send a Slack notification to #apm-alerts"
						}
					],
					"alertRuleId": "12345",
					"alertThreshold": 10000,
					"dateCreated": "2022-04-15T15:06:01.079598Z",
					"id": "56789",
					"label": "critical",
					"resolveThreshold": 0,
					"thresholdType": 0
				  }
				],
				"projects": [
				  "pump-station"
				],
				"owner": "pump-station:12345",
				"dateCreated": "2022-04-15T15:06:01.05618Z"
			}
		`)
	})

	client := NewClient(httpClient, nil, "")
	apmRule, _, err := client.APMRules.Update("the-interstellar-jurisdiction", "pump-station", "12345", params)
	assert.NoError(t, err)

	expected := &APMRule{
		ID:               "12345",
		Name:             "pump-station-alert",
		Environment:      &environment,
		DataSet:          "transactions",
		Query:            "http.url:http://service/unreadmessages",
		Aggregate:        "p50(transaction.duration)",
		ThresholdType:    int(0),
		ResolveThreshold: float64(0),
		TimeWindow:       float64(10.0),
		Triggers: []Trigger{
			{
				"id":               "56789",
				"alertRuleId":      "12345",
				"label":            "critical",
				"thresholdType":    float64(0),
				"alertThreshold":   float64(10000),
				"resolveThreshold": float64(0),
				"dateCreated":      "2022-04-15T15:06:01.079598Z",
				"actions": []interface{}{map[string]interface{}{
					"id":                 "12389",
					"alertRuleTriggerId": "56789",
					"type":               "slack",
					"targetType":         "specific",
					"targetIdentifier":   "#apm-alerts",
					"inputChannelId":     "C0XXXFKLXXX",
					"integrationId":      float64(111),
					"sentryAppId":        interface{}(nil),
					"dateCreated":        "2022-04-15T15:06:01.087054Z",
					"desc":               "Send a Slack notification to #apm-alerts",
				}},
			},
		},
		Projects: []string{"pump-station"},
		Owner:    "pump-station:12345",
		Created:  mustParseTime("2022-04-15T15:06:01.05618Z"),
	}

	require.Equal(t, expected, apmRule)
}

func TestAPMRuleService_Delete(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/api/0/projects/the-interstellar-jurisdiction/pump-station/alert-rules/12345/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "DELETE", r)
	})

	client := NewClient(httpClient, nil, "")
	_, err := client.APMRules.Delete("the-interstellar-jurisdiction", "pump-station", "12345")
	require.NoError(t, err)
}
