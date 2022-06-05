package sentry

import (
	"net/http"
	"time"

	"github.com/dghubble/sling"
)

// OrganizationStatus represents a Sentry organization's status.
type OrganizationStatus struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// OrganizationQuota represents a Sentry organization's quota.
type OrganizationQuota struct {
	MaxRate         int `json:"maxRate"`
	MaxRateInterval int `json:"maxRateInterval"`
	AccountLimit    int `json:"accountLimit"`
	ProjectLimit    int `json:"projectLimit"`
}

// OrganizationAvailableRole represents a Sentry organization's available role.
type OrganizationAvailableRole struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Organization represents a Sentry organization.
// Based on https://github.com/getsentry/sentry/blob/22.5.0/src/sentry/api/serializers/models/organization.py#L110-L120
type Organization struct {
	// Basic
	ID                       string             `json:"id"`
	Slug                     string             `json:"slug"`
	Status                   OrganizationStatus `json:"status"`
	Name                     string             `json:"name"`
	DateCreated              time.Time          `json:"dateCreated"`
	IsEarlyAdopter           bool               `json:"isEarlyAdopter"`
	Require2FA               bool               `json:"require2FA"`
	RequireEmailVerification bool               `json:"requireEmailVerification"`
	Avatar                   Avatar             `json:"avatar"`
	Features                 []string           `json:"features"`
}

// DetailedOrganization represents detailed information about a Sentry organization.
// Based on https://github.com/getsentry/sentry/blob/22.5.0/src/sentry/api/serializers/models/organization.py#L263-L288
type DetailedOrganization struct {
	// Basic
	ID                       string             `json:"id"`
	Slug                     string             `json:"slug"`
	Status                   OrganizationStatus `json:"status"`
	Name                     string             `json:"name"`
	DateCreated              time.Time          `json:"dateCreated"`
	IsEarlyAdopter           bool               `json:"isEarlyAdopter"`
	Require2FA               bool               `json:"require2FA"`
	RequireEmailVerification bool               `json:"requireEmailVerification"`
	Avatar                   Avatar             `json:"avatar"`
	Features                 []string           `json:"features"`

	// Detailed
	// TODO: experiments
	Quota                OrganizationQuota           `json:"quota"`
	IsDefault            bool                        `json:"isDefault"`
	DefaultRole          string                      `json:"defaultRole"`
	AvailableRoles       []OrganizationAvailableRole `json:"availableRoles"`
	OpenMembership       bool                        `json:"openMembership"`
	AllowSharedIssues    bool                        `json:"allowSharedIssues"`
	EnhancedPrivacy      bool                        `json:"enhancedPrivacy"`
	DataScrubber         bool                        `json:"dataScrubber"`
	DataScrubberDefaults bool                        `json:"dataScrubberDefaults"`
	SensitiveFields      []string                    `json:"sensitiveFields"`
	SafeFields           []string                    `json:"safeFields"`
	StoreCrashReports    int                         `json:"storeCrashReports"`
	AttachmentsRole      string                      `json:"attachmentsRole"`
	DebugFilesRole       string                      `json:"debugFilesRole"`
	EventsMemberAdmin    bool                        `json:"eventsMemberAdmin"`
	AlertsMemberWrite    bool                        `json:"alertsMemberWrite"`
	ScrubIPAddresses     bool                        `json:"scrubIPAddresses"`
	ScrapeJavaScript     bool                        `json:"scrapeJavaScript"`
	AllowJoinRequests    bool                        `json:"allowJoinRequests"`
	RelayPiiConfig       *string                     `json:"relayPiiConfig"`
	// TODO: trustedRelays
	Access                []string `json:"access"`
	Role                  string   `json:"role"`
	PendingAccessRequests int      `json:"pendingAccessRequests"`
	// TODO: onboardingTasks
}

// OrganizationService provides methods for accessing Sentry organization API endpoints.
// https://docs.sentry.io/api/organizations/
type OrganizationService struct {
	sling *sling.Sling
}

func newOrganizationService(sling *sling.Sling) *OrganizationService {
	return &OrganizationService{
		sling: sling.Path("organizations/"),
	}
}

// ListOrganizationParams are the parameters for OrganizationService.List.
type ListOrganizationParams struct {
	Cursor string `url:"cursor,omitempty"`
}

// List organizations available to the authenticated session.
// https://docs.sentry.io/api/organizations/get-organization-index/
func (s *OrganizationService) List(params *ListOrganizationParams) ([]Organization, *http.Response, error) {
	organizations := new([]Organization)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("").QueryStruct(params).Receive(organizations, apiError)
	return *organizations, resp, relevantError(err, *apiError)
}

// CreateOrganizationParams are the parameters for OrganizationService.Create.
type CreateOrganizationParams struct {
	Name       string `json:"name,omitempty"`
	Slug       string `json:"slug,omitempty"`
	AgreeTerms *bool  `json:"agreeTerms,omitempty"`
}

// Get a Sentry organization.
// https://docs.sentry.io/api/organizations/get-organization-details/
func (s *OrganizationService) Get(slug string) (*DetailedOrganization, *http.Response, error) {
	org := new(DetailedOrganization)
	apiError := new(APIError)
	resp, err := s.sling.New().Get(slug+"/").Receive(org, apiError)
	return org, resp, relevantError(err, *apiError)
}

// Create a new Sentry organization.
// https://docs.sentry.io/api/organizations/post-organization-index/
func (s *OrganizationService) Create(params *CreateOrganizationParams) (*Organization, *http.Response, error) {
	org := new(Organization)
	apiError := new(APIError)
	resp, err := s.sling.New().Post("").BodyJSON(params).Receive(org, apiError)
	return org, resp, relevantError(err, *apiError)
}

// UpdateOrganizationParams are the parameters for OrganizationService.Update.
type UpdateOrganizationParams struct {
	Name string `json:"name,omitempty"`
	Slug string `json:"slug,omitempty"`
}

// Update a Sentry organization.
// https://docs.sentry.io/api/organizations/put-organization-details/
func (s *OrganizationService) Update(slug string, params *UpdateOrganizationParams) (*Organization, *http.Response, error) {
	org := new(Organization)
	apiError := new(APIError)
	resp, err := s.sling.New().Put(slug+"/").BodyJSON(params).Receive(org, apiError)
	return org, resp, relevantError(err, *apiError)
}

// Delete a Sentry organization.
func (s *OrganizationService) Delete(slug string) (*http.Response, error) {
	apiError := new(APIError)
	resp, err := s.sling.New().Delete(slug+"/").Receive(nil, apiError)
	return resp, relevantError(err, *apiError)
}
