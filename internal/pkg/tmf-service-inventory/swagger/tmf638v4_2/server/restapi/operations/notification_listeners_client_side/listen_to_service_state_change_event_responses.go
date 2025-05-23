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
// ListenToServiceStateChangeEventCreatedCode is the HTTP code returned for type ListenToServiceStateChangeEventCreated
const ListenToServiceStateChangeEventCreatedCode int = 201

/*
ListenToServiceStateChangeEventCreated Notified

swagger:response listenToServiceStateChangeEventCreated
*/
type ListenToServiceStateChangeEventCreated struct {

	/*
	  In: Body
	*/
	Payload *models.EventSubscription `json:"body,omitempty"`
}

type ListenToServiceStateChangeEventCreatedRaw struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewListenToServiceStateChangeEventCreated creates ListenToServiceStateChangeEventCreated with default headers values
func NewListenToServiceStateChangeEventCreated() *ListenToServiceStateChangeEventCreated {

	return &ListenToServiceStateChangeEventCreated{}
}

// NewListenToServiceStateChangeEventCreated creates ListenToServiceStateChangeEventCreatedRaw with default headers values
func NewListenToServiceStateChangeEventCreatedRaw() *ListenToServiceStateChangeEventCreatedRaw {

	return &ListenToServiceStateChangeEventCreatedRaw{}
}

// WithPayload adds the payload to the listen to service state change event created response
func (o *ListenToServiceStateChangeEventCreated) WithPayload(payload *models.EventSubscription) *ListenToServiceStateChangeEventCreated {
	o.Payload = payload
	return o
}

