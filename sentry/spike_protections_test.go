package sentry

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpikeProtectionsService_Enable(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/0/organizations/organization_slug/spike-protections/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, http.MethodPost, r)
		assertPostJSON(t, map[string]interface{}{
			"projects": []interface{}{"$all"},
		}, r)
	})

	params := &SpikeProtectionParams{
		Projects: []string{"$all"},
	}
	ctx := context.Background()
	_, err := client.SpikeProtections.Enable(ctx, "organization_slug", params)
	assert.NoError(t, err)
}

func TestSpikeProtectionsService_Disable(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/0/organizations/organization_slug/spike-protections/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, http.MethodDelete, r)
		assertPostJSON(t, map[string]interface{}{
			"projects": []interface{}{"$all"},
		}, r)
	})

	params := &SpikeProtectionParams{
		Projects: []string{"$all"},
	}
	ctx := context.Background()
	_, err := client.SpikeProtections.Disable(ctx, "organization_slug", params)
	assert.NoError(t, err)
}
