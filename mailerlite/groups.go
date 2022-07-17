package mailerlite

import (
	"context"
	"fmt"
	"net/http"
)

// GroupsService handles communication with the group related
// methods of the MailerLite API.
//
// MailerLite API docs: https://developers.mailerlite.com/reference/groups
type GroupsService service

// NewSubscriberInGroup represents a new subscriber in a group.
type NewSubscriberInGroup struct {
	Email          string            `json:"email,omitempty"`
	Name           string            `json:"name,omitempty"`
	Fields         map[string]string `json:"fields,omitempty"`
	Resubscribe    *bool             `json:"resubscribe,omitempty"`
	AutoResponders *bool             `json:"autoresponders,omitempty"`
	Type           SubscriptionType  `json:"type,omitempty"`
}

// Add a subscriber to a group.
//
// MailerLite API docs: https://developers.mailerlite.com/reference/add-single-subscriber
func (s *GroupsService) AddSubscriber(ctx context.Context, id int, newSubscriber NewSubscriberInGroup) (*Subscriber, *Response, error) {
	u := fmt.Sprintf("groups/%d/subscribers", id)

	req, err := s.client.NewRequest(ctx, http.MethodPost, u, newSubscriber)
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
