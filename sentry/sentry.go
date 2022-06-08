package sentry

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/peterhellberg/link"
)

const (
	defaultBaseURL = "https://sentry.io/api/"
	userAgent      = "go-sentry"

	// https://docs.sentry.io/api/ratelimits/
	headerRateLimit               = "X-Sentry-Rate-Limit-Limit"
	headerRateRemaining           = "X-Sentry-Rate-Limit-Remaining"
	headerRateReset               = "X-Sentry-Rate-Limit-Reset"
	headerRateConcurrentLimit     = "X-Sentry-Rate-Limit-ConcurrentLimit"
	headerRateConcurrentRemaining = "X-Sentry-Rate-Limit-ConcurrentRemaining"
)

var errNonNilContext = errors.New("context must be non-nil")

// Client for Sentry API.
type Client struct {
	client *http.Client

	// BaseURL for API requests.
	BaseURL *url.URL

	// User agent used when communicating with Sentry.
	UserAgent string

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

	c := &Client{
		client:    httpClient,
		BaseURL:   baseURL,
		UserAgent: userAgent,
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

	Rate Rate
}

func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	response.Rate = parseRate(r)
	response.populatePaginationCursor()
	return response
}

func (r *Response) populatePaginationCursor() {
	rels := link.ParseResponse(r.Response)
	if nextRel, ok := rels["next"]; ok && nextRel.Extra["results"] == "true" {
		r.Cursor = nextRel.Extra["cursor"]
	}
}

// parseRate parses the rate limit headers.
func parseRate(r *http.Response) Rate {
	var rate Rate
	if limit := r.Header.Get(headerRateLimit); limit != "" {
		rate.Limit, _ = strconv.Atoi(limit)
	}
	if remaining := r.Header.Get(headerRateRemaining); remaining != "" {
		rate.Remaining, _ = strconv.Atoi(remaining)
	}
	if reset := r.Header.Get(headerRateReset); reset != "" {
		if v, _ := strconv.ParseInt(reset, 10, 64); v != 0 {
			rate.Reset = time.Unix(v, 0).UTC()
		}
	}
	if concurrentLimit := r.Header.Get(headerRateConcurrentLimit); concurrentLimit != "" {
		rate.ConcurrentLimit, _ = strconv.Atoi(concurrentLimit)
	}
	if concurrentRemaining := r.Header.Get(headerRateConcurrentRemaining); concurrentRemaining != "" {
		rate.ConcurrentRemaining, _ = strconv.Atoi(concurrentRemaining)
	}

	return rate
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
		dec := json.NewDecoder(resp.Body)
		dec.UseNumber()
		decErr := dec.Decode(v)
		if decErr == io.EOF {
			decErr = nil
		}
		if decErr != nil {
			err = decErr
		}
	}
	return resp, err
}

func (c *Client) checkRateLimit() *RateLimitError {
	return nil
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

type RateLimitError struct {
	Rate     Rate
	Response *http.Response
	Detail   string
}

func (r *RateLimitError) Error() string {
	return fmt.Sprintf(
		"%v %v: %d %v %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Detail, fmt.Sprintf("[rate reset in %v]", time.Until(r.Rate.Reset)))
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		apiError := new(APIError)
		json.Unmarshal(data, apiError)
		if apiError.Empty() {
			errorResponse.Detail = strings.TrimSpace(string(data))
		} else {
			errorResponse.Detail = apiError.Detail()
		}
	}
	// Re-populate error response body.
	r.Body = ioutil.NopCloser(bytes.NewBuffer(data))

	switch {
	case r.StatusCode == http.StatusTooManyRequests &&
		(r.Header.Get(headerRateRemaining) == "0" || r.Header.Get(headerRateConcurrentRemaining) == "0"):
		return &RateLimitError{
			Rate:     parseRate(r),
			Response: errorResponse.Response,
			Detail:   errorResponse.Detail,
		}
	}

	return errorResponse
}

// Rate represents the rate limit for the current client.
type Rate struct {
	// The maximum number of requests allowed within the window.
	Limit int

	// The number of requests this caller has left on this endpoint within the current window
	Remaining int

	// The time when the next rate limit window begins and the count resets, measured in UTC seconds from epoch
	Reset time.Time

	// The maximum number of concurrent requests allowed within the window
	ConcurrentLimit int

	// The number of concurrent requests this caller has left on this endpoint within the current window
	ConcurrentRemaining int
}

// Bool returns a pointer to the bool value passed in.
func Bool(v bool) *bool { return &v }

// Int returns a pointer to the int value passed in.
func Int(v int) *int { return &v }

// String returns a pointer to the string value passed in.
func String(v string) *string { return &v }

// Time returns a pointer to the time.Time value passed in.
func Time(v time.Time) *time.Time { return &v }
