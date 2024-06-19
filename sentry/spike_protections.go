package sentry

import (
	"context"
	"fmt"
	"net/http"
)

type SpikeProtectionsService service

type SpikeProtectionParams struct {
	ProjectSlug string   `json:"-" url:"projectSlug,omitempty"`
	Projects    []string `json:"projects" url:"-"`
}

func (s *SpikeProtectionsService) Enable(ctx context.Context, organizationSlug string, params *SpikeProtectionParams) (*Response, error) {
	u := fmt.Sprintf("0/organizations/%v/spike-protections/", organizationSlug)
	u, err := addQuery(u, params)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest(http.MethodPost, u, params)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

func (s *SpikeProtectionsService) Disable(ctx context.Context, organizationSlug string, params *SpikeProtectionParams) (*Response, error) {
	u := fmt.Sprintf("0/organizations/%v/spike-protections/", organizationSlug)
	u, err := addQuery(u, params)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest(http.MethodDelete, u, params)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
