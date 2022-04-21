package sentry

import (
	"encoding/json"
	"net/http"

	"github.com/dghubble/sling"
)

// ProjectFilter represents inbounding filters applied to a project.
type ProjectFilter struct {
	ID     string          `json:"id"`
	Active json.RawMessage `json:"active"`
}

// ProjectOwnershipService provides methods for accessing Sentry project
// filters API endpoints.
type ProjectFilterService struct {
	sling *sling.Sling
}

func newProjectFilterService(sling *sling.Sling) *ProjectFilterService {
	return &ProjectFilterService{
		sling: sling,
	}
}

// Get the filters.
func (s *ProjectFilterService) Get(organizationSlug string, projectSlug string) ([]*ProjectFilter, *http.Response, error) {
	var filters []*ProjectFilter
	apiError := new(APIError)
	resp, err := s.sling.New().Get("projects/"+organizationSlug+"/"+projectSlug+"/filters/").Receive(&filters, apiError)
	return filters, resp, relevantError(err, *apiError)
}

// FilterConfig represents configuration for project filter
type FilterConfig struct {
	BrowserExtension bool
	LegacyBrowsers   []string
}

// GetFilterConfig retrieves filter configuration.
func (s *ProjectFilterService) GetFilterConfig(organizationSlug string, projectSlug string) (*FilterConfig, *http.Response, error) {
	filters, resp, err := s.Get(organizationSlug, projectSlug)
	if err != nil {
		return nil, resp, err
	}

	var filterConfig FilterConfig

	for _, filter := range filters {
		switch filter.ID {
		case "browser-extensions":
			if string(filter.Active) == "true" {
				filterConfig.BrowserExtension = true
			}

		case "legacy-browsers":
			if string(filter.Active) != "false" {
				err = json.Unmarshal(filter.Active, &filterConfig.LegacyBrowsers)
				if err != nil {
					return nil, resp, err
				}
			}
		}
	}

	return &filterConfig, resp, err
}

// BrowserExtensionParams defines parameters for browser extension request
type BrowserExtensionParams struct {
	Active bool `json:"active"`
}

// UpdateBrowserExtensions updates configuration for browser extension filter
func (s *ProjectFilterService) UpdateBrowserExtensions(organizationSlug string, projectSlug string, active bool) (*http.Response, error) {
	apiError := new(APIError)
	params := BrowserExtensionParams{active}
	resp, err := s.sling.New().Put("projects/"+organizationSlug+"/"+projectSlug+"/filters/browser-extensions/").BodyJSON(params).Receive(nil, apiError)
	return resp, relevantError(err, *apiError)
}

// LegactBrowserParams defines parameters for legacy browser request
type LegactBrowserParams struct {
	Browsers []string `json:"subfilters"`
}

// UpdateLegacyBrowser updates configuration for legacy browser filters
func (s *ProjectFilterService) UpdateLegacyBrowser(organizationSlug string, projectSlug string, browsers []string) (*http.Response, error) {
	apiError := new(APIError)
	params := LegactBrowserParams{browsers}
	resp, err := s.sling.New().Put("projects/"+organizationSlug+"/"+projectSlug+"/filters/legacy-browsers/").BodyJSON(params).Receive(nil, apiError)
	return resp, relevantError(err, *apiError)
}
