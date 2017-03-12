package sentry

import (
	"net/http"
	"time"

	"github.com/dghubble/sling"
)

// Organization represents a Sentry organization.
type Organization struct {
	ID             string    `json:"id"`
	Slug           string    `json:"slug"`
	Name           string    `json:"name"`
	DateCreated    time.Time `json:"dateCreated"`
	IsEarlyAdopter bool      `json:"isEarlyAdopter"`

	// TODO:
	// Teams []Team `json:"teams"`

	Quota struct {
		MaxRate      int `json:"maxRate"`
		ProjectLimit int `json:"projectLimit"`
	} `json:"quota"`
	Access                []string `json:"access"`
	Features              []string `json:"features"`
	PendingAccessRequests int      `json:"pendingAccessRequests"`
}

// OrganizationsService provides methods for accessing Sentry organization API endpoints.
// https://docs.sentry.io/api/organizations/
type OrganizationsService struct {
	sling *sling.Sling
}

func newOrganizationsService(sling *sling.Sling) *OrganizationsService {
	return &OrganizationsService{
		sling: sling.Path("organizations/"),
	}
}

// CreateOrganizationParams are the parameters for OrganizationsService.Create.
type CreateOrganizationParams struct {
	Name string `json:"name,omitempty"`
	Slug string `json:"slug,omitempty"`
}

// Create a new Sentry organization.
// https://docs.sentry.io/api/organizations/post-organization-index/
func (s *OrganizationsService) Create(params *CreateOrganizationParams) (*Organization, *http.Response, error) {
	org := new(Organization)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("").BodyJSON(params).Receive(org, apiError)
	return org, resp, relevantError(err, *apiError)
}

// Get a Sentry organization.
// https://docs.sentry.io/api/organizations/get-organization-details/
func (s *OrganizationsService) Get(slug string) (*Organization, *http.Response, error) {
	org := new(Organization)
	apiError := new(APIError)
	resp, err := s.sling.New().Get(slug+"/").Receive(org, apiError)
	return org, resp, relevantError(err, *apiError)
}

// UpdateOrganizationParams are the parameters for OrganizationsService.Update.
type UpdateOrganizationParams struct {
	Name string `json:"name,omitempty"`
	Slug string `json:"slug,omitempty"`
}

// Update a Sentry organization.
// https://docs.sentry.io/api/organizations/put-organization-details/
func (s *OrganizationsService) Update(slug string, params *UpdateOrganizationParams) (*Organization, *http.Response, error) {
	if params == nil {
		params = &UpdateOrganizationParams{}
	}
	org := new(Organization)
	apiError := new(APIError)
	resp, err := s.sling.New().Put(slug+"/").BodyJSON(params).Receive(org, apiError)
	return org, resp, relevantError(err, *apiError)
}

// Delete a Sentry organization.
func (s *OrganizationsService) Delete(slug string) (*http.Response, error) {
	apiError := new(APIError)
	resp, err := s.sling.New().Delete(slug+"/").Receive(nil, apiError)
	return resp, relevantError(err, *apiError)
}
