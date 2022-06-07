package sentry

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/dghubble/sling"
	"github.com/google/go-querystring/query"
	"github.com/peterhellberg/link"
)

const (
	defaultBaseURL = "https://sentry.io/api/"
	userAgent      = "go-sentry"

	APIVersion = "0"
)

var errNonNilContext = errors.New("context must be non-nil")

// Client for Sentry API.
type Client struct {
	client *http.Client

	// BaseURL for API requests.
	BaseURL *url.URL

	// User agent used when communicating with Sentry.
	UserAgent string

	// TODO: Remove sling
	sling *sling.Sling

	// Common struct used by all services.
	common service

	// Services
	IssueAlerts         *IssueAlertsService
	MetricAlerts        *MetricAlertsService
	OrganizationMembers *OrganizationMembersService
	Organizations       *OrganizationsService
	ProjectKeys         *ProjectKeysService
	ProjectOwnerships   *ProjectOwnershipsService
	ProjectPlugins      *ProjectPluginsService
	Projects            *ProjectsService
	Teams               *TeamsService
}

type service struct {
	client *Client
}

// NewClient returns a new Sentry API client.
// If a nil httpClient is provided, the http.DefaultClient will be used.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)

	base := sling.New().Client(httpClient)

	c := &Client{
		sling: base,

		client:  httpClient,
		BaseURL: baseURL,
	}
	c.common.client = c
	c.IssueAlerts = (*IssueAlertsService)(&c.common)
	c.MetricAlerts = (*MetricAlertsService)(&c.common)
	c.OrganizationMembers = (*OrganizationMembersService)(&c.common)
	c.Organizations = (*OrganizationsService)(&c.common)
	c.ProjectKeys = (*ProjectKeysService)(&c.common)
	c.ProjectOwnerships = (*ProjectOwnershipsService)(&c.common)
	c.ProjectPlugins = (*ProjectPluginsService)(&c.common)
	c.Projects = (*ProjectsService)(&c.common)
	c.Teams = (*TeamsService)(&c.common)
	return c
}

// NewOnPremiseClient returns a new Sentry API client with the provided base URL.
// Note that the base URL must be in the format "http(s)://[hostname]/api/".
// If the base URL does not have the suffix "/api/", it will be added automatically.
// If a nil httpClient is provided, the http.DefaultClient will be used.
func NewOnPremiseClient(baseURL string, httpClient *http.Client) (*Client, error) {
	baseEndpoint, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	if !strings.HasSuffix(baseEndpoint.Path, "/") {
		baseEndpoint.Path += "/"
	}
	if !strings.HasSuffix(baseEndpoint.Path, "/api/") {
		baseEndpoint.Path += "api/"
	}

	c := NewClient(httpClient)
	c.BaseURL = baseEndpoint
	return c, nil
}

type ListCursorParams struct {
	// A cursor, as given in the Link header.
	// If specified, the query continues the search using this cursor.
	Cursor string `url:"cursor,omitempty"`
}

func addQuery(s string, params interface{}) (string, error) {
	v := reflect.ValueOf(params)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(params)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

// NewRequest creates an API request.
func (c *Client) NewRequest(method, urlRef string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}

	u, err := c.BaseURL.Parse(urlRef)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

// Response is a Sentry API response. This wraps the standard http.Response
// and provides convenient access to things like pagination links and rate limits.
type Response struct {
	*http.Response

	// For APIs that support cursor pagination, the following field will be populated
	// to point to the next page if more results are available.
	// Set ListCursorParams.Cursor to this value when calling the endpoint again.
	Cursor string

	// TODO: Parse rate limit
}

func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	response.populatePaginationCursor()
	return response
}

func (r *Response) populatePaginationCursor() {
	rels := link.ParseResponse(r.Response)
	if nextRel, ok := rels["next"]; ok && nextRel.Extra["results"] == "true" {
		r.Cursor = nextRel.Extra["cursor"]
	}
}

func (c *Client) BareDo(ctx context.Context, req *http.Request) (*Response, error) {
	if ctx == nil {
		return nil, errNonNilContext
	}

	resp, err := c.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			return nil, err
		}
	}

	response := newResponse(resp)

	err = CheckResponse(resp)
	return response, err
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.BareDo(ctx, req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	switch v := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	default:
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if decErr == io.EOF {
			decErr = nil
		}
		if decErr != nil {
			err = decErr
		}
	}
	return resp, err
}

type ErrorResponse struct {
	Response *http.Response
	Detail   string `json:"detail"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf(
		"%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Detail)
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	// TODO: Handle API errors

	return errorResponse
}

// Avatar represents an avatar.
type Avatar struct {
	UUID *string `json:"avatarUuid"`
	Type string  `json:"avatarType"`
}
