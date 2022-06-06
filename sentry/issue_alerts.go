package sentry

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// IssueAlert represents an issue alert configured for this project.
// https://github.com/getsentry/sentry/blob/22.5.0/src/sentry/api/serializers/models/rule.py#L131-L155
type IssueAlert struct {
	ID          string          `json:"id"`
	ActionMatch string          `json:"actionMatch"`
	FilterMatch string          `json:"filterMatch"`
	Environment *string         `json:"environment,omitempty"`
	Frequency   int             `json:"frequency"`
	Name        string          `json:"name"`
	Conditions  []ConditionType `json:"conditions"`
	Actions     []ActionType    `json:"actions"`
	Filters     []FilterType    `json:"filters"`
	Created     time.Time       `json:"dateCreated"`
	TaskUUID    string          `json:"uuid,omitempty"` // This is actually the UUID of the async task that can be spawned to create the rule
}

// IssueAlertTaskDetail represents the inline struct Sentry defines for task details
// https://github.com/getsentry/sentry/blob/22.5.0/src/sentry/api/endpoints/project_rule_task_details.py#L29
type IssueAlertTaskDetail struct {
	Status string     `json:"status"`
	Rule   IssueAlert `json:"rule"`
	Error  string     `json:"error"`
}

// IssueAlertsService provides methods for accessing Sentry project
// client key API endpoints.
// https://docs.sentry.io/api/projects/
type IssueAlertsService service

// List issue alerts configured for a project.
func (s *IssueAlertsService) List(ctx context.Context, organizationSlug string, projectSlug string) ([]*IssueAlert, *Response, error) {
	u := fmt.Sprintf("0/projects/%v/%v/rules/", organizationSlug, projectSlug)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	alerts := []*IssueAlert{}
	resp, err := s.client.Do(ctx, req, &alerts)
	if err != nil {
		return nil, resp, err
	}
	return alerts, resp, nil
}

// ConditionType for defining conditions.
type ConditionType map[string]interface{}

// ActionType for defining actions.
type ActionType map[string]interface{}

// FilterType for defining actions.
type FilterType map[string]interface{}

// CreateIssueAlertParams are the parameters for IssueAlertsService.Create.
type CreateIssueAlertParams struct {
	ActionMatch string          `json:"actionMatch"`
	FilterMatch string          `json:"filterMatch"`
	Environment string          `json:"environment,omitempty"`
	Frequency   int             `json:"frequency"`
	Name        string          `json:"name"`
	Conditions  []ConditionType `json:"conditions"`
	Actions     []ActionType    `json:"actions"`
	Filters     []FilterType    `json:"filters"`
}

// CreateIssueAlertActionParams models the actions when creating the action for the rule.
type CreateIssueAlertActionParams struct {
	ID        string `json:"id"`
	Tags      string `json:"tags"`
	Channel   string `json:"channel"`
	Workspace string `json:"workspace"`
}

// CreateIssueAlertConditionParams models the conditions when creating the action for the rule.
type CreateIssueAlertConditionParams struct {
	ID       string `json:"id"`
	Interval string `json:"interval"`
	Value    int    `json:"value"`
	Level    int    `json:"level"`
	Match    string `json:"match"`
}

// Create a new issue alert bound to a project.
func (s *IssueAlertsService) Create(ctx context.Context, organizationSlug string, projectSlug string, params *CreateIssueAlertParams) (*IssueAlert, *Response, error) {
	u := fmt.Sprintf("0/projects/%v/%v/rules/", organizationSlug, projectSlug)
	req, err := s.client.NewRequest("POST", u, params)
	if err != nil {
		return nil, nil, err
	}

	alert := new(IssueAlert)
	resp, err := s.client.Do(ctx, req, alert)
	if err != nil {
		return nil, resp, err
	}

	if resp.StatusCode == 202 {
		// We just received a reference to an async task, we need to check another endpoint to retrieve the issue alert we created
		return s.getIssueAlertFromTaskDetail(ctx, organizationSlug, projectSlug, alert.TaskUUID)
	}

	return alert, resp, nil
}

// getIssueAlertFromTaskDetail is called when Sentry offloads the issue alert creation process to an async task and sends us back the task's uuid.
// It usually doesn't happen, but when creating Slack notification rules, it seemed to be sometimes the case. During testing it
// took very long for a task to finish (10+ seconds) which is why this method can take long to return.
func (s *IssueAlertsService) getIssueAlertFromTaskDetail(ctx context.Context, organizationSlug string, projectSlug string, taskUuid string) (*IssueAlert, *Response, error) {
	u := fmt.Sprintf("0/projects/%v/%v/rule-task/%v/", organizationSlug, projectSlug, taskUuid)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var resp *Response
	for i := 0; i < 5; i++ {
		// TODO: Read poll interval from context
		time.Sleep(5 * time.Second)

		taskDetail := new(IssueAlertTaskDetail)
		resp, err := s.client.Do(ctx, req, taskDetail)
		if err != nil {
			return nil, resp, err
		}

		if taskDetail.Status == "success" {
			return &taskDetail.Rule, resp, err
		} else if taskDetail.Status == "failed" {
			return &taskDetail.Rule, resp, errors.New(taskDetail.Error)
		} else if resp.StatusCode == 404 {
			return nil, resp, fmt.Errorf("couldn't find the issue alert creation task for uuid %v in Sentry (HTTP 404)", taskUuid)
		}
	}
	return nil, resp, errors.New("getting the status of the issue alert creation from Sentry took too long")
}

// Update an issue alert.
func (s *IssueAlertsService) Update(ctx context.Context, organizationSlug string, projectSlug string, issueAlertID string, params *IssueAlert) (*IssueAlert, *Response, error) {
	u := fmt.Sprintf("0/projects/%v/%v/rules/%v/", organizationSlug, projectSlug, issueAlertID)
	req, err := s.client.NewRequest("PUT", u, params)
	if err != nil {
		return nil, nil, err
	}

	alert := new(IssueAlert)
	resp, err := s.client.Do(ctx, req, alert)
	if err != nil {
		return nil, resp, err
	}
	return alert, resp, nil
}

// Delete an issue alert.
func (s *IssueAlertsService) Delete(ctx context.Context, organizationSlug string, projectSlug string, issueAlertID string) (*Response, error) {
	u := fmt.Sprintf("0/projects/%v/%v/rules/%v/", organizationSlug, projectSlug, issueAlertID)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
