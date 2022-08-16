package sentry

import (
	"context"
	"fmt"
)

// OrganizationCodeMapping represents a code mapping added for the organization.
// https://github.com/getsentry/sentry/blob/22.7.0/src/sentry/api/serializers/models/repository_project_path_config.py
type OrganizationCodeMapping struct {
	ID            string                           `json:"id"`
	ProjectId     string                           `json:"projectId"`
	ProjectSlug   string                           `json:"projectSlug"`
	RepoId        string                           `json:"repoId"`
	RepoName      string                           `json:"repoName"`
	IntegrationId string                           `json:"integrationId"`
	Provider      *OrganizationIntegrationProvider `json:"provider"`
	StackRoot     string                           `json:"stackRoot"`
	SourceRoot    string                           `json:"sourceRoot"`
	DefaultBranch string                           `json:"defaultBranch"`
}

// OrganizationCodeMappingsService provides methods for accessing Sentry organization code mappings API endpoints.
// Paths: https://github.com/getsentry/sentry/blob/22.7.0/src/sentry/api/urls.py#L929-L938
// Endpoints: https://github.com/getsentry/sentry/blob/22.7.0/src/sentry/api/endpoints/organization_code_mappings.py
// Endpoints: https://github.com/getsentry/sentry/blob/22.7.0/src/sentry/api/endpoints/organization_code_mapping_details.py
type OrganizationCodeMappingsService service

type ListOrganizationCodeMappingsParams struct {
	ListCursorParams
	IntegrationId string `url:"integrationId,omitempty"`
}

// List organization integrations.
func (s *OrganizationCodeMappingsService) List(ctx context.Context, organizationSlug string, params *ListOrganizationCodeMappingsParams) ([]*OrganizationCodeMapping, *Response, error) {
	u := fmt.Sprintf("0/organizations/%v/code-mappings/", organizationSlug)
	u, err := addQuery(u, params)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	integrations := []*OrganizationCodeMapping{}
	resp, err := s.client.Do(ctx, req, &integrations)
	if err != nil {
		return nil, resp, err
	}
	return integrations, resp, nil
}
