// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/hunterBhough/bitmask.api/models"
)

// GetTransactionOKCode is the HTTP code returned for type GetTransactionOK
const GetTransactionOKCode int = 200

/*GetTransactionOK successful operation

swagger:response getTransactionOK
*/
type GetTransactionOK struct {

	/*
	  In: Body
	*/
	Payload []models.Question `json:"body,omitempty"`
}

// NewGetTransactionOK creates GetTransactionOK with default headers values
func NewGetTransactionOK() *GetTransactionOK {
	return &GetTransactionOK{}
}

// WithPayload adds the payload to the get transaction o k response
func (o *GetTransactionOK) WithPayload(payload []models.Question) *GetTransactionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get transaction o k response
func (o *GetTransactionOK) SetPayload(payload []models.Question) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTransactionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]models.Question, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetTransactionBadRequestCode is the HTTP code returned for type GetTransactionBadRequest
const GetTransactionBadRequestCode int = 400

/*GetTransactionBadRequest Invalid or not enough inputs.

swagger:response getTransactionBadRequest
*/
type GetTransactionBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTransactionBadRequest creates GetTransactionBadRequest with default headers values
func NewGetTransactionBadRequest() *GetTransactionBadRequest {
	return &GetTransactionBadRequest{}
}

// WithPayload adds the payload to the get transaction bad request response
func (o *GetTransactionBadRequest) WithPayload(payload *models.Error) *GetTransactionBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get transaction bad request response
func (o *GetTransactionBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTransactionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTransactionTooManyRequestsCode is the HTTP code returned for type GetTransactionTooManyRequests
const GetTransactionTooManyRequestsCode int = 429

/*GetTransactionTooManyRequests Too many requests.

swagger:response getTransactionTooManyRequests
*/
type GetTransactionTooManyRequests struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetTransactionTooManyRequests creates GetTransactionTooManyRequests with default headers values
func NewGetTransactionTooManyRequests() *GetTransactionTooManyRequests {
	return &GetTransactionTooManyRequests{}
}

// WithPayload adds the payload to the get transaction too many requests response
func (o *GetTransactionTooManyRequests) WithPayload(payload *models.Error) *GetTransactionTooManyRequests {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get transaction too many requests response
func (o *GetTransactionTooManyRequests) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTransactionTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(429)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
