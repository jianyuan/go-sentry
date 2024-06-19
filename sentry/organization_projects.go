package sentry

import (
	"context"
	"fmt"
)

type OrganizationProjectsService service

type ListOrganizationProjectsParams struct {
	ListCursorParams

	Options string `url:"options,omitempty"`
	Query   string `url:"query,omitempty"`
}

// List an Organization's Projects
// https://docs.sentry.io/api/organizations/list-an-organizations-projects/
func (s *OrganizationProjectsService) List(ctx context.Context, organizationSlug string, params *ListOrganizationProjectsParams) ([]*Project, *Response, error) {
	u := fmt.Sprintf("0/organizations/%v/projects/", organizationSlug)
	u, err := addQuery(u, params)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	projects := []*Project{}
	resp, err := s.client.Do(ctx, req, &projects)
	if err != nil {
		return nil, resp, err
	}
	return projects, resp, nil
}
