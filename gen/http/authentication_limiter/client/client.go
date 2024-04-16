// Code generated by goa v3.11.0, DO NOT EDIT.
//
// authenticationLimiter client HTTP transport
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

// Client lists the authenticationLimiter service endpoint HTTP clients.
type Client struct {
	// GetAuthenticationLimiter Doer is the HTTP client used to make requests to
	// the getAuthenticationLimiter endpoint.
	GetAuthenticationLimiterDoer goahttp.Doer

	// SetAuthenticationLimiter Doer is the HTTP client used to make requests to
	// the setAuthenticationLimiter endpoint.
	SetAuthenticationLimiterDoer goahttp.Doer

	// DelAuthenticationLimiter Doer is the HTTP client used to make requests to
	// the delAuthenticationLimiter endpoint.
	DelAuthenticationLimiterDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the authenticationLimiter
// service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		GetAuthenticationLimiterDoer: doer,
		SetAuthenticationLimiterDoer: doer,
		DelAuthenticationLimiterDoer: doer,
		RestoreResponseBody:          restoreBody,
		scheme:                       scheme,
		host:                         host,
		decoder:                      dec,
		encoder:                      enc,
	}
}

// GetAuthenticationLimiter returns an endpoint that makes HTTP requests to the
// authenticationLimiter service getAuthenticationLimiter server.
func (c *Client) GetAuthenticationLimiter() goa.Endpoint {
	var (
		decodeResponse = DecodeGetAuthenticationLimiterResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetAuthenticationLimiterRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetAuthenticationLimiterDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("authenticationLimiter", "getAuthenticationLimiter", err)
		}
		return decodeResponse(resp)
	}
}

// SetAuthenticationLimiter returns an endpoint that makes HTTP requests to the
// authenticationLimiter service setAuthenticationLimiter server.
func (c *Client) SetAuthenticationLimiter() goa.Endpoint {
	var (
		encodeRequest  = EncodeSetAuthenticationLimiterRequest(c.encoder)
		decodeResponse = DecodeSetAuthenticationLimiterResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildSetAuthenticationLimiterRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.SetAuthenticationLimiterDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("authenticationLimiter", "setAuthenticationLimiter", err)
		}
		return decodeResponse(resp)
	}
}

// DelAuthenticationLimiter returns an endpoint that makes HTTP requests to the
// authenticationLimiter service delAuthenticationLimiter server.
func (c *Client) DelAuthenticationLimiter() goa.Endpoint {
	var (
		decodeResponse = DecodeDelAuthenticationLimiterResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDelAuthenticationLimiterRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DelAuthenticationLimiterDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("authenticationLimiter", "delAuthenticationLimiter", err)
		}
		return decodeResponse(resp)
	}
}