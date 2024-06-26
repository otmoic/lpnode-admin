// Code generated by goa v3.11.0, DO NOT EDIT.
//
// taskManager HTTP client CLI support package
//
// Command:
// $ goa gen admin-panel/design

package client

import (
	taskmanager "admin-panel/gen/task_manager"
	"encoding/json"
	"fmt"
)

// BuildTaskDeployPayload builds the payload for the taskManager taskDeploy
// endpoint from CLI flags.
func BuildTaskDeployPayload(taskManagerTaskDeployBody string) (*taskmanager.TaskDeploy2, error) {
	var err error
	var body TaskDeployRequestBody
	{
		err = json.Unmarshal([]byte(taskManagerTaskDeployBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"_id\": \"Totam saepe ad et.\"\n   }'")
		}
	}
	v := &taskmanager.TaskDeploy2{
		ID: body.ID,
	}

	return v, nil
}

// BuildUnDeployPayload builds the payload for the taskManager unDeploy
// endpoint from CLI flags.
func BuildUnDeployPayload(taskManagerUnDeployBody string) (*taskmanager.TaskDeploy2, error) {
	var err error
	var body UnDeployRequestBody
	{
		err = json.Unmarshal([]byte(taskManagerUnDeployBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"_id\": \"Non labore animi.\"\n   }'")
		}
	}
	v := &taskmanager.TaskDeploy2{
		ID: body.ID,
	}

	return v, nil
}

// BuildTaskCreatePayload builds the payload for the taskManager taskCreate
// endpoint from CLI flags.
func BuildTaskCreatePayload(taskManagerTaskCreateBody string) (*taskmanager.TaskItem, error) {
	var err error
	var body TaskCreateRequestBody
	{
		err = json.Unmarshal([]byte(taskManagerTaskCreateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"_id\": \"Eos placeat.\",\n      \"deployMessage\": \"Aliquam dicta molestiae qui sed quia.\",\n      \"deployed\": false,\n      \"schedule\": \"Aut nihil dolorem natus dolorum ut.\",\n      \"scriptBody\": \"Aperiam aut explicabo sint est voluptas.\",\n      \"scriptPath\": \"Enim quos deleniti vero ipsa omnis.\",\n      \"taskType\": \"customize\"\n   }'")
		}
	}
	v := &taskmanager.TaskItem{
		ID:            body.ID,
		Schedule:      body.Schedule,
		TaskType:      body.TaskType,
		Deployed:      body.Deployed,
		DeployMessage: body.DeployMessage,
		ScriptPath:    body.ScriptPath,
		ScriptBody:    body.ScriptBody,
	}

	return v, nil
}
