package mailerlite

import (
	"context"
	"fmt"
	"net/http"
)

// SubscribersService handles communication with the subscriber related
// methods of the MailerLite API.
//
// MailerLite API docs: https://developers.mailerlite.com/reference/subscribers
type SubscribersService service

// Subscriber represents an email subscriber.
type Subscriber struct {
	ID                    int               `json:"id"`
	Name                  string            `json:"name"`
	Email                 string            `json:"email"`
	Sent                  int               `json:"sent"`
	Opened                int               `json:"opened"`
	Clicked               int               `json:"clicked"`
	Type                  SubscriptionType  `json:"type"`
	CountryID             string            `json:"country_id"`
	SignupIP              *string           `json:"signup_ip"`
	SignupTimestamp       *string           `json:"signup_timestamp"`
	ConfirmationIP        *string           `json:"confirmation_ip"`
	ConfirmationTimestamp *string           `json:"confirmation_timestamp"`
	Fields                []SubscriberField `json:"fields"`
	DateSubscribe         Timestamp         `json:"date_subscribe"`
	DateUnsubscribe       *Timestamp        `json:"date_unsubscribe"`
	DateCreated           Timestamp         `json:"date_created"`
	DateUpdated           *Timestamp        `json:"date_updated"`
}

// SubscriptionType represents the current state of the subscription.
type SubscriptionType string

const (
	Unsubscribed SubscriptionType = "unsubscribed"
	Active       SubscriptionType = "active"
	Unconfirmed  SubscriptionType = "unconfirmed"
	Bounced      SubscriptionType = "bounced"
	Junk         SubscriptionType = "junk"
)

// SubscriberField represents a custom field and its value.
type SubscriberField struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

// SubscriberListOptions specifies the optional parameters to the
// SubscribersService.List method.
type SubscriberListOptions struct {
	Type SubscriptionType `url:"type,omitempty"`

	ListOptions
}

// List all subscribers.
//
// MailerLite API docs: https://developers.mailerlite.com/reference/subscribers
func (s *SubscribersService) List(ctx context.Context, opts *SubscriberListOptions) ([]Subscriber, *Response, error) {
	u := "subscribers"

	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var subscribers []Subscriber
	resp, err := s.client.Do(req, &subscribers)
	if err != nil {
		return nil, resp, err
	}

	return subscribers, resp, nil
}

// TODO: create

// Get fetches a subscriber.
//
// MailerLite API docs: https://developers.mailerlite.com/reference/single-subscriber
func (s *SubscribersService) Get(ctx context.Context, email string) (*Subscriber, *Response, error) {
	u := fmt.Sprintf("subscribers/%s", email)

	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var subscriber Subscriber
	resp, err := s.client.Do(req, &subscriber)
	if err != nil {
		return nil, resp, err
	}

	return &subscriber, resp, nil
}

// TODO: update
