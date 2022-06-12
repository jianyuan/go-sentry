package sentry

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDashboardsService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/0/organizations/the-interstellar-jurisdiction/dashboards/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `[
			{
				"id": "11833",
				"title": "General",
				"dateCreated": "2022-06-07T16:48:26.255520Z"
			},

			{
				"id": "11832",
				"title": "Mobile Template",
				"dateCreated": "2022-06-07T16:43:40.456607Z"
			}
		]`)
	})

	ctx := context.Background()
	widgetErrors, _, err := client.Dashboards.List(ctx, "the-interstellar-jurisdiction", nil)

	expected := []*Dashboard{
		{
			ID:          String("11833"),
			Title:       String("General"),
			DateCreated: Time(mustParseTime("2022-06-07T16:48:26.255520Z")),
		},
		{
			ID:          String("11832"),
			Title:       String("Mobile Template"),
			DateCreated: Time(mustParseTime("2022-06-07T16:43:40.456607Z")),
		},
	}
	assert.Equal(t, expected, widgetErrors)
	assert.NoError(t, err)
}
