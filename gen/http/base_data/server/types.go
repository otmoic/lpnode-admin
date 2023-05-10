// Code generated by goa v3.11.0, DO NOT EDIT.
//
// baseData HTTP server types
//
// Command:
// $ goa gen admin-panel/design

package server

import (
	basedata "admin-panel/gen/base_data"
)

// ChainDataListResponseBody is the type of the "baseData" service
// "chainDataList" endpoint HTTP response body.
type ChainDataListResponseBody struct {
	Code *int64 `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// 列表
	Result  []*ChainDataItemResponseBody `form:"result,omitempty" json:"result,omitempty" xml:"result,omitempty"`
	Message *string                      `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ChainDataItemResponseBody is used to define fields on response body types.
type ChainDataItemResponseBody struct {
	// 链在数据库中的id
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// 链的Id
	ChainID *int64 `form:"chainId,omitempty" json:"chainId,omitempty" xml:"chainId,omitempty"`
	// 链名称
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// 链全称
	ChainName *string `form:"chainName,omitempty" json:"chainName,omitempty" xml:"chainName,omitempty"`
	// Token币的名称
	TokenName *string `form:"tokenName,omitempty" json:"tokenName,omitempty" xml:"tokenName,omitempty"`
}

// NewChainDataListResponseBody builds the HTTP response body from the result
// of the "chainDataList" endpoint of the "baseData" service.
func NewChainDataListResponseBody(res *basedata.ChainDataListResult) *ChainDataListResponseBody {
	body := &ChainDataListResponseBody{
		Code:    res.Code,
		Message: res.Message,
	}
	if res.Result != nil {
		body.Result = make([]*ChainDataItemResponseBody, len(res.Result))
		for i, val := range res.Result {
			body.Result[i] = marshalBasedataChainDataItemToChainDataItemResponseBody(val)
		}
	}
	return body
}