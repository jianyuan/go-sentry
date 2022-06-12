package sentry

import (
	"context"
	"fmt"
	"time"
)

// Dashboard represents a Dashboard.
type Dashboard struct {
	ID          *string            `json:"id,omitempty"`
	Title       *string            `json:"title,omitempty"`
	DateCreated *time.Time         `json:"dateCreated,omitempty"`
	Widgets     []*DashboardWidget `json:"widgets,omitempty"`
}

// DashboardsService provides methods for accessing Sentry dashboard API endpoints.
type DashboardsService service

// List dashboards in an organization.
func (s *DashboardsService) List(ctx context.Context, organizationSlug string, params *ListCursorParams) ([]*Dashboard, *Response, error) {
	u := fmt.Sprintf("0/organizations/%v/dashboards/", organizationSlug)
	u, err := addQuery(u, params)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var dashboards []*Dashboard
	resp, err := s.client.Do(ctx, req, &dashboards)
	if err != nil {
		return nil, resp, err
	}
	return dashboards, resp, nil
}
