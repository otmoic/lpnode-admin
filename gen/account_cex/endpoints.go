// Code generated by goa v3.11.0, DO NOT EDIT.
//
// accountCex endpoints
//
// Command:
// $ goa gen admin-panel/design

package accountcex

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "accountCex" service endpoints.
type Endpoints struct {
	WalletInfo goa.Endpoint
}

// NewEndpoints wraps the methods of the "accountCex" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		WalletInfo: NewWalletInfoEndpoint(s),
	}
}

// Use applies the given middleware to all the "accountCex" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.WalletInfo = m(e.WalletInfo)
}

// NewWalletInfoEndpoint returns an endpoint function that calls the method
// "walletInfo" of service "accountCex".
func NewWalletInfoEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.WalletInfo(ctx)
	}
}