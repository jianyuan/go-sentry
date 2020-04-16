package sentry

import (
	"net/http"

	"github.com/dghubble/sling"
)

// OrganizationUser represents a Sentry organization user.
type OrganizationUser struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// OrganizationUserService provides methods for accessing Sentry organization users
// https://docs.sentry.io/api/organizations/get-organization-users/
type OrganizationUserService struct {
	sling *sling.Sling
}

func newOrganizationUserService(sling *sling.Sling) *OrganizationUserService {
	return &OrganizationUserService{
		sling: sling,
	}
}

// List an organization's users
// https://docs.sentry.io/api/organizations/get-organization-users/
func (s *OrganizationUserService) List(organizationSlug string) ([]OrganizationUser, *http.Response, error) {
	users := new([]OrganizationUser)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("organizations/"+organizationSlug+"/users/").Receive(users, apiError)
	return *users, resp, relevantError(err, *apiError)
}
