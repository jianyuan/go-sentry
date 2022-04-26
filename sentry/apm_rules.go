package sentry

import (
	"net/http"
	"time"

	"github.com/dghubble/sling"
)

type APMRuleService struct {
	sling *sling.Sling
}

type APMRule struct {
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

func newAPMRuleService(sling *sling.Sling) *APMRuleService {
	return &APMRuleService{
		sling: sling,
	}
}

// List APM rules configured for a project
func (s *APMRuleService) List(organizationSlug string, projectSlug string) ([]APMRule, *http.Response, error) {
	apmRules := new([]APMRule)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("projects/"+organizationSlug+"/"+projectSlug+"/alert-rules/").Receive(apmRules, apiError)
	return *apmRules, resp, relevantError(err, *apiError)
}

type CreateAPMRuleParams struct {
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

// Create a new APM rule bound to a project.
func (s *APMRuleService) Create(organizationSlug string, projectSlug string, params *CreateAPMRuleParams) (*APMRule, *http.Response, error) {
	apmRule := new(APMRule)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("projects/"+organizationSlug+"/"+projectSlug+"/alert-rules/").BodyJSON(params).Receive(apmRule, apiError)
	return apmRule, resp, relevantError(err, *apiError)
}

// Update a APM rule.
func (s *APMRuleService) Update(organizationSlug string, projectSlug string, apmRuleID string, params *APMRule) (*APMRule, *http.Response, error) {
	apmRule := new(APMRule)
	apiError := new(APIError)
	resp, err := s.sling.New().Put("projects/"+organizationSlug+"/"+projectSlug+"/alert-rules/"+apmRuleID+"/").BodyJSON(params).Receive(apmRule, apiError)
	return apmRule, resp, relevantError(err, *apiError)
}

// Delete a APM rule.
func (s *APMRuleService) Delete(organizationSlug string, projectSlug string, apmRuleID string) (*http.Response, error) {
	apiError := new(APIError)
	resp, err := s.sling.New().Delete("projects/"+organizationSlug+"/"+projectSlug+"/alert-rules/"+apmRuleID+"/").Receive(nil, apiError)
	return resp, relevantError(err, *apiError)
}
