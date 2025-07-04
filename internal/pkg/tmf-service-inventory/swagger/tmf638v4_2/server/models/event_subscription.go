// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EventSubscription Sets the communication endpoint address the service instance must use to deliver notification information
//
// swagger:model EventSubscription
type EventSubscription struct {

	// The callback being registered.
	// Required: true
	Callback string `json:"callback" bson:"callback,omitempty"`

	// Id of the listener
	// Required: true
	ID string `json:"id" bson:"id,omitempty"`

	// additional data to be passed
	Query *string `json:"query,omitempty" bson:"query,omitempty"`
}

// Validate validates this event subscription
func (m *EventSubscription) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCallback(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EventSubscription) validateCallback(formats strfmt.Registry) error {

	if err := validate.RequiredString("callback", "body", m.Callback); err != nil {
		return err
	}

	return nil
}

func (m *EventSubscription) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this event subscription based on context it is used
func (m *EventSubscription) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EventSubscription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventSubscription) UnmarshalBinary(b []byte) error {
	var res EventSubscription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
