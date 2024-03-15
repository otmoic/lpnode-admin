// Code generated by goa v3.11.0, DO NOT EDIT.
//
// configResource HTTP client types
//
// Command:
// $ goa gen admin-panel/design

package client

import (
	configresource "admin-panel/gen/config_resource"

	goa "goa.design/goa/v3/pkg"
)

// CreateResourceRequestBody is the type of the "configResource" service
// "createResource" endpoint HTTP request body.
type CreateResourceRequestBody struct {
	AppName  string  `form:"appName" json:"appName" xml:"appName"`
	Version  *string `form:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	ClientID string  `form:"clientId" json:"clientId" xml:"clientId"`
	Template *string `form:"template,omitempty" json:"template,omitempty" xml:"template,omitempty"`
}

// GetResourceRequestBody is the type of the "configResource" service
// "getResource" endpoint HTTP request body.
type GetResourceRequestBody struct {
	ClientID string `form:"clientId" json:"clientId" xml:"clientId"`
}

// EditResultRequestBody is the type of the "configResource" service
// "editResult" endpoint HTTP request body.
type EditResultRequestBody struct {
	TemplateResult string  `form:"templateResult" json:"templateResult" xml:"templateResult"`
	Template       *string `form:"template,omitempty" json:"template,omitempty" xml:"template,omitempty"`
	ClientID       string  `form:"clientId" json:"clientId" xml:"clientId"`
	AppName        *string `form:"appName,omitempty" json:"appName,omitempty" xml:"appName,omitempty"`
	Version        *string `form:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	VersionHash    *string `form:"versionHash,omitempty" json:"versionHash,omitempty" xml:"versionHash,omitempty"`
}

