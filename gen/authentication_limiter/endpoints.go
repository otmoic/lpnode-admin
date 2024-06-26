// Code generated by goa v3.11.0, DO NOT EDIT.
//
// authenticationLimiter endpoints
//
// Command:
// $ goa gen admin-panel/design

package authenticationlimiter

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "authenticationLimiter" service endpoints.
type Endpoints struct {
	GetAuthenticationLimiter goa.Endpoint
	SetAuthenticationLimiter goa.Endpoint
	DelAuthenticationLimiter goa.Endpoint
}

// NewEndpoints wraps the methods of the "authenticationLimiter" service with
// endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		GetAuthenticationLimiter: NewGetAuthenticationLimiterEndpoint(s),
		SetAuthenticationLimiter: NewSetAuthenticationLimiterEndpoint(s),
		DelAuthenticationLimiter: NewDelAuthenticationLimiterEndpoint(s),
	}
}

// Use applies the given middleware to all the "authenticationLimiter" service
// endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.GetAuthenticationLimiter = m(e.GetAuthenticationLimiter)
	e.SetAuthenticationLimiter = m(e.SetAuthenticationLimiter)
	e.DelAuthenticationLimiter = m(e.DelAuthenticationLimiter)
}

// NewGetAuthenticationLimiterEndpoint returns an endpoint function that calls
// the method "getAuthenticationLimiter" of service "authenticationLimiter".
func NewGetAuthenticationLimiterEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetAuthenticationLimiter(ctx)
	}
}

// NewSetAuthenticationLimiterEndpoint returns an endpoint function that calls
// the method "setAuthenticationLimiter" of service "authenticationLimiter".
func NewSetAuthenticationLimiterEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SetAuthenticationLimiterPayload)
		return s.SetAuthenticationLimiter(ctx, p)
	}
}

// NewDelAuthenticationLimiterEndpoint returns an endpoint function that calls
// the method "delAuthenticationLimiter" of service "authenticationLimiter".
func NewDelAuthenticationLimiterEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.DelAuthenticationLimiter(ctx)
	}
}
