// Code generated by goa v3.11.0, DO NOT EDIT.
//
// configResource client
//
// Command:
// $ goa gen admin-panel/design

package configresource

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "configResource" service client.
type Client struct {
	CreateResourceEndpoint goa.Endpoint
	GetResourceEndpoint    goa.Endpoint
	ListResourceEndpoint   goa.Endpoint
	DeleteResultEndpoint   goa.Endpoint
	EditResultEndpoint     goa.Endpoint
}

// NewClient initializes a "configResource" service client given the endpoints.
func NewClient(createResource, getResource, listResource, deleteResult, editResult goa.Endpoint) *Client {
	return &Client{
		CreateResourceEndpoint: createResource,
		GetResourceEndpoint:    getResource,
		ListResourceEndpoint:   listResource,
		DeleteResultEndpoint:   deleteResult,
		EditResultEndpoint:     editResult,
	}
}

// CreateResource calls the "createResource" endpoint of the "configResource"
// service.
func (c *Client) CreateResource(ctx context.Context, p *CreateResourcePayload) (res *CreateResourceResult, err error) {
	var ires interface{}
	ires, err = c.CreateResourceEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*CreateResourceResult), nil
}

// GetResource calls the "getResource" endpoint of the "configResource" service.
func (c *Client) GetResource(ctx context.Context, p *GetResourcePayload) (res *GetResourceResult, err error) {
	var ires interface{}
	ires, err = c.GetResourceEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*GetResourceResult), nil
}

// ListResource calls the "listResource" endpoint of the "configResource"
// service.
func (c *Client) ListResource(ctx context.Context) (res *ListResourceResult, err error) {
	var ires interface{}
	ires, err = c.ListResourceEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*ListResourceResult), nil
}

// DeleteResult calls the "deleteResult" endpoint of the "configResource"
// service.
func (c *Client) DeleteResult(ctx context.Context) (res *DeleteResultResult, err error) {
	var ires interface{}
	ires, err = c.DeleteResultEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*DeleteResultResult), nil
}

// EditResult calls the "editResult" endpoint of the "configResource" service.
func (c *Client) EditResult(ctx context.Context, p *EditResultPayload) (res *EditResultResult, err error) {
	var ires interface{}
	ires, err = c.EditResultEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*EditResultResult), nil
}
