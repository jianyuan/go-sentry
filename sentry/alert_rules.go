package sentry

import (
	"net/http"
	"time"

	"github.com/dghubble/sling"
)

type AlertRuleService struct {
	sling *sling.Sling
}

type AlertRule struct {
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

func newAlertRuleService(sling *sling.Sling) *AlertRuleService {
	return &AlertRuleService{
		sling: sling,
	}
}

// List Alert Rules configured for a project
func (s *AlertRuleService) List(organizationSlug string, projectSlug string) ([]AlertRule, *http.Response, error) {
	alertRules := new([]AlertRule)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("projects/"+organizationSlug+"/"+projectSlug+"/alert-rules/").Receive(alertRules, apiError)
	return *alertRules, resp, relevantError(err, *apiError)
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
func (s *AlertRuleService) Create(organizationSlug string, projectSlug string, params *CreateAlertRuleParams) (*AlertRule, *http.Response, error) {
	alertRule := new(AlertRule)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("projects/"+organizationSlug+"/"+projectSlug+"/alert-rules/").BodyJSON(params).Receive(alertRule, apiError)
	return alertRule, resp, relevantError(err, *apiError)
}

// Update an Alert Rule.
func (s *AlertRuleService) Update(organizationSlug string, projectSlug string, alertRuleID string, params *AlertRule) (*AlertRule, *http.Response, error) {
	alertRule := new(AlertRule)
	apiError := new(APIError)
	resp, err := s.sling.New().Put("projects/"+organizationSlug+"/"+projectSlug+"/alert-rules/"+alertRuleID+"/").BodyJSON(params).Receive(alertRule, apiError)
	return alertRule, resp, relevantError(err, *apiError)
}

// Delete an Alert Rule.
func (s *AlertRuleService) Delete(organizationSlug string, projectSlug string, alertRuleID string) (*http.Response, error) {
	apiError := new(APIError)
	resp, err := s.sling.New().Delete("projects/"+organizationSlug+"/"+projectSlug+"/alert-rules/"+alertRuleID+"/").Receive(nil, apiError)
	return resp, relevantError(err, *apiError)
}
