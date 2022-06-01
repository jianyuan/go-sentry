package sentry

import (
	"net/http"
	"time"

	"github.com/dghubble/sling"
)

// OrganizationMember represents a User's membership to the organization.
// https://github.com/getsentry/sentry/blob/275e6efa0f364ce05d9bfd09386b895b8a5e0671/src/sentry/api/serializers/models/organization_member.py#L12
type OrganizationMember struct {
	ID           string          `json:"id"`
	Email        string          `json:"email"`
	Name         string          `json:"name"`
	User         User            `json:"user"`
	Role         string          `json:"role"`
	RoleName     string          `json:"roleName"`
	Pending      bool            `json:"pending"`
	Expired      bool            `json:"expired"`
	Flags        map[string]bool `json:"flags"`
	DateCreated  time.Time       `json:"dateCreated"`
	InviteStatus string          `json:"inviteStatus"`
	InviterName  *string         `json:"inviterName"`
	Teams        []string        `json:"teams"`
}

// OrganizationMemberService provides methods for accessing Sentry membership API endpoints.
type OrganizationMemberService struct {
	sling *sling.Sling
}

func newOrganizationMemberService(sling *sling.Sling) *OrganizationMemberService {
	return &OrganizationMemberService{
		sling: sling,
	}
}

// ListOrganizationMemberParams are the parameters for OrganizationMemberService.List.
type ListOrganizationMemberParams struct {
	Cursor string `url:"cursor,omitempty"`
}

// List organization members.
func (s *OrganizationMemberService) List(organizationSlug string, params *ListOrganizationMemberParams) ([]OrganizationMember, *http.Response, error) {
	members := new([]OrganizationMember)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("organizations/"+organizationSlug+"/members/").QueryStruct(params).Receive(members, apiError)
	return *members, resp, relevantError(err, *apiError)
}

const (
	RoleMember  string = "member"
	RoleBilling string = "billing"
	RoleAdmin   string = "admin"
	RoleOwner   string = "owner"
	RoleManager string = "manager"
)

type CreateOrganizationMemberParams struct {
	Email string   `json:"email"`
	Role  string   `json:"role"`
	Teams []string `json:"teams,omitempty"`
}

func (s *OrganizationMemberService) Create(organizationSlug string, params *CreateOrganizationMemberParams) (*OrganizationMember, *http.Response, error) {
	apiError := new(APIError)
	organizationMember := new(OrganizationMember)
	resp, err := s.sling.New().Post("organizations/"+organizationSlug+"/members/").BodyJSON(params).Receive(organizationMember, apiError)
	return organizationMember, resp, relevantError(err, *apiError)
}

func (s *OrganizationMemberService) Delete(organizationSlug string, memberId string) (*http.Response, error) {
	apiError := new(APIError)
	resp, err := s.sling.New().Delete("organizations/"+organizationSlug+"/members/"+memberId).Receive(nil, apiError)
	return resp, relevantError(err, *apiError)
}

func (s *OrganizationMemberService) Get(organizationSlug string, memberId string) (*OrganizationMember, *http.Response, error) {
	apiError := new(APIError)
	organizationMember := new(OrganizationMember)
	resp, err := s.sling.New().Get("organizations/"+organizationSlug+"/members/"+memberId).Receive(organizationMember, apiError)
	return organizationMember, resp, relevantError(err, *apiError)
}

type UpdateOrganizationMemberParams struct {
	Role  string   `json:"role"`
	Teams []string `json:"teams,omitempty"`
}

func (s *OrganizationMemberService) Update(organizationSlug string, memberId string, params *UpdateOrganizationMemberParams) (*OrganizationMember, *http.Response, error) {
	apiError := new(APIError)
	organizationMember := new(OrganizationMember)
	resp, err := s.sling.New().Put("organizations/"+organizationSlug+"/members/"+memberId).BodyJSON(params).Receive(organizationMember, apiError)
	return organizationMember, resp, relevantError(err, *apiError)
}
