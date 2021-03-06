// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetRecordsHandlerFunc turns a function with the right signature into a get records handler
type GetRecordsHandlerFunc func(GetRecordsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRecordsHandlerFunc) Handle(params GetRecordsParams) middleware.Responder {
	return fn(params)
}

// GetRecordsHandler interface for that can handle valid get records params
type GetRecordsHandler interface {
	Handle(GetRecordsParams) middleware.Responder
}

// NewGetRecords creates a new http.Handler for the get records operation
func NewGetRecords(ctx *middleware.Context, handler GetRecordsHandler) *GetRecords {
	return &GetRecords{Context: ctx, Handler: handler}
}

/*GetRecords swagger:route GET /getRecords getRecords

get all records from dogechain by transactionId. Served as an array of transmission_types containing a name and an array of decrypted records

*/
type GetRecords struct {
	Context *middleware.Context
	Handler GetRecordsHandler
}

func (o *GetRecords) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRecordsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
