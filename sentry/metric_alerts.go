package sentry

import (
	"context"
	"fmt"
	"time"
)

type MetricAlertsService service

type MetricAlert struct {
	ID               *string               `json:"id"`
	Name             *string               `json:"name"`
	Environment      *string               `json:"environment,omitempty"`
	DataSet          *string               `json:"dataset"`
	Query            *string               `json:"query"`
	Aggregate        *string               `json:"aggregate"`
	TimeWindow       *float64              `json:"timeWindow"`
	ThresholdType    *int                  `json:"thresholdType"`
	ResolveThreshold *float64              `json:"resolveThreshold"`
	Triggers         []*MetricAlertTrigger `json:"triggers"`
	Projects         []string              `json:"projects"`
	Owner            *string               `json:"owner"`
	DateCreated      *time.Time            `json:"dateCreated"`
}

type MetricAlertTrigger map[string]interface{}

// List Alert Rules configured for a project
func (s *MetricAlertsService) List(ctx context.Context, organizationSlug string, projectSlug string) ([]*MetricAlert, *Response, error) {
	u := fmt.Sprintf("0/projects/%v/%v/alert-rules/", organizationSlug, projectSlug)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	alerts := []*MetricAlert{}
	resp, err := s.client.Do(ctx, req, &alerts)
	if err != nil {
		return nil, resp, err
	}
	return alerts, resp, nil
}

// Get details on an issue alert.
func (s *MetricAlertsService) Get(ctx context.Context, organizationSlug string, projectSlug string, id string) (*MetricAlert, *Response, error) {
	u := fmt.Sprintf("0/projects/%v/%v/alert-rules/%v/", organizationSlug, projectSlug, id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	alert := new(MetricAlert)
	resp, err := s.client.Do(ctx, req, alert)
	if err != nil {
		return nil, resp, err
	}
	return alert, resp, nil
}

// Create a new Alert Rule bound to a project.
func (s *MetricAlertsService) Create(ctx context.Context, organizationSlug string, projectSlug string, params *MetricAlert) (*MetricAlert, *Response, error) {
	u := fmt.Sprintf("0/projects/%v/%v/alert-rules/", organizationSlug, projectSlug)
	req, err := s.client.NewRequest("POST", u, params)
	if err != nil {
		return nil, nil, err
	}

	alert := new(MetricAlert)
	resp, err := s.client.Do(ctx, req, alert)
	if err != nil {
		return nil, resp, err
	}
	return alert, resp, nil
}

// Update an Alert Rule.
func (s *MetricAlertsService) Update(ctx context.Context, organizationSlug string, projectSlug string, alertRuleID string, params *MetricAlert) (*MetricAlert, *Response, error) {
	u := fmt.Sprintf("0/projects/%v/%v/alert-rules/%v/", organizationSlug, projectSlug, alertRuleID)
	req, err := s.client.NewRequest("PUT", u, params)
	if err != nil {
		return nil, nil, err
	}

	alert := new(MetricAlert)
	resp, err := s.client.Do(ctx, req, alert)
	if err != nil {
		return nil, resp, err
	}
	return alert, resp, nil
}

// Delete an Alert Rule.
func (s *MetricAlertsService) Delete(ctx context.Context, organizationSlug string, projectSlug string, alertRuleID string) (*Response, error) {
	u := fmt.Sprintf("0/projects/%v/%v/alert-rules/%v/", organizationSlug, projectSlug, alertRuleID)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
