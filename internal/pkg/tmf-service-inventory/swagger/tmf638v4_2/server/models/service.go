// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Service Service is a base class for defining the Service hierarchy. All Services are characterized as either being possibly visible and usable by a Customer or not. This gives rise to the two subclasses of Service: CustomerFacingService and ResourceFacingService.
//
// swagger:model Service
type Service struct {

	// When sub-classing, this defines the super-class
	AtBaseType *string `json:"@baseType,omitempty" bson:"@baseType,omitempty"`

	// A URI to a JSON-Schema file that defines additional attributes and relationships
	// Format: uri
	AtSchemaLocation *strfmt.URI `json:"@schemaLocation,omitempty" bson:"@schemaLocation,omitempty"`

	// When sub-classing, this defines the sub-class Extensible name
	AtType *string `json:"@type,omitempty" bson:"@type,omitempty"`

	// Is it a customer facing or resource facing service
	Category *string `json:"category,omitempty" bson:"category,omitempty"`

	// Free-text description of the service
	Description *string `json:"description,omitempty" bson:"description,omitempty"`

	// Date when the service ends
	// Format: date-time
	EndDate *strfmt.DateTime `json:"endDate,omitempty" bson:"endDate,omitempty"`

	// A list of external identifiers assoicated with this service
	ExternalIdentifier []*ExternalIdentifier `json:"externalIdentifier,omitempty" bson:"externalIdentifier,omitempty"`

	// A list of feature associated with this service
	Feature []*Feature `json:"feature,omitempty" bson:"feature,omitempty"`

	// If TRUE, this Service has already been started
	HasStarted *bool `json:"hasStarted,omitempty" bson:"hasStarted,omitempty"`

	// Reference of the service
	Href *string `json:"href,omitempty" bson:"href,omitempty"`

	// Unique identifier of the service
	ID *string `json:"id,omitempty" bson:"_id,omitempty"`

	// If true, the service is a ServiceBundle which regroup a service hierachy. If false, the service is a 'atomic' service (hierachy leaf).
	IsBundle *bool `json:"isBundle,omitempty" bson:"isBundle,omitempty"`

	// If FALSE and hasStarted is FALSE, this particular Service has NOT been enabled for use - if FALSE and hasStarted is TRUE then the service has failed
	IsServiceEnabled *bool `json:"isServiceEnabled,omitempty" bson:"isServiceEnabled,omitempty"`

	// If TRUE, this Service can be changed without affecting any other services
	IsStateful *bool `json:"isStateful,omitempty" bson:"isStateful,omitempty"`

	// Name of the service
	Name *string `json:"name,omitempty" bson:"name,omitempty"`

	// A list of notes made on this service
	Note []*Note `json:"note,omitempty" bson:"note,omitempty"`

	// Indicates how a service is currently performing or operating. It is a logical representation of the service operating behaviour and is determined/managed by the service provider.
	OperatingStatus ServiceOperatingStatusType `json:"operatingStatus,omitempty" bson:"operatingStatus,omitempty"`

	// Additional information describing the context of operatingStatus and is determined/managed by the service provider.
	OperatingStatusContextUpdate *ContextUpdate `json:"operatingStatusContextUpdate,omitempty" bson:"operatingStatusContextUpdate,omitempty"`

	// A list of places (Place [*]). Used to define a place useful for the service (for example a geographical place whre the service is installed)
	Place []*RelatedPlaceRefOrValue `json:"place,omitempty" bson:"place,omitempty"`

	// A list of related  entity in relationship with this service
	RelatedEntity []*RelatedEntityRefOrValue `json:"relatedEntity,omitempty" bson:"relatedEntity,omitempty"`

	// A list of related party references (RelatedParty [*]). A related party defines party or party role linked to a specific entity
	RelatedParty []*RelatedParty `json:"relatedParty,omitempty" bson:"relatedParty,omitempty"`

	// A list of characteristics that characterize this service (ServiceCharacteristic [*])
	ServiceCharacteristic []*Characteristic `json:"serviceCharacteristic,omitempty" bson:"serviceCharacteristic,omitempty"`

	// Date when the service was created (whatever its status).
	ServiceDate *string `json:"serviceDate,omitempty" bson:"serviceDate,omitempty"`

	// A list of service order items related to this service
	ServiceOrderItem []*RelatedServiceOrderItem `json:"serviceOrderItem,omitempty" bson:"serviceOrderItem,omitempty"`

	// A list of service relationships (ServiceRelationship [*]). Describes links with other service(s) in the inventory.
	ServiceRelationship []*ServiceRelationship `json:"serviceRelationship,omitempty" bson:"serviceRelationship,omitempty"`

	// The specification from which this service was instantiated
	ServiceSpecification *ServiceSpecificationRef `json:"serviceSpecification,omitempty" bson:"serviceSpecification,omitempty"`

	// Business type of the service
	ServiceType *string `json:"serviceType,omitempty" bson:"serviceType,omitempty"`

	// Date when the service starts
	// Format: date-time
	StartDate *strfmt.DateTime `json:"startDate,omitempty" bson:"startDate,omitempty"`

	// This attribute is an enumerated integer that indicates how the Service is started, such as: 0: Unknown; 1: Automatically by the managed environment; 2: Automatically by the owning device; 3: Manually by the Provider of the Service; 4: Manually by a Customer of the Provider; 5: Any of the above
	StartMode *string `json:"startMode,omitempty" bson:"startMode,omitempty"`

	// The life cycle state of the service, such as designed, reserved, active, etc...
	State ServiceStateType `json:"state,omitempty"`

	// A list of supporting resources (SupportingResource [*]).Note: only Service of type RFS can be associated with Resources
	SupportingResource []*ResourceRef `json:"supportingResource,omitempty" bson:"supportingResource,omitempty"`

	// A list of supporting services (SupportingService [*]). A collection of services that support this service (bundling, link CFS to RFS)
	SupportingService []*ServiceRefOrValue `json:"supportingService,omitempty" bson:"supportingService,omitempty"`
}

