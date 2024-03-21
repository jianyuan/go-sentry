package sentry

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type TeamMember struct {
	ID          *string         `json:"id"`
	Slug        *string         `json:"slug"`
	Name        *string         `json:"name"`
	DateCreated *time.Time      `json:"dateCreated"`
	IsMember    *bool           `json:"isMember"`
	TeamRole    *string         `json:"teamRole"`
	Flags       map[string]bool `json:"flags"`
	Access      []string        `json:"access"`
	HasAccess   *bool           `json:"hasAccess"`
	IsPending   *bool           `json:"isPending"`
	MemberCount *int            `json:"memberCount"`
	Avatar      *Avatar         `json:"avatar"`
}

// TeamMember provides methods for accessing Sentry team member API endpoints.
type TeamMembersService service

func (s *TeamMembersService) Create(ctx context.Context, organizationSlug string, memberID string, teamSlug string) (*TeamMember, *Response, error) {
	u := fmt.Sprintf("0/organizations/%v/members/%v/teams/%v/", organizationSlug, memberID, teamSlug)
	req, err := s.client.NewRequest(http.MethodPost, u, nil)
	if err != nil {
		return nil, nil, err
	}

	member := new(TeamMember)
	resp, err := s.client.Do(ctx, req, member)
	if err != nil {
		return nil, resp, err
	}
	return member, resp, nil
}

type UpdateTeamMemberParams struct {
	TeamRole *string `json:"teamRole,omitempty"`
}

type UpdateTeamMemberResponse struct {
	IsActive *bool   `json:"isActive,omitempty"`
	TeamRole *string `json:"teamRole,omitempty"`
}

func (s *TeamMembersService) Update(ctx context.Context, organizationSlug string, memberID string, teamSlug string, params *UpdateTeamMemberParams) (*UpdateTeamMemberResponse, *Response, error) {
	u := fmt.Sprintf("0/organizations/%v/members/%v/teams/%v/", organizationSlug, memberID, teamSlug)
	req, err := s.client.NewRequest(http.MethodPut, u, params)
	if err != nil {
		return nil, nil, err
	}

	member := new(UpdateTeamMemberResponse)
	resp, err := s.client.Do(ctx, req, member)
	if err != nil {
		return nil, resp, err
	}
	return member, resp, nil
}

func (s *TeamMembersService) Delete(ctx context.Context, organizationSlug string, memberID string, teamSlug string) (*TeamMember, *Response, error) {
	u := fmt.Sprintf("0/organizations/%v/members/%v/teams/%v/", organizationSlug, memberID, teamSlug)
	req, err := s.client.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return nil, nil, err
	}

	member := new(TeamMember)
	resp, err := s.client.Do(ctx, req, member)
	if err != nil {
		return nil, resp, err
	}
	return member, resp, nil
}
