package mailerlite

import (
	"context"
	"fmt"
	"net/http"
)

// FieldsService handles communication with the field related
// methods of the MailerLite API.
//
// MailerLite API docs: https://developers.mailerlite.com/reference/all-fields
type FieldsService service

// Field represents a custom field in a subscriber profile.
type Field struct {
	ID          WeakInt   `json:"id"`
	Title       string    `json:"title"`
	Key         string    `json:"key"`
	Type        FieldType `json:"type"`
	DateUpdated Timestamp `json:"date_updated"`
	DateCreated Timestamp `json:"date_created"`
}

// FieldType represents the type of data a field can store.
type FieldType string

const (
	Text   FieldType = "TEXT"
	Number FieldType = "NUMBER"
	Date   FieldType = "DATE"
)

// List all fields.
//
// MailerLite API docs: https://developers.mailerlite.com/reference/all-fields
func (s *FieldsService) List(ctx context.Context) ([]Field, *Response, error) {
	u := "fields"

	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var fields []Field
	resp, err := s.client.Do(req, &fields)
	if err != nil {
		return nil, resp, err
	}

	return fields, resp, nil
}

// NewField represents a new field to be created.
type NewField struct {
	Title string    `json:"title,omitempty"`
	Type  FieldType `json:"type,omitempty"`
}

// Create a new field.
//
// MailerLite API docs: https://developers.mailerlite.com/reference/create-field
func (s *FieldsService) Create(ctx context.Context, newField NewField) (*Field, *Response, error) {
	u := "fields"

	req, err := s.client.NewRequest(ctx, http.MethodPost, u, newField)
	if err != nil {
		return nil, nil, err
	}

	var field Field
	resp, err := s.client.Do(req, &field)
	if err != nil {
		return nil, resp, err
	}

	return &field, resp, nil
}

// FieldUpdate represents information that needs to be updated in a field.
type FieldUpdate struct {
	Title string `json:"title,omitempty"`
}

// Update a field.
//
// MailerLite API docs: https://developers.mailerlite.com/reference/update-field
func (s *FieldsService) Update(ctx context.Context, id int, update FieldUpdate) (*Field, *Response, error) {
	u := fmt.Sprintf("fields/%d", id)

	req, err := s.client.NewRequest(ctx, http.MethodPut, u, update)
	if err != nil {
		return nil, nil, err
	}

	var field Field
	resp, err := s.client.Do(req, &field)
	if err != nil {
		return nil, resp, err
	}

	return &field, resp, nil
}

// Delete a field.
//
// MailerLite API docs: https://developers.mailerlite.com/reference/all-fields
func (s *FieldsService) Delete(ctx context.Context, id int) (*Response, error) {
	u := fmt.Sprintf("fields/%d", id)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
