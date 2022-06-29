package mailerlite

import (
	"net/http"
	"net/url"
	"sync"
)

const (
	defaultBaseURL = "https://api.mailerlite.com/api/v2"
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

	rateLimits [categories]Rate // Rate limits for the client as determined by the most recent API calls.
	rateMu     sync.Mutex

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
func NewClient(apiKey string, opts ...ClientOption) *Client {
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

	return c
}
