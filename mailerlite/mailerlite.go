package mailerlite

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/google/go-querystring/query"
)

const (
	defaultBaseURL = "https://api.mailerlite.com/api/v2/"
	userAgent      = "go-mailerlite"

	headerAPIKey = "X-MailerLite-ApiKey"
)

// A Client manages communication with the MailerLite API.
type Client struct {
	httpClient *http.Client // HTTP client used to communicate with the API.

	// Base URL for API requests.
	// Defaults to the public MailerLite API.
	// BaseURL should always be specified with a trailing slash.
	baseURL *url.URL

	// User agent used when communicating with the MailerLite API.
	userAgent string

	// API key that authenticates each request sent to the API.
	apiKey string

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the MailerLite API.
	Campaigns   *CampaignsService
	Segments    *SegmentsService
	Subscribers *SubscribersService
	Groups      *GroupsService
	Fields      *FieldsService
	Webhooks    *WebhooksService
	Stats       *StatsService
	Settings    *SettingsService
}

type service struct {
	client *Client
}

// ClientOption configures a Client.
type ClientOption interface {
	apply(c *Client)
}

type clientOptionFunc func(c *Client)

func (fn clientOptionFunc) apply(c *Client) {
	fn(c)
}

// HTTPClient configures a Client to use a specific HTTP client for communicating with the API.
func HTTPClient(httpClient *http.Client) ClientOption {
	return clientOptionFunc(func(c *Client) {
		if httpClient == nil {
			return
		}

		c.httpClient = httpClient
	})
}

// BaseURL tells a Client where to send API requests to.
// Defaults to the public MailerLite API, but can be set to a different URL.
// BaseURL should always be specified with a trailing slash.
func BaseURL(baseURL *url.URL) ClientOption {
	return clientOptionFunc(func(c *Client) {
		if baseURL == nil {
			return
		}

		c.baseURL = baseURL
	})
}

// UserAgent configures a Client to send API requests with a specific user agent header.
func UserAgent(userAgent string) ClientOption {
	return clientOptionFunc(func(c *Client) {
		if userAgent == "" {
			return
		}

		c.userAgent = userAgent
	})
}

// NewClient returns a new MailerLite API client.
func NewClient(apiKey string, opts ...ClientOption) (*Client, error) {
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		httpClient: http.DefaultClient,
		baseURL:    baseURL,
		userAgent:  userAgent,
		apiKey:     apiKey,
	}

	for _, opt := range opts {
		opt.apply(c)
	}

	if !strings.HasSuffix(c.baseURL.Path, "/") {
		return nil, fmt.Errorf("base URL must have a trailing slash, but %q does not", c.baseURL)
	}

	c.common.client = c

	// Services
	c.Campaigns = (*CampaignsService)(&c.common)
	c.Segments = (*SegmentsService)(&c.common)
	c.Subscribers = (*SubscribersService)(&c.common)
	c.Groups = (*GroupsService)(&c.common)
	c.Fields = (*FieldsService)(&c.common)
	c.Webhooks = (*WebhooksService)(&c.common)
	c.Stats = (*StatsService)(&c.common)
	c.Settings = (*SettingsService)(&c.common)

	return c, nil
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the baseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(ctx context.Context, method string, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.baseURL.Parse(urlStr)
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

	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set(headerAPIKey, c.apiKey)

	req.Header.Set("Accept", "application/json")

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if c.userAgent != "" {
		req.Header.Set("User-Agent", c.userAgent)
	}

	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer interface,
// the raw response body will be written to v, without attempting to first
// decode it. If v is nil, and no error hapens, the response is returned as is.
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.BareDo(req)
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
			decErr = nil // ignore EOF errors caused by empty response body
		}

		if decErr != nil {
			err = decErr
		}
	}

	return resp, err
}

// Response is a MailerLite API response. This wraps the standard http.Response
// returned from MailerLite and provides convenient access to things like
// rate limit information.
type Response struct {
	*http.Response
}

// newResponse creates a new Response for the provided http.Response.
// r must not be nil.
func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}

	return response
}

// BareDo sends an API request and lets you handle the api response. If an error
// or API Error occurs, the error will contain more information. Otherwise you
// are supposed to read and close the response's Body. If rate limit is exceeded
// and reset time is in the future, BareDo returns *RateLimitError immediately
// without making a network API call.
func (c *Client) BareDo(req *http.Request) (*Response, error) {
	// TODO: rate limit

	resp, err := c.httpClient.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-req.Context().Done():
			return nil, req.Context().Err()
		default:
		}

		return nil, err
	}

	response := newResponse(resp)

	err = CheckResponse(resp)
	if err != nil {
		defer resp.Body.Close()
	}

	return response, err
}

// CheckResponse checks the API response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range.
// API error responses are expected to have response
// body, and a JSON response body that maps to ErrorResponse.
//
// The error type will be *RateLimitError for rate limit exceeded errors,
// *AcceptedError for 202 Accepted status codes,
// and *TwoFactorAuthError for two-factor authentication errors.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}

	return errorResponse
}

// ErrorResponse wraps the returned HTTP response and error.
type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error

	Err Error `json:"error"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Err.Message, r.Err.Code,
	)
}

// Error contains details about what went wrong during a failed API request.
type Error struct {
	Code    int    `json:"code"`    // Error code (optional)
	Message string `json:"message"` // Human-readable error message
}

func (e Error) Error() string {
	if e.Code == 0 {
		return fmt.Sprintf("%s (code: %d)", e.Message, e.Code)
	}

	return e.Message
}

// addOptions adds the parameters in opts as URL query parameters to s. opts
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opts interface{}) (string, error) {
	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opts)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}
