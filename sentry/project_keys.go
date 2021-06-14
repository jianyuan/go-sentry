package sentry

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dghubble/sling"
	"github.com/tomnomnom/linkheader"
)

// ProjectKeyRateLimit represents a project key's rate limit.
type ProjectKeyRateLimit struct {
	Window int `json:"window"`
	Count  int `json:"count"`
}

// ProjectKeyDSN represents a project key's DSN.
type ProjectKeyDSN struct {
	Secret   string `json:"secret"`
	Public   string `json:"public"`
	CSP      string `json:"csp"`
	Security string `json:"security"`
	Minidump string `json:"minidump"`
	CDN      string `json:"cdn"`
}

// ProjectKey represents a client key bound to a project.
// https://github.com/getsentry/sentry/blob/9.0.0/src/sentry/api/serializers/models/project_key.py
type ProjectKey struct {
	ID          string               `json:"id"`
	Name        string               `json:"name"`
	Label       string               `json:"label"`
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
	cursor := ""
	return s.listPerPage(organizationSlug, projectSlug, cursor)
}

// https://docs.sentry.io/api/projects/get-project-keys/
func (s *ProjectKeyService) listPerPage(organizationSlug string, projectSlug string, cursor string) ([]ProjectKey, *http.Response, error) {
	projectKeys := new([]ProjectKey)
	apiError := new(APIError)

	URL := "projects/"+organizationSlug+"/"+projectSlug+"/keys/" + cursor
	resp, err := s.sling.New().Get(URL).Receive(projectKeys, apiError)
	if resp != nil && resp.StatusCode == 200 {
		linkHeaders := linkheader.Parse(resp.Header.Get("Link"))
		// If the next Link has results query it as well
		nextLink := linkHeaders[len(linkHeaders) - 1]

		if nextLink.Param("results") == "true" {
			c := fmt.Sprintf("?&cursor=%s", nextLink.Param("cursor"))
			pagedProjectKeys, pagedResp, err2 := s.listPerPage(organizationSlug, projectSlug, c)
			if err2 != nil {
				return nil, pagedResp, relevantError(err2, *apiError)
			}
			*projectKeys = append(*projectKeys, pagedProjectKeys...)
		}
	}

	return *projectKeys, resp, relevantError(err, *apiError)
}

// CreateProjectKeyParams are the parameters for ProjectKeyService.Create.
type CreateProjectKeyParams struct {
	Name      string               `json:"name,omitempty"`
	RateLimit *ProjectKeyRateLimit `json:"rateLimit,omitempty"`
}

// Create a new client key bound to a project.
// https://docs.sentry.io/api/projects/post-project-keys/
func (s *ProjectKeyService) Create(organizationSlug string, projectSlug string, params *CreateProjectKeyParams) (*ProjectKey, *http.Response, error) {
	projectKey := new(ProjectKey)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("projects/"+organizationSlug+"/"+projectSlug+"/keys/").BodyJSON(params).Receive(projectKey, apiError)

	if err != nil {
		return projectKey, resp, relevantError(err, *apiError)
	}

	// Hack as currently the API does not support setting rate limits on Create
	if params.RateLimit != nil {
		updateParams := &UpdateProjectKeyParams{
			Name:      params.Name,
			RateLimit: params.RateLimit,
		}
		projectKey, resp, err = s.Update(organizationSlug, projectSlug, projectKey.ID, updateParams)
	}

	return projectKey, resp, relevantError(err, *apiError)
}

// UpdateProjectKeyParams are the parameters for ProjectKeyService.Update.
type UpdateProjectKeyParams struct {
	Name      string               `json:"name,omitempty"`
	RateLimit *ProjectKeyRateLimit `json:"rateLimit,omitempty"`
}

// Update a client key.
// https://docs.sentry.io/api/projects/put-project-key-details/
func (s *ProjectKeyService) Update(organizationSlug string, projectSlug string, keyID string, params *UpdateProjectKeyParams) (*ProjectKey, *http.Response, error) {
	projectKey := new(ProjectKey)
	apiError := new(APIError)
	resp, err := s.sling.New().Put("projects/"+organizationSlug+"/"+projectSlug+"/keys/"+keyID+"/").BodyJSON(params).Receive(projectKey, apiError)
	return projectKey, resp, relevantError(err, *apiError)
}

// Delete a project.
// https://docs.sentry.io/api/projects/delete-project-details/
func (s *ProjectKeyService) Delete(organizationSlug string, projectSlug string, keyID string) (*http.Response, error) {
	apiError := new(APIError)
	resp, err := s.sling.New().Delete("projects/"+organizationSlug+"/"+projectSlug+"/keys/"+keyID+"/").Receive(nil, apiError)
	return resp, relevantError(err, *apiError)
}
