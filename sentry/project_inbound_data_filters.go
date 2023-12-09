package sentry

import (
	"context"
	"net/http"
)

type ProjectInboundDataFilter struct {
	ID     string            `json:"id"`
	Active BoolOrStringSlice `json:"active"`
}

type ProjectInboundDataFiltersService service

func (s *ProjectInboundDataFiltersService) List(ctx context.Context, organizationSlug string, projectSlug string) ([]*ProjectInboundDataFilter, *Response, error) {
	u := "0/projects/" + organizationSlug + "/" + projectSlug + "/filters/"
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	filters := []*ProjectInboundDataFilter{}
	resp, err := s.client.Do(ctx, req, &filters)
	if err != nil {
		return nil, resp, err
	}
	return filters, resp, nil
}

type UpdateProjectInboundDataFilterParams struct {
	Active     *bool    `json:"active,omitempty"`
	Subfilters []string `json:"subfilters,omitempty"`
}

func (s *ProjectInboundDataFiltersService) Update(ctx context.Context, organizationSlug string, projectSlug string, filterID string, params *UpdateProjectInboundDataFilterParams) (*Response, error) {
	u := "0/projects/" + organizationSlug + "/" + projectSlug + "/filters/" + filterID + "/"
	req, err := s.client.NewRequest(http.MethodPut, u, params)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}
