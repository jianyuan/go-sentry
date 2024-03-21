package sentry

import (
	"context"
	"net/http"
)

type ProjectSymbolSourceLayout struct {
	Type   *string `json:"type"`
	Casing *string `json:"casing"`
}

type ProjectSymbolSourceHiddenSecret struct {
	HiddenSecret *bool `json:"hidden-secret"`
}

type ProjectSymbolSource struct {
	ID     *string                    `json:"id"`
	Type   *string                    `json:"type"`
	Name   *string                    `json:"name"`
	Layout *ProjectSymbolSourceLayout `json:"layout"`

	AppConnectIssuer     *string                          `json:"appconnectIssuer,omitempty"`
	AppConnectPrivateKey *ProjectSymbolSourceHiddenSecret `json:"appconnectPrivateKey,omitempty"`
	AppId                *string                          `json:"appId,omitempty"`
	Url                  *string                          `json:"url,omitempty"`
	Username             *string                          `json:"username,omitempty"`
	Password             *ProjectSymbolSourceHiddenSecret `json:"password,omitempty"`
	Bucket               *string                          `json:"bucket,omitempty"`
	Region               *string                          `json:"region,omitempty"`
	AccessKey            *string                          `json:"access_key,omitempty"`
	SecretKey            *ProjectSymbolSourceHiddenSecret `json:"secret_key,omitempty"`
	Prefix               *string                          `json:"prefix,omitempty"`
	ClientEmail          *string                          `json:"client_email,omitempty"`
	PrivateKey           *ProjectSymbolSourceHiddenSecret `json:"private_key,omitempty"`
}

type ProjectSymbolSourcesService service

type ProjectSymbolSourceQueryParams struct {
	ID *string `url:"id,omitempty"`
}

func (s *ProjectSymbolSourcesService) List(ctx context.Context, organizationSlug string, projectSlug string, params *ProjectSymbolSourceQueryParams) ([]*ProjectSymbolSource, *Response, error) {
	u := "0/projects/" + organizationSlug + "/" + projectSlug + "/symbol-sources/"
	u, err := addQuery(u, params)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	filters := []*ProjectSymbolSource{}
	resp, err := s.client.Do(ctx, req, &filters)
	if err != nil {
		return nil, resp, err
	}
	return filters, resp, nil
}

type CreateProjectSymbolSourceParams struct {
	Type   *string                    `json:"type"`
	Name   *string                    `json:"name"`
	Layout *ProjectSymbolSourceLayout `json:"layout"`

	AppConnectIssuer     *string `json:"appconnectIssuer,omitempty"`
	AppConnectPrivateKey *string `json:"appconnectPrivateKey,omitempty"`
	AppId                *string `json:"appId,omitempty"`
	Url                  *string `json:"url,omitempty"`
	Username             *string `json:"username,omitempty"`
	Password             *string `json:"password,omitempty"`
	Bucket               *string `json:"bucket,omitempty"`
	Region               *string `json:"region,omitempty"`
	AccessKey            *string `json:"access_key,omitempty"`
	SecretKey            *string `json:"secret_key,omitempty"`
	Prefix               *string `json:"prefix,omitempty"`
	ClientEmail          *string `json:"client_email,omitempty"`
	PrivateKey           *string `json:"private_key,omitempty"`
}

func (s *ProjectSymbolSourcesService) Create(ctx context.Context, organizationSlug string, projectSlug string, params *CreateProjectSymbolSourceParams) (*ProjectSymbolSource, *Response, error) {
	u := "0/projects/" + organizationSlug + "/" + projectSlug + "/symbol-sources/"
	req, err := s.client.NewRequest(http.MethodPost, u, params)
	if err != nil {
		return nil, nil, err
	}

	filter := &ProjectSymbolSource{}
	resp, err := s.client.Do(ctx, req, filter)
	if err != nil {
		return nil, resp, err
	}
	return filter, resp, nil
}

type UpdateProjectSymbolSourceParams struct {
	ID     *string                    `json:"id"`
	Type   *string                    `json:"type"`
	Name   *string                    `json:"name"`
	Layout *ProjectSymbolSourceLayout `json:"layout"`

	AppConnectIssuer     *string `json:"appconnectIssuer,omitempty"`
	AppConnectPrivateKey *string `json:"appconnectPrivateKey,omitempty"`
	AppId                *string `json:"appId,omitempty"`
	Url                  *string `json:"url,omitempty"`
	Username             *string `json:"username,omitempty"`
	Password             *string `json:"password,omitempty"`
	Bucket               *string `json:"bucket,omitempty"`
	Region               *string `json:"region,omitempty"`
	AccessKey            *string `json:"access_key,omitempty"`
	SecretKey            *string `json:"secret_key,omitempty"`
	Prefix               *string `json:"prefix,omitempty"`
	ClientEmail          *string `json:"client_email,omitempty"`
	PrivateKey           *string `json:"private_key,omitempty"`
}

func (s *ProjectSymbolSourcesService) Update(ctx context.Context, organizationSlug string, projectSlug string, symbolSourceId string, params *UpdateProjectSymbolSourceParams) (*ProjectSymbolSource, *Response, error) {
	u := "0/projects/" + organizationSlug + "/" + projectSlug + "/symbol-sources/"
	u, err := addQuery(u, &ProjectSymbolSourceQueryParams{
		ID: String(symbolSourceId),
	})
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodPut, u, params)
	if err != nil {
		return nil, nil, err
	}

	filter := &ProjectSymbolSource{}
	resp, err := s.client.Do(ctx, req, filter)
	if err != nil {
		return nil, resp, err
	}
	return filter, resp, nil
}

func (s *ProjectSymbolSourcesService) Delete(ctx context.Context, organizationSlug string, projectSlug string, symbolSourceId string) (*Response, error) {
	u := "0/projects/" + organizationSlug + "/" + projectSlug + "/symbol-sources/"
	u, err := addQuery(u, &ProjectSymbolSourceQueryParams{
		ID: String(symbolSourceId),
	})
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