// Validate validates this service
func (m *Service) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAtSchemaLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeature(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNote(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatingStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatingStatusContextUpdate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelatedEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelatedParty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceCharacteristic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceOrderItem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceRelationship(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceSpecification(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupportingResource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupportingService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Service) validateAtSchemaLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.AtSchemaLocation) { // not required
		return nil
	}

	if err := validate.FormatOf("@schemaLocation", "body", "uri", m.AtSchemaLocation.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Service) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Service) validateExternalIdentifier(formats strfmt.Registry) error {
	if swag.IsZero(m.ExternalIdentifier) { // not required
		return nil
	}

	for i := 0; i < len(m.ExternalIdentifier); i++ {
		if swag.IsZero(m.ExternalIdentifier[i]) { // not required
			continue
		}

		if m.ExternalIdentifier[i] != nil {
			if err := m.ExternalIdentifier[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("externalIdentifier" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("externalIdentifier" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Service) validateFeature(formats strfmt.Registry) error {
	if swag.IsZero(m.Feature) { // not required
		return nil
	}

	for i := 0; i < len(m.Feature); i++ {
		if swag.IsZero(m.Feature[i]) { // not required
			continue
		}

		if m.Feature[i] != nil {
			if err := m.Feature[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("feature" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("feature" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Service) validateNote(formats strfmt.Registry) error {
	if swag.IsZero(m.Note) { // not required
		return nil
	}

	for i := 0; i < len(m.Note); i++ {
		if swag.IsZero(m.Note[i]) { // not required
			continue
		}

		if m.Note[i] != nil {
			if err := m.Note[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("note" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("note" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Service) validateOperatingStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.OperatingStatus) { // not required
		return nil
	}

	if err := m.OperatingStatus.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("operatingStatus")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("operatingStatus")
		}
		return err
	}

	return nil
}

func (m *Service) validateOperatingStatusContextUpdate(formats strfmt.Registry) error {
	if swag.IsZero(m.OperatingStatusContextUpdate) { // not required
		return nil
	}

	if m.OperatingStatusContextUpdate != nil {
		if err := m.OperatingStatusContextUpdate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operatingStatusContextUpdate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operatingStatusContextUpdate")
			}
			return err
		}
	}

	return nil
}

func (m *Service) validatePlace(formats strfmt.Registry) error {
	if swag.IsZero(m.Place) { // not required
		return nil
	}

	for i := 0; i < len(m.Place); i++ {
		if swag.IsZero(m.Place[i]) { // not required
			continue
		}

		if m.Place[i] != nil {
			if err := m.Place[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("place" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("place" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Service) validateRelatedEntity(formats strfmt.Registry) error {
	if swag.IsZero(m.RelatedEntity) { // not required
		return nil
	}

	for i := 0; i < len(m.RelatedEntity); i++ {
		if swag.IsZero(m.RelatedEntity[i]) { // not required
			continue
		}

		if m.RelatedEntity[i] != nil {
			if err := m.RelatedEntity[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relatedEntity" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("relatedEntity" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Service) validateRelatedParty(formats strfmt.Registry) error {
	if swag.IsZero(m.RelatedParty) { // not required
		return nil
	}

	for i := 0; i < len(m.RelatedParty); i++ {
		if swag.IsZero(m.RelatedParty[i]) { // not required
			continue
		}

		if m.RelatedParty[i] != nil {
			if err := m.RelatedParty[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relatedParty" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("relatedParty" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Service) validateServiceCharacteristic(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceCharacteristic) { // not required
		return nil
	}

	for i := 0; i < len(m.ServiceCharacteristic); i++ {
		if swag.IsZero(m.ServiceCharacteristic[i]) { // not required
			continue
		}

		if m.ServiceCharacteristic[i] != nil {
			if err := m.ServiceCharacteristic[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serviceCharacteristic" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("serviceCharacteristic" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Service) validateServiceOrderItem(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceOrderItem) { // not required
		return nil
	}

	for i := 0; i < len(m.ServiceOrderItem); i++ {
		if swag.IsZero(m.ServiceOrderItem[i]) { // not required
			continue
		}

		if m.ServiceOrderItem[i] != nil {
			if err := m.ServiceOrderItem[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serviceOrderItem" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("serviceOrderItem" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Service) validateServiceRelationship(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceRelationship) { // not required
		return nil
	}

	for i := 0; i < len(m.ServiceRelationship); i++ {
		if swag.IsZero(m.ServiceRelationship[i]) { // not required
			continue
		}

		if m.ServiceRelationship[i] != nil {
			if err := m.ServiceRelationship[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serviceRelationship" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("serviceRelationship" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Service) validateServiceSpecification(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceSpecification) { // not required
		return nil
	}

	if m.ServiceSpecification != nil {
		if err := m.ServiceSpecification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serviceSpecification")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("serviceSpecification")
			}
			return err
		}
	}

	return nil
}

func (m *Service) validateStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Service) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("state")
		}
		return err
	}

	return nil
}

func (m *Service) validateSupportingResource(formats strfmt.Registry) error {
	if swag.IsZero(m.SupportingResource) { // not required
		return nil
	}

	for i := 0; i < len(m.SupportingResource); i++ {
		if swag.IsZero(m.SupportingResource[i]) { // not required
			continue
		}

		if m.SupportingResource[i] != nil {
			if err := m.SupportingResource[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("supportingResource" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("supportingResource" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Service) validateSupportingService(formats strfmt.Registry) error {
	if swag.IsZero(m.SupportingService) { // not required
		return nil
	}

	for i := 0; i < len(m.SupportingService); i++ {
		if swag.IsZero(m.SupportingService[i]) { // not required
			continue
		}

		if m.SupportingService[i] != nil {
			if err := m.SupportingService[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("supportingService" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("supportingService" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this service based on the context it is used
func (m *Service) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExternalIdentifier(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFeature(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNote(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOperatingStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOperatingStatusContextUpdate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePlace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRelatedEntity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRelatedParty(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServiceCharacteristic(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServiceOrderItem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServiceRelationship(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServiceSpecification(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSupportingResource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSupportingService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Service) contextValidateExternalIdentifier(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExternalIdentifier); i++ {

		if m.ExternalIdentifier[i] != nil {

			if swag.IsZero(m.ExternalIdentifier[i]) { // not required
				return nil
			}

			if err := m.ExternalIdentifier[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("externalIdentifier" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("externalIdentifier" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Service) contextValidateFeature(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Feature); i++ {

		if m.Feature[i] != nil {

			if swag.IsZero(m.Feature[i]) { // not required
				return nil
			}

			if err := m.Feature[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("feature" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("feature" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Service) contextValidateNote(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Note); i++ {

		if m.Note[i] != nil {

			if swag.IsZero(m.Note[i]) { // not required
				return nil
			}

			if err := m.Note[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("note" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("note" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Service) contextValidateOperatingStatus(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.OperatingStatus) { // not required
		return nil
	}

	if err := m.OperatingStatus.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("operatingStatus")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("operatingStatus")
		}
		return err
	}

	return nil
}

func (m *Service) contextValidateOperatingStatusContextUpdate(ctx context.Context, formats strfmt.Registry) error {

	if m.OperatingStatusContextUpdate != nil {

		if swag.IsZero(m.OperatingStatusContextUpdate) { // not required
			return nil
		}

		if err := m.OperatingStatusContextUpdate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("operatingStatusContextUpdate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("operatingStatusContextUpdate")
			}
			return err
		}
	}

	return nil
}

func (m *Service) contextValidatePlace(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Place); i++ {

		if m.Place[i] != nil {

			if swag.IsZero(m.Place[i]) { // not required
				return nil
			}

			if err := m.Place[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("place" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("place" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Service) contextValidateRelatedEntity(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RelatedEntity); i++ {

		if m.RelatedEntity[i] != nil {

			if swag.IsZero(m.RelatedEntity[i]) { // not required
				return nil
			}

			if err := m.RelatedEntity[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relatedEntity" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("relatedEntity" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Service) contextValidateRelatedParty(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RelatedParty); i++ {

		if m.RelatedParty[i] != nil {

			if swag.IsZero(m.RelatedParty[i]) { // not required
				return nil
			}

			if err := m.RelatedParty[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relatedParty" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("relatedParty" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Service) contextValidateServiceCharacteristic(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ServiceCharacteristic); i++ {

		if m.ServiceCharacteristic[i] != nil {

			if swag.IsZero(m.ServiceCharacteristic[i]) { // not required
				return nil
			}

			if err := m.ServiceCharacteristic[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serviceCharacteristic" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("serviceCharacteristic" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Service) contextValidateServiceOrderItem(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ServiceOrderItem); i++ {

		if m.ServiceOrderItem[i] != nil {

			if swag.IsZero(m.ServiceOrderItem[i]) { // not required
				return nil
			}

			if err := m.ServiceOrderItem[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serviceOrderItem" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("serviceOrderItem" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Service) contextValidateServiceRelationship(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ServiceRelationship); i++ {

		if m.ServiceRelationship[i] != nil {

			if swag.IsZero(m.ServiceRelationship[i]) { // not required
				return nil
			}

			if err := m.ServiceRelationship[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serviceRelationship" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("serviceRelationship" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Service) contextValidateServiceSpecification(ctx context.Context, formats strfmt.Registry) error {

	if m.ServiceSpecification != nil {

		if swag.IsZero(m.ServiceSpecification) { // not required
			return nil
		}

		if err := m.ServiceSpecification.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serviceSpecification")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("serviceSpecification")
			}
			return err
		}
	}

	return nil
}

func (m *Service) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("state")
		}
		return err
	}

	return nil
}

func (m *Service) contextValidateSupportingResource(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SupportingResource); i++ {

		if m.SupportingResource[i] != nil {

			if swag.IsZero(m.SupportingResource[i]) { // not required
				return nil
			}

			if err := m.SupportingResource[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("supportingResource" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("supportingResource" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Service) contextValidateSupportingService(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SupportingService); i++ {

		if m.SupportingService[i] != nil {

			if swag.IsZero(m.SupportingService[i]) { // not required
				return nil
			}

			if err := m.SupportingService[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("supportingService" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("supportingService" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Service) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Service) UnmarshalBinary(b []byte) error {
	var res Service
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
