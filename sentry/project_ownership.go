package sentry

import (
	"net/http"
	"time"

	"github.com/dghubble/sling"
)

// https://github.com/getsentry/sentry/blob/master/src/sentry/api/serializers/models/projectownership.py
type ProjectOwnership struct {
	Raw                string    `json:"raw"`
	FallThrough        bool      `json:"fallthrough"`
	DateCreated        time.Time `json:"dateCreated"`
	LastUpdated        time.Time `json:"lastUpdated"`
	IsActive           bool      `json:"isActive"`
	AutoAssignment     bool      `json:"autoAssignment"`
	CodeownersAutoSync *bool     `json:"codeownersAutoSync,omitempty"`
}

// ProjectOwnershipService provides methods for accessing Sentry project
// client key API endpoints.
type ProjectOwnershipService struct {
	sling *sling.Sling
}

func newProjectOwnershipService(sling *sling.Sling) *ProjectOwnershipService {
	return &ProjectOwnershipService{
		sling: sling,
	}
}

// Get details on a project's ownership configuration.
func (s *ProjectOwnershipService) Get(organizationSlug string, projectSlug string) (*ProjectOwnership, *http.Response, error) {
	project := new(ProjectOwnership)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("projects/"+organizationSlug+"/"+projectSlug+"/ownership/").Receive(project, apiError)
	return project, resp, relevantError(err, *apiError)
}

// CreateProjectParams are the parameters for ProjectOwnershipService.Update.
type UpdateProjectOwnershipParams struct {
	Raw                string `json:"raw,omitempty"`
	FallThrough        *bool  `json:"fallthrough,omitempty"`
	AutoAssignment     *bool  `json:"autoAssignment,omitempty"`
	CodeownersAutoSync *bool  `json:"codeownersAutoSync,omitempty"`
}

// Update a Project's Ownership configuration
func (s *ProjectOwnershipService) Update(organizationSlug string, projectSlug string, params *UpdateProjectOwnershipParams) (*ProjectOwnership, *http.Response, error) {
	project := new(ProjectOwnership)
	apiError := new(APIError)
	resp, err := s.sling.New().Put("projects/"+organizationSlug+"/"+projectSlug+"/ownership/").BodyJSON(params).Receive(project, apiError)
	return project, resp, relevantError(err, *apiError)
}
