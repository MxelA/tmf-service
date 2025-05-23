// Code generated by go-swagger; DO NOT EDIT.

package notification_listeners_client_side

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/MxelA/tmf-service/internal/pkg/tmf-service-inventory/swagger/tmf638v4_2/server/models"
)

// Alex Server response
// ListenToServiceAttributeValueChangeEventCreatedCode is the HTTP code returned for type ListenToServiceAttributeValueChangeEventCreated
const ListenToServiceAttributeValueChangeEventCreatedCode int = 201

/*
ListenToServiceAttributeValueChangeEventCreated Notified

swagger:response listenToServiceAttributeValueChangeEventCreated
*/
type ListenToServiceAttributeValueChangeEventCreated struct {

	/*
	  In: Body
	*/
	Payload *models.EventSubscription `json:"body,omitempty"`
}

type ListenToServiceAttributeValueChangeEventCreatedRaw struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewListenToServiceAttributeValueChangeEventCreated creates ListenToServiceAttributeValueChangeEventCreated with default headers values
func NewListenToServiceAttributeValueChangeEventCreated() *ListenToServiceAttributeValueChangeEventCreated {

	return &ListenToServiceAttributeValueChangeEventCreated{}
}

// NewListenToServiceAttributeValueChangeEventCreated creates ListenToServiceAttributeValueChangeEventCreatedRaw with default headers values
func NewListenToServiceAttributeValueChangeEventCreatedRaw() *ListenToServiceAttributeValueChangeEventCreatedRaw {

	return &ListenToServiceAttributeValueChangeEventCreatedRaw{}
}

// WithPayload adds the payload to the listen to service attribute value change event created response
func (o *ListenToServiceAttributeValueChangeEventCreated) WithPayload(payload *models.EventSubscription) *ListenToServiceAttributeValueChangeEventCreated {
	o.Payload = payload
	return o
}

