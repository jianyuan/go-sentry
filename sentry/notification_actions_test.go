package sentry

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotificationActionsService_Get(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/0/organizations/organization_slug/notifications/actions/action_id/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, http.MethodGet, r)
		fmt.Fprintf(w, `{
			"id": "836501735",
			"organizationId": "62848264",
			"serviceType": "sentry_notification",
			"targetDisplay": "default",
			"targetIdentifier": "default",
			"targetType": "specific",
			"triggerType": "spike-protection",
			"projects": [
				4505321021243392
			]
		}`)
	})

	ctx := context.Background()
	action, _, err := client.NotificationActions.Get(ctx, "organization_slug", "action_id")
	assert.NoError(t, err)

	expected := &NotificationAction{
		ID:               JsonNumber(json.Number("836501735")),
		TriggerType:      String("spike-protection"),
		ServiceType:      String("sentry_notification"),
		TargetIdentifier: "default",
		TargetDisplay:    String("default"),
		TargetType:       String("specific"),
		Projects:         []json.Number{"4505321021243392"},
	}
	assert.Equal(t, expected, action)
}

func TestNotificationActionsService_Create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/0/organizations/organization_slug/notifications/actions/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, http.MethodPost, r)
		assertPostJSON(t, map[string]interface{}{
			"projects":         []interface{}{"go"},
			"serviceType":      "sentry_notification",
			"targetDisplay":    "default",
			"targetIdentifier": "default",
			"targetType":       "specific",
			"triggerType":      "spike-protection",
		}, r)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `{
			"id": "836501735",
			"organizationId": "62848264",
			"serviceType": "sentry_notification",
			"targetDisplay": "default",
			"targetIdentifier": "default",
			"targetType": "specific",
			"triggerType": "spike-protection",
			"projects": [
				4505321021243392
			]
		}`)
	})

	params := &CreateNotificationActionParams{
		TriggerType:      String("spike-protection"),
		ServiceType:      String("sentry_notification"),
		TargetIdentifier: String("default"),
		TargetDisplay:    String("default"),
		TargetType:       String("specific"),
		Projects:         []string{"go"},
	}
	ctx := context.Background()
	action, _, err := client.NotificationActions.Create(ctx, "organization_slug", params)
	assert.NoError(t, err)

	expected := &NotificationAction{
		ID:               JsonNumber(json.Number("836501735")),
		TriggerType:      String("spike-protection"),
		ServiceType:      String("sentry_notification"),
		TargetIdentifier: "default",
		TargetDisplay:    String("default"),
		TargetType:       String("specific"),
		Projects:         []json.Number{"4505321021243392"},
	}
	assert.Equal(t, expected, action)
}

func TestNotificationActionsService_Delete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/0/organizations/organization_slug/notifications/actions/action_id/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, http.MethodDelete, r)
	})

	ctx := context.Background()
	_, err := client.NotificationActions.Delete(ctx, "organization_slug", "action_id")
	assert.NoError(t, err)
}
