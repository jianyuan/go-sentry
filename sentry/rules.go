package sentry

import (
	"errors"
	"net/http"
	"time"

	"github.com/dghubble/sling"
)

// Rule represents an alert rule configured for this project.
// https://github.com/getsentry/sentry/blob/9.0.0/src/sentry/api/serializers/models/rule.py
type Rule struct {
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

// RuleTaskDetail represents the inline struct Sentry defines for task details
// https://github.com/getsentry/sentry/blob/7ce8f5a4bbc3429eef4b2e273148baf6525fede2/src/sentry/api/endpoints/project_rule_task_details.py#L29
type RuleTaskDetail struct {
	Status string `json:"status"`
	Rule   Rule   `json:"rule"`
	Error  string `json:"error"`
}

// RuleService provides methods for accessing Sentry project
// client key API endpoints.
// https://docs.sentry.io/api/projects/
type RuleService struct {
	sling *sling.Sling
}

func newRuleService(sling *sling.Sling) *RuleService {
	return &RuleService{
		sling: sling,
	}
}

// List alert rules configured for a project.
func (s *RuleService) List(organizationSlug string, projectSlug string) ([]Rule, *http.Response, error) {
	rules := new([]Rule)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("projects/"+organizationSlug+"/"+projectSlug+"/rules/").Receive(rules, apiError)
	return *rules, resp, relevantError(err, *apiError)
}

// ConditionType for defining conditions.
type ConditionType map[string]interface{}

// ActionType for defining actions.
type ActionType map[string]interface{}

// FilterType for defining actions.
type FilterType map[string]interface{}

// CreateRuleParams are the parameters for RuleService.Create.
type CreateRuleParams struct {
	ActionMatch string          `json:"actionMatch"`
	FilterMatch string          `json:"filterMatch"`
	Environment string          `json:"environment,omitempty"`
	Frequency   int             `json:"frequency"`
	Name        string          `json:"name"`
	Conditions  []ConditionType `json:"conditions"`
	Actions     []ActionType    `json:"actions"`
	Filters     []FilterType    `json:"filters"`
}

// CreateRuleActionParams models the actions when creating the action for the rule.
type CreateRuleActionParams struct {
	ID        string `json:"id"`
	Tags      string `json:"tags"`
	Channel   string `json:"channel"`
	Workspace string `json:"workspace"`
}

// CreateRuleConditionParams models the conditions when creating the action for the rule.
type CreateRuleConditionParams struct {
	ID       string `json:"id"`
	Interval string `json:"interval"`
	Value    int    `json:"value"`
	Level    int    `json:"level"`
	Match    string `json:"match"`
}

// Create a new alert rule bound to a project.
func (s *RuleService) Create(organizationSlug string, projectSlug string, params *CreateRuleParams) (*Rule, *http.Response, error) {
	rule := new(Rule)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("projects/"+organizationSlug+"/"+projectSlug+"/rules/").BodyJSON(params).Receive(rule, apiError)
	if resp.StatusCode == 202 {
		// We just received a reference to an async task, we need to check another endpoint to retrieve the rule we created
		return s.getRuleFromTaskDetail(organizationSlug, projectSlug, rule.TaskUUID)
	}
	return rule, resp, relevantError(err, *apiError)
}

// getRuleFromTaskDetail is called when Sentry offloads the rule creation process to an async task and sends us back the task's uuid.
// It usually doesn't happen, but when creating Slack notification rules, it seemed to be sometimes the case. During testing it
// took very long for a task to finish (10+ seconds) which is why this method can take long to return.
func (s *RuleService) getRuleFromTaskDetail(organizationSlug string, projectSlug string, taskUuid string) (*Rule, *http.Response, error) {
	taskDetail := &RuleTaskDetail{}
	var resp *http.Response
	for i := 0; i < 5; i++ {
		time.Sleep(5 * time.Second)
		resp, err := s.sling.New().Get("projects/" + organizationSlug + "/" + projectSlug + "/rule-task/" + taskUuid + "/").ReceiveSuccess(taskDetail)
		if taskDetail.Status == "success" {
			return &taskDetail.Rule, resp, err
		} else if taskDetail.Status == "failed" {
			return &taskDetail.Rule, resp, errors.New(taskDetail.Error)
		} else if resp.StatusCode == 404 {
			return &Rule{}, resp, errors.New("couldn't find the rule creation task for uuid '" + taskUuid + "' in Sentry (HTTP 404)")
		}
	}
	return &Rule{}, resp, errors.New("getting the status of the rule creation from Sentry took too long")
}

// Update a rule.
func (s *RuleService) Update(organizationSlug string, projectSlug string, ruleID string, params *Rule) (*Rule, *http.Response, error) {
	rule := new(Rule)
	apiError := new(APIError)
	resp, err := s.sling.New().Put("projects/"+organizationSlug+"/"+projectSlug+"/rules/"+ruleID+"/").BodyJSON(params).Receive(rule, apiError)
	return rule, resp, relevantError(err, *apiError)
}

// Delete a rule.
func (s *RuleService) Delete(organizationSlug string, projectSlug string, ruleID string) (*http.Response, error) {
	apiError := new(APIError)
	resp, err := s.sling.New().Delete("projects/"+organizationSlug+"/"+projectSlug+"/rules/"+ruleID+"/").Receive(nil, apiError)
	return resp, relevantError(err, *apiError)
}
