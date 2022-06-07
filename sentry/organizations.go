package sentry

import (
	"context"
	"fmt"
	"time"
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

// OrganizationsService provides methods for accessing Sentry organization API endpoints.
// https://docs.sentry.io/api/organizations/
type OrganizationsService service

// List organizations available to the authenticated session.
// https://docs.sentry.io/api/organizations/list-your-organizations/
func (s *OrganizationsService) List(ctx context.Context, params *ListCursorParams) ([]*Organization, *Response, error) {
	u, err := addQuery("0/organizations/", params)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	orgs := []*Organization{}
	resp, err := s.client.Do(ctx, req, &orgs)
	if err != nil {
		return nil, resp, err
	}
	return orgs, resp, nil
}

// Get a Sentry organization.
// https://docs.sentry.io/api/organizations/retrieve-an-organization/
func (s *OrganizationsService) Get(ctx context.Context, slug string) (*DetailedOrganization, *Response, error) {
	u := fmt.Sprintf("0/organizations/%v/", slug)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	org := new(DetailedOrganization)
	resp, err := s.client.Do(ctx, req, org)
	if err != nil {
		return nil, resp, err
	}
	return org, resp, nil
}

// CreateOrganizationParams are the parameters for OrganizationService.Create.
type CreateOrganizationParams struct {
	Name       string `json:"name,omitempty"`
	Slug       string `json:"slug,omitempty"`
	AgreeTerms *bool  `json:"agreeTerms,omitempty"`
}

// Create a new Sentry organization.
func (s *OrganizationsService) Create(ctx context.Context, params *CreateOrganizationParams) (*Organization, *Response, error) {
	u := "0/organizations/"
	req, err := s.client.NewRequest("POST", u, params)
	if err != nil {
		return nil, nil, err
	}

	org := new(Organization)
	resp, err := s.client.Do(ctx, req, org)
	if err != nil {
		return nil, resp, err
	}
	return org, resp, nil
}

// UpdateOrganizationParams are the parameters for OrganizationService.Update.
type UpdateOrganizationParams struct {
	Name string `json:"name,omitempty"`
	Slug string `json:"slug,omitempty"`
}

// Update a Sentry organization.
// https://docs.sentry.io/api/organizations/update-an-organization/
func (s *OrganizationsService) Update(ctx context.Context, slug string, params *UpdateOrganizationParams) (*Organization, *Response, error) {
	u := fmt.Sprintf("0/organizations/%v/", slug)
	req, err := s.client.NewRequest("PUT", u, params)
	if err != nil {
		return nil, nil, err
	}

	org := new(Organization)
	resp, err := s.client.Do(ctx, req, org)
	if err != nil {
		return nil, resp, err
	}
	return org, resp, nil
}

// Delete a Sentry organization.
func (s *OrganizationsService) Delete(ctx context.Context, slug string) (*Response, error) {
	u := fmt.Sprintf("0/organizations/%v/", slug)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
