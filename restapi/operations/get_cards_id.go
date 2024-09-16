// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetCardsIDHandlerFunc turns a function with the right signature into a get cards ID handler
type GetCardsIDHandlerFunc func(GetCardsIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCardsIDHandlerFunc) Handle(params GetCardsIDParams) middleware.Responder {
	return fn(params)
}

// GetCardsIDHandler interface for that can handle valid get cards ID params
type GetCardsIDHandler interface {
	Handle(GetCardsIDParams) middleware.Responder
}

// NewGetCardsID creates a new http.Handler for the get cards ID operation
func NewGetCardsID(ctx *middleware.Context, handler GetCardsIDHandler) *GetCardsID {
	return &GetCardsID{Context: ctx, Handler: handler}
}

/*
	GetCardsID swagger:route GET /cards/{id} getCardsId

Get a card by ID
*/
type GetCardsID struct {
	Context *middleware.Context
	Handler GetCardsIDHandler
}

func (o *GetCardsID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetCardsIDParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}