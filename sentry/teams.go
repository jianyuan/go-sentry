package sentry

import (
	"net/http"
	"time"

	"github.com/dghubble/sling"
)

// Team represents a Sentry team that is bound to an organization.
type Team struct {
	ID          string    `json:"id"`
	Slug        string    `json:"slug"`
	Name        string    `json:"name"`
	DateCreated time.Time `json:"dateCreated"`
	HasAccess   bool      `json:"hasAccess"`
	IsPending   bool      `json:"isPending"`
	IsMember    bool      `json:"isMember"`
}

// TeamService provides methods for accessing Sentry team API endpoints.
// https://docs.sentry.io/api/teams/
type TeamService struct {
	sling *sling.Sling
}

func newTeamService(sling *sling.Sling) *TeamService {
	return &TeamService{
		sling: sling.Path("organizations/"),
	}
}

// CreateTeamParams are the parameters for TeamService.Create.
type CreateTeamParams struct {
	Name string `json:"name,omitempty"`
	Slug string `json:"slug,omitempty"`
}

// Create a new Sentry team bound to an organization.
// https://docs.sentry.io/api/teams/post-organization-teams/
func (s *TeamService) Create(organizationSlug string, params *CreateTeamParams) (*Team, *http.Response, error) {
	team := new(Team)
	apiError := new(APIError)
	resp, err := s.sling.New().Post(organizationSlug+"/teams/").BodyJSON(params).Receive(team, apiError)
	return team, resp, relevantError(err, *apiError)
}
