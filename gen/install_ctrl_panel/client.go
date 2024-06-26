// Code generated by goa v3.11.0, DO NOT EDIT.
//
// installCtrlPanel client
//
// Command:
// $ goa gen admin-panel/design

package installctrlpanel

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "installCtrlPanel" service client.
type Client struct {
	ListInstallEndpoint         goa.Endpoint
	InstallLpClientEndpoint     goa.Endpoint
	UninstallLpClientEndpoint   goa.Endpoint
	InstallDeploymentEndpoint   goa.Endpoint
	UninstallDeploymentEndpoint goa.Endpoint
	UpdateDeploymentEndpoint    goa.Endpoint
}

// NewClient initializes a "installCtrlPanel" service client given the
// endpoints.
func NewClient(listInstall, installLpClient, uninstallLpClient, installDeployment, uninstallDeployment, updateDeployment goa.Endpoint) *Client {
	return &Client{
		ListInstallEndpoint:         listInstall,
		InstallLpClientEndpoint:     installLpClient,
		UninstallLpClientEndpoint:   uninstallLpClient,
		InstallDeploymentEndpoint:   installDeployment,
		UninstallDeploymentEndpoint: uninstallDeployment,
		UpdateDeploymentEndpoint:    updateDeployment,
	}
}

// ListInstall calls the "listInstall" endpoint of the "installCtrlPanel"
// service.
func (c *Client) ListInstall(ctx context.Context, p *ListInstallPayload) (res *ListInstallResult, err error) {
	var ires interface{}
	ires, err = c.ListInstallEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ListInstallResult), nil
}

// InstallLpClient calls the "installLpClient" endpoint of the
// "installCtrlPanel" service.
// InstallLpClient may return the following errors:
//   - "an_error" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) InstallLpClient(ctx context.Context, p *InstallLpClientPayload) (res *InstallLpClientResult, err error) {
	var ires interface{}
	ires, err = c.InstallLpClientEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*InstallLpClientResult), nil
}

// UninstallLpClient calls the "uninstallLpClient" endpoint of the
// "installCtrlPanel" service.
func (c *Client) UninstallLpClient(ctx context.Context, p *UninstallLpClientPayload) (res *UninstallLpClientResult, err error) {
	var ires interface{}
	ires, err = c.UninstallLpClientEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*UninstallLpClientResult), nil
}

// InstallDeployment calls the "installDeployment" endpoint of the
// "installCtrlPanel" service.
func (c *Client) InstallDeployment(ctx context.Context, p *InstallDeploymentPayload) (res *InstallDeploymentResult, err error) {
	var ires interface{}
	ires, err = c.InstallDeploymentEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*InstallDeploymentResult), nil
}

// UninstallDeployment calls the "uninstallDeployment" endpoint of the
// "installCtrlPanel" service.
func (c *Client) UninstallDeployment(ctx context.Context, p *UninstallDeploymentPayload) (res *UninstallDeploymentResult, err error) {
	var ires interface{}
	ires, err = c.UninstallDeploymentEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*UninstallDeploymentResult), nil
}

// UpdateDeployment calls the "updateDeployment" endpoint of the
// "installCtrlPanel" service.
func (c *Client) UpdateDeployment(ctx context.Context, p *UpdateDeploymentPayload) (res *UpdateDeploymentResult, err error) {
	var ires interface{}
	ires, err = c.UpdateDeploymentEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*UpdateDeploymentResult), nil
}
