// Code generated by goa v3.11.0, DO NOT EDIT.
//
// relayAccount client HTTP transport
//
// Command:
// $ goa gen admin-panel/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the relayAccount service endpoint HTTP clients.
type Client struct {
	// ListAccount Doer is the HTTP client used to make requests to the listAccount
	// endpoint.
	ListAccountDoer goahttp.Doer

	// RegisterAccount Doer is the HTTP client used to make requests to the
	// registerAccount endpoint.
	RegisterAccountDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the relayAccount service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		ListAccountDoer:     doer,
		RegisterAccountDoer: doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// ListAccount returns an endpoint that makes HTTP requests to the relayAccount
// service listAccount server.
func (c *Client) ListAccount() goa.Endpoint {
	var (
		decodeResponse = DecodeListAccountResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildListAccountRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ListAccountDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("relayAccount", "listAccount", err)
		}
		return decodeResponse(resp)
	}
}

// RegisterAccount returns an endpoint that makes HTTP requests to the
// relayAccount service registerAccount server.
func (c *Client) RegisterAccount() goa.Endpoint {
	var (
		encodeRequest  = EncodeRegisterAccountRequest(c.encoder)
		decodeResponse = DecodeRegisterAccountResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildRegisterAccountRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.RegisterAccountDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("relayAccount", "registerAccount", err)
		}
		return decodeResponse(resp)
	}
}