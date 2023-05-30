// Code generated by goa v3.11.0, DO NOT EDIT.
//
// relayAccount HTTP server types
//
// Command:
// $ goa gen admin-panel/design

package server

import (
	relayaccount "admin-panel/gen/relay_account"

	goa "goa.design/goa/v3/pkg"
)

// RegisterAccountRequestBody is the type of the "relayAccount" service
// "registerAccount" endpoint HTTP request body.
type RegisterAccountRequestBody struct {
	Name    *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Profile *string `form:"profile,omitempty" json:"profile,omitempty" xml:"profile,omitempty"`
}

// DeleteAccountRequestBody is the type of the "relayAccount" service
// "deleteAccount" endpoint HTTP request body.
type DeleteAccountRequestBody struct {
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// ListAccountResponseBody is the type of the "relayAccount" service
// "listAccount" endpoint HTTP response body.
type ListAccountResponseBody struct {
	Code    int64                           `form:"code" json:"code" xml:"code"`
	Result  []*RelayAccountItemResponseBody `form:"result,omitempty" json:"result,omitempty" xml:"result,omitempty"`
	Message *string                         `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// RegisterAccountResponseBody is the type of the "relayAccount" service
// "registerAccount" endpoint HTTP response body.
type RegisterAccountResponseBody struct {
	Code    *int64                        `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	Result  *RelayAccountItemResponseBody `form:"result,omitempty" json:"result,omitempty" xml:"result,omitempty"`
	Message *string                       `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// DeleteAccountResponseBody is the type of the "relayAccount" service
// "deleteAccount" endpoint HTTP response body.
type DeleteAccountResponseBody struct {
	Code    *int64  `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	Result  *string `form:"result,omitempty" json:"result,omitempty" xml:"result,omitempty"`
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// RelayAccountItemResponseBody is used to define fields on response body types.
type RelayAccountItemResponseBody struct {
	ID           *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Name         *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Profile      *string `form:"profile,omitempty" json:"profile,omitempty" xml:"profile,omitempty"`
	LpIDFake     *string `form:"lpIdFake,omitempty" json:"lpIdFake,omitempty" xml:"lpIdFake,omitempty"`
	LpNodeAPIKey *string `form:"lpNodeApiKey,omitempty" json:"lpNodeApiKey,omitempty" xml:"lpNodeApiKey,omitempty"`
	RelayAPIKey  *string `form:"relayApiKey,omitempty" json:"relayApiKey,omitempty" xml:"relayApiKey,omitempty"`
}

// NewListAccountResponseBody builds the HTTP response body from the result of
// the "listAccount" endpoint of the "relayAccount" service.
func NewListAccountResponseBody(res *relayaccount.ListAccountResult) *ListAccountResponseBody {
	body := &ListAccountResponseBody{
		Code:    res.Code,
		Message: res.Message,
	}
	if res.Result != nil {
		body.Result = make([]*RelayAccountItemResponseBody, len(res.Result))
		for i, val := range res.Result {
			body.Result[i] = marshalRelayaccountRelayAccountItemToRelayAccountItemResponseBody(val)
		}
	}
	return body
}

// NewRegisterAccountResponseBody builds the HTTP response body from the result
// of the "registerAccount" endpoint of the "relayAccount" service.
func NewRegisterAccountResponseBody(res *relayaccount.RegisterAccountResult) *RegisterAccountResponseBody {
	body := &RegisterAccountResponseBody{
		Code:    res.Code,
		Message: res.Message,
	}
	if res.Result != nil {
		body.Result = marshalRelayaccountRelayAccountItemToRelayAccountItemResponseBody(res.Result)
	}
	return body
}

// NewDeleteAccountResponseBody builds the HTTP response body from the result
// of the "deleteAccount" endpoint of the "relayAccount" service.
func NewDeleteAccountResponseBody(res *relayaccount.DeleteAccountResult) *DeleteAccountResponseBody {
	body := &DeleteAccountResponseBody{
		Code:    res.Code,
		Result:  res.Result,
		Message: res.Message,
	}
	return body
}

// NewRegisterAccountPayload builds a relayAccount service registerAccount
// endpoint payload.
func NewRegisterAccountPayload(body *RegisterAccountRequestBody) *relayaccount.RegisterAccountPayload {
	v := &relayaccount.RegisterAccountPayload{
		Name:    *body.Name,
		Profile: body.Profile,
	}

	return v
}

// NewDeleteAccountPayload builds a relayAccount service deleteAccount endpoint
// payload.
func NewDeleteAccountPayload(body *DeleteAccountRequestBody) *relayaccount.DeleteAccountPayload {
	v := &relayaccount.DeleteAccountPayload{
		ID: *body.ID,
	}

	return v
}

// ValidateRegisterAccountRequestBody runs the validations defined on
// RegisterAccountRequestBody
func ValidateRegisterAccountRequestBody(body *RegisterAccountRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	return
}

// ValidateDeleteAccountRequestBody runs the validations defined on
// DeleteAccountRequestBody
func ValidateDeleteAccountRequestBody(body *DeleteAccountRequestBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	return
}
