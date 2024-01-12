package sentry

import (
	"context"
	"fmt"
	"time"
)

// OrganizationMember represents a User's membership to the organization.
// https://github.com/getsentry/sentry/blob/22.5.0/src/sentry/api/serializers/models/organization_member/response.py#L57-L69
type OrganizationMember struct {
	ID           string                     `json:"id"`
	Email        string                     `json:"email"`
	Name         string                     `json:"name"`
	User         User                       `json:"user"`
	OrgRole      string                     `json:"orgRole"`
	OrgRoleList  []OrganizationRoleListItem `json:"orgRoleList"`
	Pending      bool                       `json:"pending"`
	Expired      bool                       `json:"expired"`
	Flags        map[string]bool            `json:"flags"`
	DateCreated  time.Time                  `json:"dateCreated"`
	InviteStatus string                     `json:"inviteStatus"`
	InviterName  *string                    `json:"inviterName"`
	TeamRoleList []TeamRoleListItem         `json:"teamRoleList"`
	TeamRoles    []TeamRole                 `json:"teamRoles"`
	Teams        []string                   `json:"teams"`
}

const (
	OrganizationRoleBilling string = "billing"
	OrganizationRoleMember  string = "member"
	OrganizationRoleManager string = "manager"
	OrganizationRoleOwner   string = "owner"

	TeamRoleContributor string = "contributor"
	TeamRoleAdmin       string = "admin"
)

// OrganizationMembersService provides methods for accessing Sentry membership API endpoints.
type OrganizationMembersService service

// List organization members.
func (s *OrganizationMembersService) List(ctx context.Context, organizationSlug string, params *ListCursorParams) ([]*OrganizationMember, *Response, error) {
	u := fmt.Sprintf("0/organizations/%v/members/", organizationSlug)
	u, err := addQuery(u, params)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	members := []*OrganizationMember{}
	resp, err := s.client.Do(ctx, req, &members)
	if err != nil {
		return nil, resp, err
	}
	return members, resp, nil
}

func (s *OrganizationMembersService) Get(ctx context.Context, organizationSlug string, memberID string) (*OrganizationMember, *Response, error) {
	u := fmt.Sprintf("0/organizations/%v/members/%v/", organizationSlug, memberID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	member := new(OrganizationMember)
	resp, err := s.client.Do(ctx, req, member)
	if err != nil {
		return nil, resp, err
	}
	return member, resp, nil
}

type CreateOrganizationMemberParams struct {
	Email string   `json:"email"`
	Role  string   `json:"role"`
	Teams []string `json:"teams,omitempty"`
}

func (s *OrganizationMembersService) Create(ctx context.Context, organizationSlug string, params *CreateOrganizationMemberParams) (*OrganizationMember, *Response, error) {
	u := fmt.Sprintf("0/organizations/%v/members/", organizationSlug)
	req, err := s.client.NewRequest("POST", u, params)
	if err != nil {
		return nil, nil, err
	}

	member := new(OrganizationMember)
	resp, err := s.client.Do(ctx, req, member)
	if err != nil {
		return nil, resp, err
	}
	return member, resp, nil
}

type TeamRole struct {
	TeamSlug string  `json:"teamSlug"`
	Role     *string `json:"role"`
}

type UpdateOrganizationMemberParams struct {
	OrganizationRole string     `json:"role"`
	TeamRoles        []TeamRole `json:"teamRoles"`
}

func (s *OrganizationMembersService) Update(ctx context.Context, organizationSlug string, memberID string, params *UpdateOrganizationMemberParams) (*OrganizationMember, *Response, error) {
	u := fmt.Sprintf("0/organizations/%v/members/%v/", organizationSlug, memberID)
	req, err := s.client.NewRequest("PUT", u, params)
	if err != nil {
		return nil, nil, err
	}

	member := new(OrganizationMember)
	resp, err := s.client.Do(ctx, req, member)
	if err != nil {
		return nil, resp, err
	}
	return member, resp, nil
}

func (s *OrganizationMembersService) Delete(ctx context.Context, organizationSlug string, memberID string) (*Response, error) {
	u := fmt.Sprintf("0/organizations/%v/members/%v/", organizationSlug, memberID)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
