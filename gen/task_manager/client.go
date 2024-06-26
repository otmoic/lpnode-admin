// Code generated by goa v3.11.0, DO NOT EDIT.
//
// taskManager client
//
// Command:
// $ goa gen admin-panel/design

package taskmanager

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "taskManager" service client.
type Client struct {
	TaskListEndpoint   goa.Endpoint
	TaskDeployEndpoint goa.Endpoint
	UnDeployEndpoint   goa.Endpoint
	TaskCreateEndpoint goa.Endpoint
}

// NewClient initializes a "taskManager" service client given the endpoints.
func NewClient(taskList, taskDeploy, unDeploy, taskCreate goa.Endpoint) *Client {
	return &Client{
		TaskListEndpoint:   taskList,
		TaskDeployEndpoint: taskDeploy,
		UnDeployEndpoint:   unDeploy,
		TaskCreateEndpoint: taskCreate,
	}
}

// TaskList calls the "taskList" endpoint of the "taskManager" service.
func (c *Client) TaskList(ctx context.Context) (res *TaskListResult, err error) {
	var ires interface{}
	ires, err = c.TaskListEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*TaskListResult), nil
}

// TaskDeploy calls the "taskDeploy" endpoint of the "taskManager" service.
func (c *Client) TaskDeploy(ctx context.Context, p *TaskDeploy2) (res *TaskDeployResult, err error) {
	var ires interface{}
	ires, err = c.TaskDeployEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*TaskDeployResult), nil
}

// UnDeploy calls the "unDeploy" endpoint of the "taskManager" service.
func (c *Client) UnDeploy(ctx context.Context, p *TaskDeploy2) (res *UnDeployResult, err error) {
	var ires interface{}
	ires, err = c.UnDeployEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*UnDeployResult), nil
}

// TaskCreate calls the "taskCreate" endpoint of the "taskManager" service.
func (c *Client) TaskCreate(ctx context.Context, p *TaskItem) (res *TaskCreateResult, err error) {
	var ires interface{}
	ires, err = c.TaskCreateEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*TaskCreateResult), nil
}