// WithPayload adds the payload to the listen to service state change event created response
func (o *ListenToServiceStateChangeEventCreatedRaw) WithPayload(payload interface{}) *ListenToServiceStateChangeEventCreatedRaw {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service state change event created response
func (o *ListenToServiceStateChangeEventCreated) SetPayload(payload *models.EventSubscription) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceStateChangeEventCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WriteResponseRaw to the client
func (o *ListenToServiceStateChangeEventCreatedRaw) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// Alex Server response
// ListenToServiceStateChangeEventBadRequestCode is the HTTP code returned for type ListenToServiceStateChangeEventBadRequest
const ListenToServiceStateChangeEventBadRequestCode int = 400

/*
ListenToServiceStateChangeEventBadRequest Bad Request

swagger:response listenToServiceStateChangeEventBadRequest
*/
type ListenToServiceStateChangeEventBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

type ListenToServiceStateChangeEventBadRequestRaw struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewListenToServiceStateChangeEventBadRequest creates ListenToServiceStateChangeEventBadRequest with default headers values
func NewListenToServiceStateChangeEventBadRequest() *ListenToServiceStateChangeEventBadRequest {

	return &ListenToServiceStateChangeEventBadRequest{}
}

// NewListenToServiceStateChangeEventBadRequest creates ListenToServiceStateChangeEventBadRequestRaw with default headers values
func NewListenToServiceStateChangeEventBadRequestRaw() *ListenToServiceStateChangeEventBadRequestRaw {

	return &ListenToServiceStateChangeEventBadRequestRaw{}
}

// WithPayload adds the payload to the listen to service state change event bad request response
func (o *ListenToServiceStateChangeEventBadRequest) WithPayload(payload *models.Error) *ListenToServiceStateChangeEventBadRequest {
	o.Payload = payload
	return o
}

// WithPayload adds the payload to the listen to service state change event bad request response
func (o *ListenToServiceStateChangeEventBadRequestRaw) WithPayload(payload interface{}) *ListenToServiceStateChangeEventBadRequestRaw {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service state change event bad request response
func (o *ListenToServiceStateChangeEventBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceStateChangeEventBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WriteResponseRaw to the client
func (o *ListenToServiceStateChangeEventBadRequestRaw) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// Alex Server response
// ListenToServiceStateChangeEventUnauthorizedCode is the HTTP code returned for type ListenToServiceStateChangeEventUnauthorized
const ListenToServiceStateChangeEventUnauthorizedCode int = 401

/*
ListenToServiceStateChangeEventUnauthorized Unauthorized

swagger:response listenToServiceStateChangeEventUnauthorized
*/
type ListenToServiceStateChangeEventUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

type ListenToServiceStateChangeEventUnauthorizedRaw struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewListenToServiceStateChangeEventUnauthorized creates ListenToServiceStateChangeEventUnauthorized with default headers values
func NewListenToServiceStateChangeEventUnauthorized() *ListenToServiceStateChangeEventUnauthorized {

	return &ListenToServiceStateChangeEventUnauthorized{}
}

// NewListenToServiceStateChangeEventUnauthorized creates ListenToServiceStateChangeEventUnauthorizedRaw with default headers values
func NewListenToServiceStateChangeEventUnauthorizedRaw() *ListenToServiceStateChangeEventUnauthorizedRaw {

	return &ListenToServiceStateChangeEventUnauthorizedRaw{}
}

// WithPayload adds the payload to the listen to service state change event unauthorized response
func (o *ListenToServiceStateChangeEventUnauthorized) WithPayload(payload *models.Error) *ListenToServiceStateChangeEventUnauthorized {
	o.Payload = payload
	return o
}

// WithPayload adds the payload to the listen to service state change event unauthorized response
func (o *ListenToServiceStateChangeEventUnauthorizedRaw) WithPayload(payload interface{}) *ListenToServiceStateChangeEventUnauthorizedRaw {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service state change event unauthorized response
func (o *ListenToServiceStateChangeEventUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceStateChangeEventUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WriteResponseRaw to the client
func (o *ListenToServiceStateChangeEventUnauthorizedRaw) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// Alex Server response
// ListenToServiceStateChangeEventForbiddenCode is the HTTP code returned for type ListenToServiceStateChangeEventForbidden
const ListenToServiceStateChangeEventForbiddenCode int = 403

/*
ListenToServiceStateChangeEventForbidden Forbidden

swagger:response listenToServiceStateChangeEventForbidden
*/
type ListenToServiceStateChangeEventForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

type ListenToServiceStateChangeEventForbiddenRaw struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewListenToServiceStateChangeEventForbidden creates ListenToServiceStateChangeEventForbidden with default headers values
func NewListenToServiceStateChangeEventForbidden() *ListenToServiceStateChangeEventForbidden {

	return &ListenToServiceStateChangeEventForbidden{}
}

// NewListenToServiceStateChangeEventForbidden creates ListenToServiceStateChangeEventForbiddenRaw with default headers values
func NewListenToServiceStateChangeEventForbiddenRaw() *ListenToServiceStateChangeEventForbiddenRaw {

	return &ListenToServiceStateChangeEventForbiddenRaw{}
}

// WithPayload adds the payload to the listen to service state change event forbidden response
func (o *ListenToServiceStateChangeEventForbidden) WithPayload(payload *models.Error) *ListenToServiceStateChangeEventForbidden {
	o.Payload = payload
	return o
}

// WithPayload adds the payload to the listen to service state change event forbidden response
func (o *ListenToServiceStateChangeEventForbiddenRaw) WithPayload(payload interface{}) *ListenToServiceStateChangeEventForbiddenRaw {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service state change event forbidden response
func (o *ListenToServiceStateChangeEventForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceStateChangeEventForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WriteResponseRaw to the client
func (o *ListenToServiceStateChangeEventForbiddenRaw) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// Alex Server response
// ListenToServiceStateChangeEventNotFoundCode is the HTTP code returned for type ListenToServiceStateChangeEventNotFound
const ListenToServiceStateChangeEventNotFoundCode int = 404

/*
ListenToServiceStateChangeEventNotFound Not Found

swagger:response listenToServiceStateChangeEventNotFound
*/
type ListenToServiceStateChangeEventNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

type ListenToServiceStateChangeEventNotFoundRaw struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewListenToServiceStateChangeEventNotFound creates ListenToServiceStateChangeEventNotFound with default headers values
func NewListenToServiceStateChangeEventNotFound() *ListenToServiceStateChangeEventNotFound {

	return &ListenToServiceStateChangeEventNotFound{}
}

// NewListenToServiceStateChangeEventNotFound creates ListenToServiceStateChangeEventNotFoundRaw with default headers values
func NewListenToServiceStateChangeEventNotFoundRaw() *ListenToServiceStateChangeEventNotFoundRaw {

	return &ListenToServiceStateChangeEventNotFoundRaw{}
}

// WithPayload adds the payload to the listen to service state change event not found response
func (o *ListenToServiceStateChangeEventNotFound) WithPayload(payload *models.Error) *ListenToServiceStateChangeEventNotFound {
	o.Payload = payload
	return o
}

// WithPayload adds the payload to the listen to service state change event not found response
func (o *ListenToServiceStateChangeEventNotFoundRaw) WithPayload(payload interface{}) *ListenToServiceStateChangeEventNotFoundRaw {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service state change event not found response
func (o *ListenToServiceStateChangeEventNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceStateChangeEventNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WriteResponseRaw to the client
func (o *ListenToServiceStateChangeEventNotFoundRaw) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// Alex Server response
// ListenToServiceStateChangeEventMethodNotAllowedCode is the HTTP code returned for type ListenToServiceStateChangeEventMethodNotAllowed
const ListenToServiceStateChangeEventMethodNotAllowedCode int = 405

/*
ListenToServiceStateChangeEventMethodNotAllowed Method Not allowed

swagger:response listenToServiceStateChangeEventMethodNotAllowed
*/
type ListenToServiceStateChangeEventMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

type ListenToServiceStateChangeEventMethodNotAllowedRaw struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewListenToServiceStateChangeEventMethodNotAllowed creates ListenToServiceStateChangeEventMethodNotAllowed with default headers values
func NewListenToServiceStateChangeEventMethodNotAllowed() *ListenToServiceStateChangeEventMethodNotAllowed {

	return &ListenToServiceStateChangeEventMethodNotAllowed{}
}

// NewListenToServiceStateChangeEventMethodNotAllowed creates ListenToServiceStateChangeEventMethodNotAllowedRaw with default headers values
func NewListenToServiceStateChangeEventMethodNotAllowedRaw() *ListenToServiceStateChangeEventMethodNotAllowedRaw {

	return &ListenToServiceStateChangeEventMethodNotAllowedRaw{}
}

// WithPayload adds the payload to the listen to service state change event method not allowed response
func (o *ListenToServiceStateChangeEventMethodNotAllowed) WithPayload(payload *models.Error) *ListenToServiceStateChangeEventMethodNotAllowed {
	o.Payload = payload
	return o
}

// WithPayload adds the payload to the listen to service state change event method not allowed response
func (o *ListenToServiceStateChangeEventMethodNotAllowedRaw) WithPayload(payload interface{}) *ListenToServiceStateChangeEventMethodNotAllowedRaw {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service state change event method not allowed response
func (o *ListenToServiceStateChangeEventMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceStateChangeEventMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WriteResponseRaw to the client
func (o *ListenToServiceStateChangeEventMethodNotAllowedRaw) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// Alex Server response
// ListenToServiceStateChangeEventConflictCode is the HTTP code returned for type ListenToServiceStateChangeEventConflict
const ListenToServiceStateChangeEventConflictCode int = 409

/*
ListenToServiceStateChangeEventConflict Conflict

swagger:response listenToServiceStateChangeEventConflict
*/
type ListenToServiceStateChangeEventConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

type ListenToServiceStateChangeEventConflictRaw struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewListenToServiceStateChangeEventConflict creates ListenToServiceStateChangeEventConflict with default headers values
func NewListenToServiceStateChangeEventConflict() *ListenToServiceStateChangeEventConflict {

	return &ListenToServiceStateChangeEventConflict{}
}

// NewListenToServiceStateChangeEventConflict creates ListenToServiceStateChangeEventConflictRaw with default headers values
func NewListenToServiceStateChangeEventConflictRaw() *ListenToServiceStateChangeEventConflictRaw {

	return &ListenToServiceStateChangeEventConflictRaw{}
}

// WithPayload adds the payload to the listen to service state change event conflict response
func (o *ListenToServiceStateChangeEventConflict) WithPayload(payload *models.Error) *ListenToServiceStateChangeEventConflict {
	o.Payload = payload
	return o
}

// WithPayload adds the payload to the listen to service state change event conflict response
func (o *ListenToServiceStateChangeEventConflictRaw) WithPayload(payload interface{}) *ListenToServiceStateChangeEventConflictRaw {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service state change event conflict response
func (o *ListenToServiceStateChangeEventConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceStateChangeEventConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WriteResponseRaw to the client
func (o *ListenToServiceStateChangeEventConflictRaw) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// Alex Server response
// ListenToServiceStateChangeEventInternalServerErrorCode is the HTTP code returned for type ListenToServiceStateChangeEventInternalServerError
const ListenToServiceStateChangeEventInternalServerErrorCode int = 500

/*
ListenToServiceStateChangeEventInternalServerError Internal Server Error

swagger:response listenToServiceStateChangeEventInternalServerError
*/
type ListenToServiceStateChangeEventInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

type ListenToServiceStateChangeEventInternalServerErrorRaw struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewListenToServiceStateChangeEventInternalServerError creates ListenToServiceStateChangeEventInternalServerError with default headers values
func NewListenToServiceStateChangeEventInternalServerError() *ListenToServiceStateChangeEventInternalServerError {

	return &ListenToServiceStateChangeEventInternalServerError{}
}

// NewListenToServiceStateChangeEventInternalServerError creates ListenToServiceStateChangeEventInternalServerErrorRaw with default headers values
func NewListenToServiceStateChangeEventInternalServerErrorRaw() *ListenToServiceStateChangeEventInternalServerErrorRaw {

	return &ListenToServiceStateChangeEventInternalServerErrorRaw{}
}

// WithPayload adds the payload to the listen to service state change event internal server error response
func (o *ListenToServiceStateChangeEventInternalServerError) WithPayload(payload *models.Error) *ListenToServiceStateChangeEventInternalServerError {
	o.Payload = payload
	return o
}

// WithPayload adds the payload to the listen to service state change event internal server error response
func (o *ListenToServiceStateChangeEventInternalServerErrorRaw) WithPayload(payload interface{}) *ListenToServiceStateChangeEventInternalServerErrorRaw {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the listen to service state change event internal server error response
func (o *ListenToServiceStateChangeEventInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListenToServiceStateChangeEventInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WriteResponseRaw to the client
func (o *ListenToServiceStateChangeEventInternalServerErrorRaw) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
