package mailerlite

import (
	"context"
	"net/http"
)

// StatsService handles communication with the stat related
// methods of the MailerLite API.
//
// MailerLite API docs: https://developers.mailerlite.com/reference/stats
type StatsService service

// Stats represents account statistics.
type Stats struct {
	Subscribed   int     `json:"subscribed"`
	Unsubscribed int     `json:"unsubscribed"`
	Campaigns    int     `json:"campaigns"`
	SentEmails   int     `json:"sent_emails"`
	OpenRate     float64 `json:"open_rate"`
	ClickRate    float64 `json:"click_rate"`
	BounceRate   float64 `json:"bounce_rate"`
}

// StatGetOptions specifies the optional parameters to the StatsService.Get method.
type StatGetOptions struct {
	// Specify UNIX timestamp if you want to receive stats values at the specific point in the past.
	// Default: none
	Timestamp int32 `url:"visibility,omitempty"`
}

// Get basic stats for of account, such as subscribers, open/click rates and so on.
//
// MailerLite API docs: https://developers.mailerlite.com/reference/stats
func (s *StatsService) Get(ctx context.Context, opts *StatGetOptions) (*Stats, *Response, error) {
	u := "stats"

	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var stats Stats
	resp, err := s.client.Do(req, &stats)
	if err != nil {
		return nil, resp, err
	}

	return &stats, resp, nil
}
