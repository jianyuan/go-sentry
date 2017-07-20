package sentry

import (
	"net/http"
	"time"

	"github.com/dghubble/sling"
)

// ProjectKeyRateLimit represents a project key's rate limit.
type ProjectKeyRateLimit struct {
	Window int `json:"window"`
	Count  int `json:"count"`
}

// ProjectKeyDSN represents a project key's DSN.
type ProjectKeyDSN struct {
	Secret string `json:"secret"`
	Public string `json:"public"`
	CSP    string `json:"csp"`
}

// ProjectKey represents a client key bound to a project.
// Based on https://github.com/getsentry/sentry/blob/a418072946ebd2933724945e1ea2a833cf4c9b94/src/sentry/api/serializers/models/project_key.py.
type ProjectKey struct {
	ID          string               `json:"id"`
	Name        string               `json:"name"`
	Public      string               `json:"public"`
	Secret      string               `json:"secret"`
	ProjectID   int                  `json:"projectId"`
	IsActive    bool                 `json:"isActive"`
	RateLimit   *ProjectKeyRateLimit `json:"rateLimit"`
	DSN         ProjectKeyDSN        `json:"dsn"`
	DateCreated time.Time            `json:"dateCreated"`
}

// ProjectKeyService provides methods for accessing Sentry project
// client key API endpoints.
// https://docs.sentry.io/api/projects/
type ProjectKeyService struct {
	sling *sling.Sling
}

func newProjectKeyService(sling *sling.Sling) *ProjectKeyService {
	return &ProjectKeyService{
		sling: sling,
	}
}

// List client keys bound to a project.
// https://docs.sentry.io/api/projects/get-project-keys/
func (s *ProjectKeyService) List(organizationSlug string, projectSlug string) ([]ProjectKey, *http.Response, error) {
	projectKeys := new([]ProjectKey)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("projects/"+organizationSlug+"/"+projectSlug+"/keys/").Receive(projectKeys, apiError)
	return *projectKeys, resp, relevantError(err, *apiError)
}

// CreateProjectKeyParams are the parameters for ProjectKeyService.Create.
type CreateProjectKeyParams struct {
	Name string `json:"name,omitempty"`
}

// Create a new client key bound to a project.
// https://docs.sentry.io/api/projects/post-project-keys/
func (s *ProjectKeyService) Create(organizationSlug string, projectSlug string, params *CreateProjectKeyParams) (*ProjectKey, *http.Response, error) {
	projectKey := new(ProjectKey)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("projects/"+organizationSlug+"/"+projectSlug+"/keys/").BodyJSON(params).Receive(projectKey, apiError)
	return projectKey, resp, relevantError(err, *apiError)
}
