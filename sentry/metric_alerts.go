package sentry

import (
	"context"
	"fmt"
	"time"
)

type MetricAlertsService service

type MetricAlert struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	Environment      *string   `json:"environment,omitempty"`
	DataSet          string    `json:"dataset"`
	Query            string    `json:"query"`
	Aggregate        string    `json:"aggregate"`
	TimeWindow       float64   `json:"timeWindow"`
	ThresholdType    int       `json:"thresholdType"`
	ResolveThreshold float64   `json:"resolveThreshold"`
	Triggers         []Trigger `json:"triggers"`
	Projects         []string  `json:"projects"`
	Owner            string    `json:"owner"`
	Created          time.Time `json:"dateCreated"`
}

type Trigger map[string]interface{}

// List Alert Rules configured for a project
func (s *MetricAlertsService) List(ctx context.Context, organizationSlug string, projectSlug string) ([]*MetricAlert, *Response, error) {
	u := fmt.Sprintf("0/projects/%v/%v/alert-rules/", organizationSlug, projectSlug)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	metricAlerts := []*MetricAlert{}
	resp, err := s.client.Do(ctx, req, &metricAlerts)
	if err != nil {
		return nil, resp, err
	}
	return metricAlerts, resp, nil
}

type CreateAlertRuleParams struct {
	Name             string    `json:"name"`
	Environment      *string   `json:"environment,omitempty"`
	DataSet          string    `json:"dataset"`
	Query            string    `json:"query"`
	Aggregate        string    `json:"aggregate"`
	TimeWindow       float64   `json:"timeWindow"`
	ThresholdType    int       `json:"thresholdType"`
	ResolveThreshold float64   `json:"resolveThreshold"`
	Triggers         []Trigger `json:"triggers"`
	Projects         []string  `json:"projects"`
	Owner            string    `json:"owner"`
}

// Create a new Alert Rule bound to a project.
func (s *MetricAlertsService) Create(ctx context.Context, organizationSlug string, projectSlug string, params *CreateAlertRuleParams) (*MetricAlert, *Response, error) {
	u := fmt.Sprintf("0/projects/%v/%v/alert-rules/", organizationSlug, projectSlug)
	req, err := s.client.NewRequest("POST", u, params)
	if err != nil {
		return nil, nil, err
	}

	alertRule := new(MetricAlert)
	resp, err := s.client.Do(ctx, req, alertRule)
	if err != nil {
		return nil, resp, err
	}
	return alertRule, resp, nil
}

// Update an Alert Rule.
func (s *MetricAlertsService) Update(ctx context.Context, organizationSlug string, projectSlug string, alertRuleID string, params *MetricAlert) (*MetricAlert, *Response, error) {
	u := fmt.Sprintf("0/projects/%v/%v/alert-rules/%v/", organizationSlug, projectSlug, alertRuleID)
	req, err := s.client.NewRequest("PUT", u, params)
	if err != nil {
		return nil, nil, err
	}

	alertRule := new(MetricAlert)
	resp, err := s.client.Do(ctx, req, alertRule)
	if err != nil {
		return nil, resp, err
	}
	return alertRule, resp, nil
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
