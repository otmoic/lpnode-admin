// Code generated by goa v3.11.0, DO NOT EDIT.
//
// hedge HTTP client CLI support package
//
// Command:
// $ goa gen admin-panel/design

package client

import (
	hedge "admin-panel/gen/hedge"
	"encoding/json"
	"fmt"
)

// BuildEditPayload builds the payload for the hedge edit endpoint from CLI
// flags.
func BuildEditPayload(hedgeEditBody string) (*hedge.EditPayload, error) {
	var err error
	var body EditRequestBody
	{
		err = json.Unmarshal([]byte(hedgeEditBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"hedge\": {\n         \"hedgeType\": \"Quod perspiciatis doloremque.\",\n         \"id\": \"Ratione sapiente quas impedit explicabo consectetur.\"\n      }\n   }'")
		}
	}
	v := &hedge.EditPayload{}
	if body.Hedge != nil {
		v.Hedge = marshalHedgeItemRequestBodyToHedgeHedgeItem(body.Hedge)
	}

	return v, nil
}

// BuildDelPayload builds the payload for the hedge del endpoint from CLI flags.
func BuildDelPayload(hedgeDelBody string) (*hedge.DelPayload, error) {
	var err error
	var body DelRequestBody
	{
		err = json.Unmarshal([]byte(hedgeDelBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": \"Qui culpa nostrum dicta eveniet nihil.\"\n   }'")
		}
	}
	v := &hedge.DelPayload{
		ID: body.ID,
	}

	return v, nil
}
