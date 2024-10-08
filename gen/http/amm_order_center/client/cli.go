// Code generated by goa v3.11.0, DO NOT EDIT.
//
// ammOrderCenter HTTP client CLI support package
//
// Command:
// $ goa gen admin-panel/design

package client

import (
	ammordercenter "admin-panel/gen/amm_order_center"
	"encoding/json"
	"fmt"
)

// BuildListPayload builds the payload for the ammOrderCenter list endpoint
// from CLI flags.
func BuildListPayload(ammOrderCenterListBody string) (*ammordercenter.ListPayload, error) {
	var err error
	var body ListRequestBody
	{
		err = json.Unmarshal([]byte(ammOrderCenterListBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"ammName\": \"Quibusdam inventore ad possimus non sit sit.\",\n      \"page\": 2031005260904629279,\n      \"pageSize\": 5465095945942461027,\n      \"status\": 3051054241686735991\n   }'")
		}
	}
	v := &ammordercenter.ListPayload{
		Status:   body.Status,
		AmmName:  body.AmmName,
		Page:     body.Page,
		PageSize: body.PageSize,
	}

	return v, nil
}
