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
