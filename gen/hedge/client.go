// Code generated by goa v3.11.0, DO NOT EDIT.
//
// hedge client
//
// Command:
// $ goa gen admin-panel/design

package hedge

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "hedge" service client.
type Client struct {
	ListEndpoint goa.Endpoint
	EditEndpoint goa.Endpoint
	DelEndpoint  goa.Endpoint
}

// NewClient initializes a "hedge" service client given the endpoints.
func NewClient(list, edit, del goa.Endpoint) *Client {
	return &Client{
		ListEndpoint: list,
		EditEndpoint: edit,
		DelEndpoint:  del,
	}
}

// List calls the "list" endpoint of the "hedge" service.
func (c *Client) List(ctx context.Context) (res *ListResult, err error) {
	var ires interface{}
	ires, err = c.ListEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*ListResult), nil
}

// Edit calls the "edit" endpoint of the "hedge" service.
func (c *Client) Edit(ctx context.Context, p *EditPayload) (res *EditResult, err error) {
	var ires interface{}
	ires, err = c.EditEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*EditResult), nil
}

// Del calls the "del" endpoint of the "hedge" service.
func (c *Client) Del(ctx context.Context, p *DelPayload) (res *DelResult, err error) {
	var ires interface{}
	ires, err = c.DelEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*DelResult), nil
}