package sentry

import (
	"net/http"
	"time"

	"github.com/dghubble/sling"
)

// Project represents a Sentry project.
// Based on https://github.com/getsentry/sentry/blob/cc81fff31d4f2c9cede14ce9c479d6f4f78c5e5b/src/sentry/api/serializers/models/project.py#L137.
type Project struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`

	DateCreated time.Time `json:"dateCreated"`
	FirstEvent  time.Time `json:"firstEvent"`

	IsPublic     bool     `json:"isPublic"`
	IsBookmarked bool     `json:"isBookmarked"`
	CallSign     string   `json:"callSign"`
	Color        string   `json:"color"`
	Features     []string `json:"features"`
	Status       string   `json:"status"`
}

// ProjectService provides methods for accessing Sentry project API endpoints.
// https://docs.sentry.io/api/projects/
type ProjectService struct {
	sling *sling.Sling
}

func newProjectService(sling *sling.Sling) *ProjectService {
	return &ProjectService{
		sling: sling.Path("projects/"),
	}
}

// List projects available.
// https://docs.sentry.io/api/projects/get-project-index/
func (s *ProjectService) List() ([]Project, *http.Response, error) {
	projects := new([]Project)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("").Receive(projects, apiError)
	return *projects, resp, relevantError(err, *apiError)
}
