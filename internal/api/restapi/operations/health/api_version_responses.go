// Code generated by go-swagger; DO NOT EDIT.

package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kaatinga/plantbook/internal/api/models"
)

// APIVersionOKCode is the HTTP code returned for type APIVersionOK
const APIVersionOKCode int = 200

/*APIVersionOK version info

swagger:response apiVersionOK
*/
type APIVersionOK struct {
	/*The request id this is a response to

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.APIVersion `json:"body,omitempty"`
}

// NewAPIVersionOK creates APIVersionOK with default headers values
func NewAPIVersionOK() *APIVersionOK {

	return &APIVersionOK{}
}

// WithXRequestID adds the xRequestId to the api version o k response
func (o *APIVersionOK) WithXRequestID(xRequestID string) *APIVersionOK {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the api version o k response
func (o *APIVersionOK) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the api version o k response
func (o *APIVersionOK) WithPayload(payload *models.APIVersion) *APIVersionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the api version o k response
func (o *APIVersionOK) SetPayload(payload *models.APIVersion) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *APIVersionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*APIVersionDefault unexpected error

swagger:response apiVersionDefault
*/
type APIVersionDefault struct {
	_statusCode int
	/*The request id this is a response to

	 */
	XRequestID string `json:"X-Request-Id"`

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewAPIVersionDefault creates APIVersionDefault with default headers values
func NewAPIVersionDefault(code int) *APIVersionDefault {
	if code <= 0 {
		code = 500
	}

	return &APIVersionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the api version default response
func (o *APIVersionDefault) WithStatusCode(code int) *APIVersionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the api version default response
func (o *APIVersionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithXRequestID adds the xRequestId to the api version default response
func (o *APIVersionDefault) WithXRequestID(xRequestID string) *APIVersionDefault {
	o.XRequestID = xRequestID
	return o
}

// SetXRequestID sets the xRequestId to the api version default response
func (o *APIVersionDefault) SetXRequestID(xRequestID string) {
	o.XRequestID = xRequestID
}

// WithPayload adds the payload to the api version default response
func (o *APIVersionDefault) WithPayload(payload *models.ErrorResponse) *APIVersionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the api version default response
func (o *APIVersionDefault) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *APIVersionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Request-Id

	xRequestID := o.XRequestID
	if xRequestID != "" {
		rw.Header().Set("X-Request-Id", xRequestID)
	}

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}