// CreateResourceResponseBody is the type of the "configResource" service
// "createResource" endpoint HTTP response body.
type CreateResourceResponseBody struct {
	// 0 is success
	Code    *int64                          `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	Result  *ConfigResultIDItemResponseBody `form:"result,omitempty" json:"result,omitempty" xml:"result,omitempty"`
	Message *string                         `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// GetResourceResponseBody is the type of the "configResource" service
// "getResource" endpoint HTTP response body.
type GetResourceResponseBody struct {
	// 0 is success
	Code    *int64                            `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	Result  *ResourceConfigResultResponseBody `form:"result,omitempty" json:"result,omitempty" xml:"result,omitempty"`
	Message *string                           `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ListResourceResponseBody is the type of the "configResource" service
// "listResource" endpoint HTTP response body.
type ListResourceResponseBody struct {
	Code    *int64                              `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	Result  []*ResourceConfigResultResponseBody `form:"result,omitempty" json:"result,omitempty" xml:"result,omitempty"`
	Message *string                             `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// DeleteResultResponseBody is the type of the "configResource" service
// "deleteResult" endpoint HTTP response body.
type DeleteResultResponseBody struct {
	Code *int64 `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
}

// EditResultResponseBody is the type of the "configResource" service
// "editResult" endpoint HTTP response body.
type EditResultResponseBody struct {
	Code    *int64  `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// update affected id
	Result *string `form:"result,omitempty" json:"result,omitempty" xml:"result,omitempty"`
}

// ConfigResultIDItemResponseBody is used to define fields on response body
// types.
type ConfigResultIDItemResponseBody struct {
	// mongodb primary key
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// clientId
	ClientID *string `form:"clientId,omitempty" json:"clientId,omitempty" xml:"clientId,omitempty"`
}

// ResourceConfigResultResponseBody is used to define fields on response body
// types.
type ResourceConfigResultResponseBody struct {
	ID             *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	TemplateResult *string `form:"templateResult,omitempty" json:"templateResult,omitempty" xml:"templateResult,omitempty"`
	Template       *string `form:"template,omitempty" json:"template,omitempty" xml:"template,omitempty"`
	ClientID       *string `form:"clientId,omitempty" json:"clientId,omitempty" xml:"clientId,omitempty"`
	AppName        *string `form:"appName,omitempty" json:"appName,omitempty" xml:"appName,omitempty"`
	Version        *string `form:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	VersionHash    *string `form:"versionHash,omitempty" json:"versionHash,omitempty" xml:"versionHash,omitempty"`
}

// NewCreateResourceRequestBody builds the HTTP request body from the payload
// of the "createResource" endpoint of the "configResource" service.
func NewCreateResourceRequestBody(p *configresource.CreateResourcePayload) *CreateResourceRequestBody {
	body := &CreateResourceRequestBody{
		AppName:  p.AppName,
		Version:  p.Version,
		ClientID: p.ClientID,
		Template: p.Template,
	}
	return body
}

// NewGetResourceRequestBody builds the HTTP request body from the payload of
// the "getResource" endpoint of the "configResource" service.
func NewGetResourceRequestBody(p *configresource.GetResourcePayload) *GetResourceRequestBody {
	body := &GetResourceRequestBody{
		ClientID: p.ClientID,
	}
	return body
}

// NewEditResultRequestBody builds the HTTP request body from the payload of
// the "editResult" endpoint of the "configResource" service.
func NewEditResultRequestBody(p *configresource.EditResultPayload) *EditResultRequestBody {
	body := &EditResultRequestBody{
		TemplateResult: p.TemplateResult,
		Template:       p.Template,
		ClientID:       p.ClientID,
		AppName:        p.AppName,
		Version:        p.Version,
		VersionHash:    p.VersionHash,
	}
	return body
}

// NewCreateResourceResultOK builds a "configResource" service "createResource"
// endpoint result from a HTTP "OK" response.
func NewCreateResourceResultOK(body *CreateResourceResponseBody) *configresource.CreateResourceResult {
	v := &configresource.CreateResourceResult{
		Code:    body.Code,
		Message: body.Message,
	}
	if body.Result != nil {
		v.Result = unmarshalConfigResultIDItemResponseBodyToConfigresourceConfigResultIDItem(body.Result)
	}

	return v
}

// NewGetResourceResultOK builds a "configResource" service "getResource"
// endpoint result from a HTTP "OK" response.
func NewGetResourceResultOK(body *GetResourceResponseBody) *configresource.GetResourceResult {
	v := &configresource.GetResourceResult{
		Code:    body.Code,
		Message: body.Message,
	}
	if body.Result != nil {
		v.Result = unmarshalResourceConfigResultResponseBodyToConfigresourceResourceConfigResult(body.Result)
	}

	return v
}

// NewListResourceResultOK builds a "configResource" service "listResource"
// endpoint result from a HTTP "OK" response.
func NewListResourceResultOK(body *ListResourceResponseBody) *configresource.ListResourceResult {
	v := &configresource.ListResourceResult{
		Code:    body.Code,
		Message: body.Message,
	}
	if body.Result != nil {
		v.Result = make([]*configresource.ResourceConfigResult, len(body.Result))
		for i, val := range body.Result {
			v.Result[i] = unmarshalResourceConfigResultResponseBodyToConfigresourceResourceConfigResult(val)
		}
	}

	return v
}

// NewDeleteResultResultOK builds a "configResource" service "deleteResult"
// endpoint result from a HTTP "OK" response.
func NewDeleteResultResultOK(body *DeleteResultResponseBody) *configresource.DeleteResultResult {
	v := &configresource.DeleteResultResult{
		Code: body.Code,
	}

	return v
}

// NewEditResultResultOK builds a "configResource" service "editResult"
// endpoint result from a HTTP "OK" response.
func NewEditResultResultOK(body *EditResultResponseBody) *configresource.EditResultResult {
	v := &configresource.EditResultResult{
		Code:    body.Code,
		Message: body.Message,
		Result:  body.Result,
	}

	return v
}

// ValidateCreateResourceResponseBody runs the validations defined on
// CreateResourceResponseBody
func ValidateCreateResourceResponseBody(body *CreateResourceResponseBody) (err error) {
	if body.Result != nil {
		if err2 := ValidateConfigResultIDItemResponseBody(body.Result); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateGetResourceResponseBody runs the validations defined on
// GetResourceResponseBody
func ValidateGetResourceResponseBody(body *GetResourceResponseBody) (err error) {
	if body.Result != nil {
		if err2 := ValidateResourceConfigResultResponseBody(body.Result); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateListResourceResponseBody runs the validations defined on
// ListResourceResponseBody
func ValidateListResourceResponseBody(body *ListResourceResponseBody) (err error) {
	for _, e := range body.Result {
		if e != nil {
			if err2 := ValidateResourceConfigResultResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateConfigResultIDItemResponseBody runs the validations defined on
// ConfigResultIdItemResponseBody
func ValidateConfigResultIDItemResponseBody(body *ConfigResultIDItemResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	return
}

// ValidateResourceConfigResultResponseBody runs the validations defined on
// ResourceConfigResultResponseBody
func ValidateResourceConfigResultResponseBody(body *ResourceConfigResultResponseBody) (err error) {
	if body.ClientID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("clientId", "body"))
	}
	return
}
