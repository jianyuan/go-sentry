package sentry

import (
	"net/http"
	"time"

	"github.com/dghubble/sling"
)

// Project represents a Sentry project.
// https://github.com/getsentry/sentry/blob/9.0.0/src/sentry/api/serializers/models/project.py
type Project struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`

	IsPublic     bool   `json:"isPublic"`
	IsBookmarked bool   `json:"isBookmarked"`
	Color        string `json:"color"`

	DateCreated time.Time `json:"dateCreated"`
	FirstEvent  time.Time `json:"firstEvent"`

	Features []string `json:"features"`
	Status   string   `json:"status"`
	Platform string   `json:"platform"`

	IsInternal bool `json:"isInternal"`
	IsMember   bool `json:"isMember"`
	HasAccess  bool `json:"hasAccess"`

	Avatar Avatar `json:"avatar"`

	// TODO: latestRelease
	Options map[string]interface{} `json:"options"`

	DigestsMinDelay      int      `json:"digestsMinDelay"`
	DigestsMaxDelay      int      `json:"digestsMaxDelay"`
	SubjectPrefix        string   `json:"subjectPrefix"`
	AllowedDomains       []string `json:"allowedDomains"`
	ResolveAge           int      `json:"resolveAge"`
	DataScrubber         bool     `json:"dataScrubber"`
	DataScrubberDefaults bool     `json:"dataScrubberDefaults"`
	SafeFields           []string `json:"safeFields"`
	SensitiveFields      []string `json:"sensitiveFields"`
	SubjectTemplate      string   `json:"subjectTemplate"`
	SecurityToken        string   `json:"securityToken"`
	SecurityTokenHeader  *string  `json:"securityTokenHeader"`
	VerifySSL            bool     `json:"verifySSL"`
	ScrubIPAddresses     bool     `json:"scrubIPAddresses"`
	ScrapeJavaScript     bool     `json:"scrapeJavaScript"`

	Organization Organization `json:"organization"`
	// TODO: plugins
	// TODO: platforms
	ProcessingIssues int `json:"processingIssues"`
	// TODO: defaultEnvironment

	Team  Team   `json:"team"`
	Teams []Team `json:"teams"`
}

// ProjectSummary represents the summary of a Sentry project.
// https://github.com/getsentry/sentry/blob/9.0.0/src/sentry/api/serializers/models/project.py#L258
type ProjectSummary struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	IsBookmarked bool   `json:"isBookmarked"`
	IsMember     bool   `json:"isMember"`
	HasAccess    bool   `json:"hasAccess"`

	DateCreated time.Time `json:"dateCreated"`
	FirstEvent  time.Time `json:"firstEvent"`

	Platform  *string  `json:"platform"`
	Platforms []string `json:"platforms"`

	Team  *ProjectSummaryTeam  `json:"team"`
	Teams []ProjectSummaryTeam `json:"teams"`
	// TODO: deploys
}

// ProjectSummaryTeam represents a team in a ProjectSummary.
// https://github.com/getsentry/sentry/blob/9.0.0/src/sentry/api/serializers/models/project.py#L223
type ProjectSummaryTeam struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// ProjectService provides methods for accessing Sentry project API endpoints.
// https://docs.sentry.io/api/projects/
type ProjectService struct {
	sling *sling.Sling
}

func newProjectService(sling *sling.Sling) *ProjectService {
	return &ProjectService{
		sling: sling,
	}
}

// List projects available.
// https://docs.sentry.io/api/projects/get-project-index/
func (s *ProjectService) List() ([]Project, *http.Response, error) {
	projects := new([]Project)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("projects/").Receive(projects, apiError)
	return *projects, resp, relevantError(err, *apiError)
}

// Get details on an individual project.
// https://docs.sentry.io/api/projects/get-project-details/
func (s *ProjectService) Get(organizationSlug string, slug string) (*Project, *http.Response, error) {
	project := new(Project)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("projects/"+organizationSlug+"/"+slug+"/").Receive(project, apiError)
	return project, resp, relevantError(err, *apiError)
}

// CreateProjectParams are the parameters for ProjectService.Create.
type CreateProjectParams struct {
	Name     string `json:"name,omitempty"`
	Slug     string `json:"slug,omitempty"`
	Platform string `json:"platform,omitempty"`
}

// Create a new project bound to a team.
// https://docs.sentry.io/api/teams/post-team-project-index/
func (s *ProjectService) Create(organizationSlug string, teamSlug string, params *CreateProjectParams) (*Project, *http.Response, error) {
	project := new(Project)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("teams/"+organizationSlug+"/"+teamSlug+"/projects/").BodyJSON(params).Receive(project, apiError)
	return project, resp, relevantError(err, *apiError)
}

// UpdateProjectParams are the parameters for ProjectService.Update.
type UpdateProjectParams struct {
	Name            string                 `json:"name,omitempty"`
	Slug            string                 `json:"slug,omitempty"`
	Platform        string                 `json:"platform,omitempty"`
	IsBookmarked    *bool                  `json:"isBookmarked,omitempty"`
	DigestsMinDelay *int                   `json:"digestsMinDelay,omitempty"`
	DigestsMaxDelay *int                   `json:"digestsMaxDelay,omitempty"`
	Options         map[string]interface{} `json:"options,omitempty"`
	AllowedDomains  []string               `json:"allowedDomains,omitempty"`
}

// Update various attributes and configurable settings for a given project.
// https://docs.sentry.io/api/projects/put-project-details/
func (s *ProjectService) Update(organizationSlug string, slug string, params *UpdateProjectParams) (*Project, *http.Response, error) {
	project := new(Project)
	apiError := new(APIError)
	resp, err := s.sling.New().Put("projects/"+organizationSlug+"/"+slug+"/").BodyJSON(params).Receive(project, apiError)
	return project, resp, relevantError(err, *apiError)
}

// Delete a project.
// https://docs.sentry.io/api/projects/delete-project-details/
func (s *ProjectService) Delete(organizationSlug string, slug string) (*http.Response, error) {
	apiError := new(APIError)
	resp, err := s.sling.New().Delete("projects/"+organizationSlug+"/"+slug+"/").Receive(nil, apiError)
	return resp, relevantError(err, *apiError)
}
