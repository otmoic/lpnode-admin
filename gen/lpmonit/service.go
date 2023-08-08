// Code generated by goa v3.11.0, DO NOT EDIT.
//
// lpmonit service
//
// Command:
// $ goa gen admin-panel/design

package lpmonit

import (
	"context"
)

// 监控脚本程序
type Service interface {
	// add script and save
	AddScript(context.Context, *AddScriptPayload) (res *AddScriptResult, err error)
	// task_list
	ListScript(context.Context) (res *ListScriptResult, err error)
	// task_list_delete
	DeleteScript(context.Context, *DeleteScriptPayload) (res *DeleteScriptResult, err error)
	// task_run
	RunScript(context.Context, *RunScriptPayload) (res *RunScriptResult, err error)
	// run_result
	RunResult(context.Context, *RunResultPayload) (res *RunResultResult, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "lpmonit"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"add_script", "list_script", "delete_script", "run_script", "run_result"}

// Multipart request Payload
type AddScriptPayload struct {
	Name       *string
	Cron       *string
	ScriptBody *string
}

// AddScriptResult is the result type of the lpmonit service add_script method.
type AddScriptResult struct {
	Code *int64
	// 创建后的Id
	TaskID *string
	// 创建成功后的Id
	Result  *string
	Message *string
}

// DeleteScriptPayload is the payload type of the lpmonit service delete_script
// method.
type DeleteScriptPayload struct {
	// Mongodb 的主键
	ID string
}

// DeleteScriptResult is the result type of the lpmonit service delete_script
// method.
type DeleteScriptResult struct {
	Code *int64
	// 是否删除成功
	Result  *int64
	Message *string
}

// ListScriptResult is the result type of the lpmonit service list_script
// method.
type ListScriptResult struct {
	Code *int64
	// 任务列表
	Result  []*LpMointTaskItem
	Message *string
}

type LpMointTaskItem struct {
	ID   *string
	Name string
	// 定时任务
	Cron string
	// 创建时间戳
	CreatedAt int64
	// 脚本路径
	ScriptPath *string
	// 任务类型
	TaskType string
}

// RunResultPayload is the payload type of the lpmonit service run_result
// method.
type RunResultPayload struct {
	ScriptName *string
}

// RunResultResult is the result type of the lpmonit service run_result method.
type RunResultResult struct {
	Code    *int64
	Result  *string
	Message *string
}

// RunScriptPayload is the payload type of the lpmonit service run_script
// method.
type RunScriptPayload struct {
	ScriptContent *string
}

// RunScriptResult is the result type of the lpmonit service run_script method.
type RunScriptResult struct {
	Code    *int64
	Result  *string
	Message *string
}