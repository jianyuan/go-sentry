package sentry

import (
	"net/http"

	"github.com/dghubble/sling"
)

// TeamMember represents a Sentry team member.
type TeamMember struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Teams []string `json:"teams"`
}

// TeamMemberService provides methods for accessing Sentry team members
// /api/0/organizations/{organization_slug}/members/{user_id}/
type TeamMemberService struct {
	sling *sling.Sling
}

func newTeamMemberService(sling *sling.Sling) *TeamMemberService {
	return &TeamMemberService{
		sling: sling,
	}
}

// Get a specific team member
// GET /api/0/organizations/{organization_slug}/members/{user_id}/
func (s *TeamMemberService) Get(organizationSlug string, userID string) (TeamMember, *http.Response, error) {
	member := new(TeamMember)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("organizations/"+organizationSlug+"/members/"+userID+"/").Receive(member, apiError)
	return *member, resp, relevantError(err, *apiError)
}

// UpdateTeamMemberParams are the parameters for TeamMemberService.Update.
type UpdateTeamMemberParams struct {
	Teams []string `json:"teams"`
}

// Update a specific team member
// PUT /api/0/organizations/{organization_slug}/members/{user_id}/
func (s *TeamMemberService) Update(organizationSlug string, userID string, params *UpdateTeamMemberParams) (TeamMember, *http.Response, error) {
	member := new(TeamMember)
	apiError := new(APIError)
	resp, err := s.sling.New().Put("organizations/"+organizationSlug+"/members/"+userID+"/").BodyJSON(params).Receive(member, apiError)
	return *member, resp, relevantError(err, *apiError)
}
