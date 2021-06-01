// Code generated by go-swagger; DO NOT EDIT.

package refplant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetRefPlantsHandlerFunc turns a function with the right signature into a get ref plants handler
type GetRefPlantsHandlerFunc func(GetRefPlantsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRefPlantsHandlerFunc) Handle(params GetRefPlantsParams) middleware.Responder {
	return fn(params)
}

// GetRefPlantsHandler interface for that can handle valid get ref plants params
type GetRefPlantsHandler interface {
	Handle(GetRefPlantsParams) middleware.Responder
}

// NewGetRefPlants creates a new http.Handler for the get ref plants operation
func NewGetRefPlants(ctx *middleware.Context, handler GetRefPlantsHandler) *GetRefPlants {
	return &GetRefPlants{Context: ctx, Handler: handler}
}

/* GetRefPlants swagger:route GET /api/v1/refplants refplant getRefPlants

find reference plants

find reference plants by parameters or all

*/
type GetRefPlants struct {
	Context *middleware.Context
	Handler GetRefPlantsHandler
}

func (o *GetRefPlants) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetRefPlantsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}