// WithPayload adds the payload to the listen to service attribute value change event created response
func (o *ListenToServiceAttributeValueChangeEventCreatedRaw) WithPayload(payload interface{}) *ListenToServiceAttributeValueChangeEventCreatedRaw {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service attribute value change event created response
func (o *ListenToServiceAttributeValueChangeEventCreated) SetPayload(payload *models.EventSubscription) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceAttributeValueChangeEventCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WriteResponseRaw to the client
func (o *ListenToServiceAttributeValueChangeEventCreatedRaw) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// Alex Server response
// ListenToServiceAttributeValueChangeEventBadRequestCode is the HTTP code returned for type ListenToServiceAttributeValueChangeEventBadRequest
const ListenToServiceAttributeValueChangeEventBadRequestCode int = 400

/*
ListenToServiceAttributeValueChangeEventBadRequest Bad Request

swagger:response listenToServiceAttributeValueChangeEventBadRequest
*/
type ListenToServiceAttributeValueChangeEventBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

type ListenToServiceAttributeValueChangeEventBadRequestRaw struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewListenToServiceAttributeValueChangeEventBadRequest creates ListenToServiceAttributeValueChangeEventBadRequest with default headers values
func NewListenToServiceAttributeValueChangeEventBadRequest() *ListenToServiceAttributeValueChangeEventBadRequest {

	return &ListenToServiceAttributeValueChangeEventBadRequest{}
}

// NewListenToServiceAttributeValueChangeEventBadRequest creates ListenToServiceAttributeValueChangeEventBadRequestRaw with default headers values
func NewListenToServiceAttributeValueChangeEventBadRequestRaw() *ListenToServiceAttributeValueChangeEventBadRequestRaw {

	return &ListenToServiceAttributeValueChangeEventBadRequestRaw{}
}

// WithPayload adds the payload to the listen to service attribute value change event bad request response
func (o *ListenToServiceAttributeValueChangeEventBadRequest) WithPayload(payload *models.Error) *ListenToServiceAttributeValueChangeEventBadRequest {
	o.Payload = payload
	return o
}

// WithPayload adds the payload to the listen to service attribute value change event bad request response
func (o *ListenToServiceAttributeValueChangeEventBadRequestRaw) WithPayload(payload interface{}) *ListenToServiceAttributeValueChangeEventBadRequestRaw {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service attribute value change event bad request response
func (o *ListenToServiceAttributeValueChangeEventBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceAttributeValueChangeEventBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WriteResponseRaw to the client
func (o *ListenToServiceAttributeValueChangeEventBadRequestRaw) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// Alex Server response
// ListenToServiceAttributeValueChangeEventUnauthorizedCode is the HTTP code returned for type ListenToServiceAttributeValueChangeEventUnauthorized
const ListenToServiceAttributeValueChangeEventUnauthorizedCode int = 401

/*
ListenToServiceAttributeValueChangeEventUnauthorized Unauthorized

swagger:response listenToServiceAttributeValueChangeEventUnauthorized
*/
type ListenToServiceAttributeValueChangeEventUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

type ListenToServiceAttributeValueChangeEventUnauthorizedRaw struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewListenToServiceAttributeValueChangeEventUnauthorized creates ListenToServiceAttributeValueChangeEventUnauthorized with default headers values
func NewListenToServiceAttributeValueChangeEventUnauthorized() *ListenToServiceAttributeValueChangeEventUnauthorized {

	return &ListenToServiceAttributeValueChangeEventUnauthorized{}
}

// NewListenToServiceAttributeValueChangeEventUnauthorized creates ListenToServiceAttributeValueChangeEventUnauthorizedRaw with default headers values
func NewListenToServiceAttributeValueChangeEventUnauthorizedRaw() *ListenToServiceAttributeValueChangeEventUnauthorizedRaw {

	return &ListenToServiceAttributeValueChangeEventUnauthorizedRaw{}
}

// WithPayload adds the payload to the listen to service attribute value change event unauthorized response
func (o *ListenToServiceAttributeValueChangeEventUnauthorized) WithPayload(payload *models.Error) *ListenToServiceAttributeValueChangeEventUnauthorized {
	o.Payload = payload
	return o
}

// WithPayload adds the payload to the listen to service attribute value change event unauthorized response
func (o *ListenToServiceAttributeValueChangeEventUnauthorizedRaw) WithPayload(payload interface{}) *ListenToServiceAttributeValueChangeEventUnauthorizedRaw {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service attribute value change event unauthorized response
func (o *ListenToServiceAttributeValueChangeEventUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceAttributeValueChangeEventUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WriteResponseRaw to the client
func (o *ListenToServiceAttributeValueChangeEventUnauthorizedRaw) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// Alex Server response
// ListenToServiceAttributeValueChangeEventForbiddenCode is the HTTP code returned for type ListenToServiceAttributeValueChangeEventForbidden
const ListenToServiceAttributeValueChangeEventForbiddenCode int = 403

/*
ListenToServiceAttributeValueChangeEventForbidden Forbidden

swagger:response listenToServiceAttributeValueChangeEventForbidden
*/
type ListenToServiceAttributeValueChangeEventForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

type ListenToServiceAttributeValueChangeEventForbiddenRaw struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewListenToServiceAttributeValueChangeEventForbidden creates ListenToServiceAttributeValueChangeEventForbidden with default headers values
func NewListenToServiceAttributeValueChangeEventForbidden() *ListenToServiceAttributeValueChangeEventForbidden {

	return &ListenToServiceAttributeValueChangeEventForbidden{}
}

// NewListenToServiceAttributeValueChangeEventForbidden creates ListenToServiceAttributeValueChangeEventForbiddenRaw with default headers values
func NewListenToServiceAttributeValueChangeEventForbiddenRaw() *ListenToServiceAttributeValueChangeEventForbiddenRaw {

	return &ListenToServiceAttributeValueChangeEventForbiddenRaw{}
}

// WithPayload adds the payload to the listen to service attribute value change event forbidden response
func (o *ListenToServiceAttributeValueChangeEventForbidden) WithPayload(payload *models.Error) *ListenToServiceAttributeValueChangeEventForbidden {
	o.Payload = payload
	return o
}

// WithPayload adds the payload to the listen to service attribute value change event forbidden response
func (o *ListenToServiceAttributeValueChangeEventForbiddenRaw) WithPayload(payload interface{}) *ListenToServiceAttributeValueChangeEventForbiddenRaw {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service attribute value change event forbidden response
func (o *ListenToServiceAttributeValueChangeEventForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceAttributeValueChangeEventForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WriteResponseRaw to the client
func (o *ListenToServiceAttributeValueChangeEventForbiddenRaw) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// Alex Server response
// ListenToServiceAttributeValueChangeEventNotFoundCode is the HTTP code returned for type ListenToServiceAttributeValueChangeEventNotFound
const ListenToServiceAttributeValueChangeEventNotFoundCode int = 404

/*
ListenToServiceAttributeValueChangeEventNotFound Not Found

swagger:response listenToServiceAttributeValueChangeEventNotFound
*/
type ListenToServiceAttributeValueChangeEventNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

type ListenToServiceAttributeValueChangeEventNotFoundRaw struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewListenToServiceAttributeValueChangeEventNotFound creates ListenToServiceAttributeValueChangeEventNotFound with default headers values
func NewListenToServiceAttributeValueChangeEventNotFound() *ListenToServiceAttributeValueChangeEventNotFound {

	return &ListenToServiceAttributeValueChangeEventNotFound{}
}

// NewListenToServiceAttributeValueChangeEventNotFound creates ListenToServiceAttributeValueChangeEventNotFoundRaw with default headers values
func NewListenToServiceAttributeValueChangeEventNotFoundRaw() *ListenToServiceAttributeValueChangeEventNotFoundRaw {

	return &ListenToServiceAttributeValueChangeEventNotFoundRaw{}
}

// WithPayload adds the payload to the listen to service attribute value change event not found response
func (o *ListenToServiceAttributeValueChangeEventNotFound) WithPayload(payload *models.Error) *ListenToServiceAttributeValueChangeEventNotFound {
	o.Payload = payload
	return o
}

// WithPayload adds the payload to the listen to service attribute value change event not found response
func (o *ListenToServiceAttributeValueChangeEventNotFoundRaw) WithPayload(payload interface{}) *ListenToServiceAttributeValueChangeEventNotFoundRaw {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service attribute value change event not found response
func (o *ListenToServiceAttributeValueChangeEventNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceAttributeValueChangeEventNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WriteResponseRaw to the client
func (o *ListenToServiceAttributeValueChangeEventNotFoundRaw) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// Alex Server response
// ListenToServiceAttributeValueChangeEventMethodNotAllowedCode is the HTTP code returned for type ListenToServiceAttributeValueChangeEventMethodNotAllowed
const ListenToServiceAttributeValueChangeEventMethodNotAllowedCode int = 405

/*
ListenToServiceAttributeValueChangeEventMethodNotAllowed Method Not allowed

swagger:response listenToServiceAttributeValueChangeEventMethodNotAllowed
*/
type ListenToServiceAttributeValueChangeEventMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

type ListenToServiceAttributeValueChangeEventMethodNotAllowedRaw struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewListenToServiceAttributeValueChangeEventMethodNotAllowed creates ListenToServiceAttributeValueChangeEventMethodNotAllowed with default headers values
func NewListenToServiceAttributeValueChangeEventMethodNotAllowed() *ListenToServiceAttributeValueChangeEventMethodNotAllowed {

	return &ListenToServiceAttributeValueChangeEventMethodNotAllowed{}
}

// NewListenToServiceAttributeValueChangeEventMethodNotAllowed creates ListenToServiceAttributeValueChangeEventMethodNotAllowedRaw with default headers values
func NewListenToServiceAttributeValueChangeEventMethodNotAllowedRaw() *ListenToServiceAttributeValueChangeEventMethodNotAllowedRaw {

	return &ListenToServiceAttributeValueChangeEventMethodNotAllowedRaw{}
}

// WithPayload adds the payload to the listen to service attribute value change event method not allowed response
func (o *ListenToServiceAttributeValueChangeEventMethodNotAllowed) WithPayload(payload *models.Error) *ListenToServiceAttributeValueChangeEventMethodNotAllowed {
	o.Payload = payload
	return o
}

// WithPayload adds the payload to the listen to service attribute value change event method not allowed response
func (o *ListenToServiceAttributeValueChangeEventMethodNotAllowedRaw) WithPayload(payload interface{}) *ListenToServiceAttributeValueChangeEventMethodNotAllowedRaw {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service attribute value change event method not allowed response
func (o *ListenToServiceAttributeValueChangeEventMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceAttributeValueChangeEventMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WriteResponseRaw to the client
func (o *ListenToServiceAttributeValueChangeEventMethodNotAllowedRaw) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// Alex Server response
// ListenToServiceAttributeValueChangeEventConflictCode is the HTTP code returned for type ListenToServiceAttributeValueChangeEventConflict
const ListenToServiceAttributeValueChangeEventConflictCode int = 409

/*
ListenToServiceAttributeValueChangeEventConflict Conflict

swagger:response listenToServiceAttributeValueChangeEventConflict
*/
type ListenToServiceAttributeValueChangeEventConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

type ListenToServiceAttributeValueChangeEventConflictRaw struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewListenToServiceAttributeValueChangeEventConflict creates ListenToServiceAttributeValueChangeEventConflict with default headers values
func NewListenToServiceAttributeValueChangeEventConflict() *ListenToServiceAttributeValueChangeEventConflict {

	return &ListenToServiceAttributeValueChangeEventConflict{}
}

// NewListenToServiceAttributeValueChangeEventConflict creates ListenToServiceAttributeValueChangeEventConflictRaw with default headers values
func NewListenToServiceAttributeValueChangeEventConflictRaw() *ListenToServiceAttributeValueChangeEventConflictRaw {

	return &ListenToServiceAttributeValueChangeEventConflictRaw{}
}

// WithPayload adds the payload to the listen to service attribute value change event conflict response
func (o *ListenToServiceAttributeValueChangeEventConflict) WithPayload(payload *models.Error) *ListenToServiceAttributeValueChangeEventConflict {
	o.Payload = payload
	return o
}

// WithPayload adds the payload to the listen to service attribute value change event conflict response
func (o *ListenToServiceAttributeValueChangeEventConflictRaw) WithPayload(payload interface{}) *ListenToServiceAttributeValueChangeEventConflictRaw {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service attribute value change event conflict response
func (o *ListenToServiceAttributeValueChangeEventConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceAttributeValueChangeEventConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WriteResponseRaw to the client
func (o *ListenToServiceAttributeValueChangeEventConflictRaw) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// Alex Server response
// ListenToServiceAttributeValueChangeEventInternalServerErrorCode is the HTTP code returned for type ListenToServiceAttributeValueChangeEventInternalServerError
const ListenToServiceAttributeValueChangeEventInternalServerErrorCode int = 500

/*
ListenToServiceAttributeValueChangeEventInternalServerError Internal Server Error

swagger:response listenToServiceAttributeValueChangeEventInternalServerError
*/
type ListenToServiceAttributeValueChangeEventInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

type ListenToServiceAttributeValueChangeEventInternalServerErrorRaw struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewListenToServiceAttributeValueChangeEventInternalServerError creates ListenToServiceAttributeValueChangeEventInternalServerError with default headers values
func NewListenToServiceAttributeValueChangeEventInternalServerError() *ListenToServiceAttributeValueChangeEventInternalServerError {

	return &ListenToServiceAttributeValueChangeEventInternalServerError{}
}

// NewListenToServiceAttributeValueChangeEventInternalServerError creates ListenToServiceAttributeValueChangeEventInternalServerErrorRaw with default headers values
func NewListenToServiceAttributeValueChangeEventInternalServerErrorRaw() *ListenToServiceAttributeValueChangeEventInternalServerErrorRaw {

	return &ListenToServiceAttributeValueChangeEventInternalServerErrorRaw{}
}

// WithPayload adds the payload to the listen to service attribute value change event internal server error response
func (o *ListenToServiceAttributeValueChangeEventInternalServerError) WithPayload(payload *models.Error) *ListenToServiceAttributeValueChangeEventInternalServerError {
	o.Payload = payload
	return o
}

// WithPayload adds the payload to the listen to service attribute value change event internal server error response
func (o *ListenToServiceAttributeValueChangeEventInternalServerErrorRaw) WithPayload(payload interface{}) *ListenToServiceAttributeValueChangeEventInternalServerErrorRaw {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service attribute value change event internal server error response
func (o *ListenToServiceAttributeValueChangeEventInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceAttributeValueChangeEventInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WriteResponseRaw to the client
func (o *ListenToServiceAttributeValueChangeEventInternalServerErrorRaw) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
