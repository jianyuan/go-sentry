package sentry

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
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
	NEL      string `json:"nel"`
	Unreal   string `json:"unreal"`
	CDN      string `json:"cdn"`
	Crons    string `json:"crons"`
}

type ProjectKeyDynamicSDKLoaderOptions struct {
	HasReplay      bool `json:"hasReplay"`
	HasPerformance bool `json:"hasPerformance"`
	HasDebugFiles  bool `json:"hasDebug"`
}

// ProjectKey represents a client key bound to a project.
// https://github.com/getsentry/sentry/blob/9.0.0/src/sentry/api/serializers/models/project_key.py
type ProjectKey struct {
	ID                      string                            `json:"id"`
	Name                    string                            `json:"name"`
	Label                   string                            `json:"label"`
	Public                  string                            `json:"public"`
	Secret                  string                            `json:"secret"`
	ProjectID               json.Number                       `json:"projectId"`
	IsActive                bool                              `json:"isActive"`
	RateLimit               *ProjectKeyRateLimit              `json:"rateLimit"`
	DSN                     ProjectKeyDSN                     `json:"dsn"`
	BrowserSDKVersion       string                            `json:"browserSdkVersion"`
	DateCreated             time.Time                         `json:"dateCreated"`
	DynamicSDKLoaderOptions ProjectKeyDynamicSDKLoaderOptions `json:"dynamicSdkLoaderOptions"`
}

// ProjectKeysService provides methods for accessing Sentry project
// client key API endpoints.
// https://docs.sentry.io/api/projects/
type ProjectKeysService service

type ListProjectKeysParams struct {
	ListCursorParams

	Status *string `url:"status,omitempty"`
}

// List client keys bound to a project.
// https://docs.sentry.io/api/projects/get-project-keys/
func (s *ProjectKeysService) List(ctx context.Context, organizationSlug string, projectSlug string, params *ListProjectKeysParams) ([]*ProjectKey, *Response, error) {
	u := fmt.Sprintf("0/projects/%v/%v/keys/", organizationSlug, projectSlug)
	u, err := addQuery(u, params)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	projectKeys := []*ProjectKey{}
	resp, err := s.client.Do(ctx, req, &projectKeys)
	if err != nil {
		return nil, resp, err
	}
	return projectKeys, resp, nil
}

// Get details of a client key.
// https://docs.sentry.io/api/projects/retrieve-a-client-key/
func (s *ProjectKeysService) Get(ctx context.Context, organizationSlug string, projectSlug string, id string) (*ProjectKey, *Response, error) {
	u := fmt.Sprintf("0/projects/%v/%v/keys/%v/", organizationSlug, projectSlug, id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	projectKey := new(ProjectKey)
	resp, err := s.client.Do(ctx, req, projectKey)
	if err != nil {
		return nil, resp, err
	}
	return projectKey, resp, nil
}

// CreateProjectKeyParams are the parameters for ProjectKeyService.Create.
type CreateProjectKeyParams struct {
	Name      string               `json:"name,omitempty"`
	RateLimit *ProjectKeyRateLimit `json:"rateLimit,omitempty"`
}

// Create a new client key bound to a project.
// https://docs.sentry.io/api/projects/post-project-keys/
func (s *ProjectKeysService) Create(ctx context.Context, organizationSlug string, projectSlug string, params *CreateProjectKeyParams) (*ProjectKey, *Response, error) {
	u := fmt.Sprintf("0/projects/%v/%v/keys/", organizationSlug, projectSlug)
	req, err := s.client.NewRequest("POST", u, params)
	if err != nil {
		return nil, nil, err
	}

	projectKey := new(ProjectKey)
	resp, err := s.client.Do(ctx, req, projectKey)
	if err != nil {
		return nil, resp, err
	}
	return projectKey, resp, nil
}

// UpdateProjectKeyParams are the parameters for ProjectKeyService.Update.
type UpdateProjectKeyParams struct {
	Name      string               `json:"name,omitempty"`
	RateLimit *ProjectKeyRateLimit `json:"rateLimit,omitempty"`
}

// Update a client key.
// https://docs.sentry.io/api/projects/put-project-key-details/
func (s *ProjectKeysService) Update(ctx context.Context, organizationSlug string, projectSlug string, keyID string, params *UpdateProjectKeyParams) (*ProjectKey, *Response, error) {
	u := fmt.Sprintf("0/projects/%v/%v/keys/%v/", organizationSlug, projectSlug, keyID)
	req, err := s.client.NewRequest("PUT", u, params)
	if err != nil {
		return nil, nil, err
	}

	projectKey := new(ProjectKey)
	resp, err := s.client.Do(ctx, req, projectKey)
	if err != nil {
		return nil, resp, err
	}
	return projectKey, resp, nil
}

// Delete a project.
// https://docs.sentry.io/api/projects/delete-project-details/
func (s *ProjectKeysService) Delete(ctx context.Context, organizationSlug string, projectSlug string, keyID string) (*Response, error) {
	u := fmt.Sprintf("0/projects/%v/%v/keys/%v/", organizationSlug, projectSlug, keyID)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
