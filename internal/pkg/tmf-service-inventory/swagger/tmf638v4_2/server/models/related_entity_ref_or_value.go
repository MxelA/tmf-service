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

// RelatedEntityRefOrValue A reference to an entity, where the type of the entity is not known in advance. A related entity defines a entity described by reference or by value linked to a specific entity. The polymorphic attributes @type, @schemaLocation & @referredType are related to the Entity and not the RelatedEntityRefOrValue class itself
//
// swagger:model RelatedEntityRefOrValue
type RelatedEntityRefOrValue struct {

	// When sub-classing, this defines the super-class
	AtBaseType *string `json:"@baseType,omitempty" bson:"@baseType,omitempty"`

	// The actual type of the target instance when needed for disambiguation.
	AtReferredType *string `json:"@referredType,omitempty" bson:"@referredType,omitempty"`

	// A URI to a JSON-Schema file that defines additional attributes and relationships
	// Format: uri
	AtSchemaLocation *strfmt.URI `json:"@schemaLocation,omitempty" bson:"@schemaLocation,omitempty"`

	// When sub-classing, this defines the sub-class Extensible name
	AtType *string `json:"@type,omitempty" bson:"@type,omitempty"`

	// Reference of the related entity.
	Href *string `json:"href,omitempty" bson:"href,omitempty"`

	// Unique identifier of a related entity.
	ID *string `json:"id,omitempty" bson:"id,omitempty"`

	// Name of the related entity.
	Name *string `json:"name,omitempty" bson:"name,omitempty"`

	// role
	// Required: true
	Role string `json:"role" bson:"role,omitempty"`
}

// Validate validates this related entity ref or value
func (m *RelatedEntityRefOrValue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAtSchemaLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RelatedEntityRefOrValue) validateAtSchemaLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.AtSchemaLocation) { // not required
		return nil
	}

	if err := validate.FormatOf("@schemaLocation", "body", "uri", m.AtSchemaLocation.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RelatedEntityRefOrValue) validateRole(formats strfmt.Registry) error {

	if err := validate.RequiredString("role", "body", m.Role); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this related entity ref or value based on context it is used
func (m *RelatedEntityRefOrValue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RelatedEntityRefOrValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RelatedEntityRefOrValue) UnmarshalBinary(b []byte) error {
	var res RelatedEntityRefOrValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
