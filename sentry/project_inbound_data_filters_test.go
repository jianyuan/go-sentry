package sentry

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectInboundDataFiltersService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/0/projects/organization_slug/project_slug/filters/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, http.MethodGet, r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[
			{
				"id": "browser-extensions",
				"active": false
			},
			{
				"id": "filtered-transaction",
				"active": true
			},
			{
				"id": "legacy-browsers",
				"active": [
					"ie_pre_9"
				]
			},
			{
				"id": "localhost",
				"active": false
			},
			{
				"id": "web-crawlers",
				"active": false
			}
		]`)
	})

	ctx := context.Background()
	filters, _, err := client.ProjectInboundDataFilters.List(ctx, "organization_slug", "project_slug")
	assert.NoError(t, err)

	expected := []*ProjectInboundDataFilter{
		{
			ID:     "browser-extensions",
			Active: BoolOrStringSlice{IsBool: true, BoolVal: false},
		},
		{
			ID:     "filtered-transaction",
			Active: BoolOrStringSlice{IsBool: true, BoolVal: true},
		},
		{
			ID:     "legacy-browsers",
			Active: BoolOrStringSlice{IsStringSlice: true, StringSliceVal: []string{"ie_pre_9"}},
		},
		{
			ID:     "localhost",
			Active: BoolOrStringSlice{IsBool: true, BoolVal: false},
		},
		{
			ID:     "web-crawlers",
			Active: BoolOrStringSlice{IsBool: true, BoolVal: false},
		},
	}
	assert.Equal(t, expected, filters)
}

func TestProjectInboundDataFiltersService_UpdateActive(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/0/projects/organization_slug/project_slug/filters/filter_id/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, http.MethodPut, r)
		assertPostJSON(t, map[string]interface{}{
			"active": true,
		}, r)
	})

	ctx := context.Background()
	params := &UpdateProjectInboundDataFilterParams{
		Active: Bool(true),
	}
	_, err := client.ProjectInboundDataFilters.Update(ctx, "organization_slug", "project_slug", "filter_id", params)
	assert.NoError(t, err)
}

func TestProjectInboundDataFiltersService_UpdateSubfilters(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/0/projects/organization_slug/project_slug/filters/filter_id/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, http.MethodPut, r)
		assertPostJSON(t, map[string]interface{}{
			"subfilters": []interface{}{"ie_pre_9"},
		}, r)
	})

	ctx := context.Background()
	params := &UpdateProjectInboundDataFilterParams{
		Subfilters: []string{"ie_pre_9"},
	}
	_, err := client.ProjectInboundDataFilters.Update(ctx, "organization_slug", "project_slug", "filter_id", params)
	assert.NoError(t, err)
}
