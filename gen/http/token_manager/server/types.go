// Code generated by goa v3.11.0, DO NOT EDIT.
//
// tokenManager HTTP server types
//
// Command:
// $ goa gen admin-panel/design

package server

import (
	tokenmanager "admin-panel/gen/token_manager"

	goa "goa.design/goa/v3/pkg"
)

// TokenCreateRequestBody is the type of the "tokenManager" service
// "tokenCreate" endpoint HTTP request body.
type TokenCreateRequestBody struct {
	ID      *string `form:"_id,omitempty" json:"_id,omitempty" xml:"_id,omitempty"`
	TokenID *string `form:"tokenId,omitempty" json:"tokenId,omitempty" xml:"tokenId,omitempty"`
	// chain id
	ChainID *int64 `form:"chainId,omitempty" json:"chainId,omitempty" xml:"chainId,omitempty"`
	// token address
	Address *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	// name in token contract
	TokenName *string `form:"tokenName,omitempty" json:"tokenName,omitempty" xml:"tokenName,omitempty"`
	// corresponding name in cex
	MarketName *string `form:"marketName,omitempty" json:"marketName,omitempty" xml:"marketName,omitempty"`
	Precision  *int64  `form:"precision,omitempty" json:"precision,omitempty" xml:"precision,omitempty"`
	ChainType  *string `form:"chainType,omitempty" json:"chainType,omitempty" xml:"chainType,omitempty"`
	CoinType   *string `form:"coinType,omitempty" json:"coinType,omitempty" xml:"coinType,omitempty"`
}

// TokenDeleteRequestBody is the type of the "tokenManager" service
// "tokenDelete" endpoint HTTP request body.
type TokenDeleteRequestBody struct {
	// mongodb primary key
	ID *string `form:"_id,omitempty" json:"_id,omitempty" xml:"_id,omitempty"`
}

// TokenListResponseBody is the type of the "tokenManager" service "tokenList"
// endpoint HTTP response body.
type TokenListResponseBody struct {
	Code *int64 `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// list
	Result  []*TokenItemResponseBody `form:"result,omitempty" json:"result,omitempty" xml:"result,omitempty"`
	Message *string                  `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// TokenCreateResponseBody is the type of the "tokenManager" service
// "tokenCreate" endpoint HTTP response body.
type TokenCreateResponseBody struct {
	Code    *int64  `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	Result  *int64  `form:"result,omitempty" json:"result,omitempty" xml:"result,omitempty"`
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// TokenDeleteResponseBody is the type of the "tokenManager" service
// "tokenDelete" endpoint HTTP response body.
type TokenDeleteResponseBody struct {
	Code    *int64  `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	Result  *int64  `form:"result,omitempty" json:"result,omitempty" xml:"result,omitempty"`
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// TokenItemResponseBody is used to define fields on response body types.
type TokenItemResponseBody struct {
	ID      *string `form:"_id,omitempty" json:"_id,omitempty" xml:"_id,omitempty"`
	TokenID *string `form:"tokenId,omitempty" json:"tokenId,omitempty" xml:"tokenId,omitempty"`
	// chain id
	ChainID int64 `form:"chainId" json:"chainId" xml:"chainId"`
	// token address
	Address string `form:"address" json:"address" xml:"address"`
	// name in token contract
	TokenName *string `form:"tokenName,omitempty" json:"tokenName,omitempty" xml:"tokenName,omitempty"`
	// corresponding name in cex
	MarketName string  `form:"marketName" json:"marketName" xml:"marketName"`
	Precision  int64   `form:"precision" json:"precision" xml:"precision"`
	ChainType  *string `form:"chainType,omitempty" json:"chainType,omitempty" xml:"chainType,omitempty"`
	CoinType   string  `form:"coinType" json:"coinType" xml:"coinType"`
}

// NewTokenListResponseBody builds the HTTP response body from the result of
// the "tokenList" endpoint of the "tokenManager" service.
func NewTokenListResponseBody(res *tokenmanager.TokenListResult) *TokenListResponseBody {
	body := &TokenListResponseBody{
		Code:    res.Code,
		Message: res.Message,
	}
	if res.Result != nil {
		body.Result = make([]*TokenItemResponseBody, len(res.Result))
		for i, val := range res.Result {
			body.Result[i] = marshalTokenmanagerTokenItemToTokenItemResponseBody(val)
		}
	}
	return body
}

// NewTokenCreateResponseBody builds the HTTP response body from the result of
// the "tokenCreate" endpoint of the "tokenManager" service.
func NewTokenCreateResponseBody(res *tokenmanager.TokenCreateResult) *TokenCreateResponseBody {
	body := &TokenCreateResponseBody{
		Code:    res.Code,
		Result:  res.Result,
		Message: res.Message,
	}
	return body
}

// NewTokenDeleteResponseBody builds the HTTP response body from the result of
// the "tokenDelete" endpoint of the "tokenManager" service.
func NewTokenDeleteResponseBody(res *tokenmanager.TokenDeleteResult) *TokenDeleteResponseBody {
	body := &TokenDeleteResponseBody{
		Code:    res.Code,
		Result:  res.Result,
		Message: res.Message,
	}
	return body
}

// NewTokenCreateTokenItem builds a tokenManager service tokenCreate endpoint
// payload.
func NewTokenCreateTokenItem(body *TokenCreateRequestBody) *tokenmanager.TokenItem {
	v := &tokenmanager.TokenItem{
		ID:         body.ID,
		TokenID:    body.TokenID,
		ChainID:    *body.ChainID,
		Address:    *body.Address,
		TokenName:  body.TokenName,
		MarketName: *body.MarketName,
		Precision:  *body.Precision,
		ChainType:  body.ChainType,
		CoinType:   *body.CoinType,
	}

	return v
}

// NewTokenDeleteDeleteTokenFilter builds a tokenManager service tokenDelete
// endpoint payload.
func NewTokenDeleteDeleteTokenFilter(body *TokenDeleteRequestBody) *tokenmanager.DeleteTokenFilter {
	v := &tokenmanager.DeleteTokenFilter{
		ID: *body.ID,
	}

	return v
}

// ValidateTokenCreateRequestBody runs the validations defined on
// TokenCreateRequestBody
func ValidateTokenCreateRequestBody(body *TokenCreateRequestBody) (err error) {
	if body.ChainID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("chainId", "body"))
	}
	if body.Address == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("address", "body"))
	}
	if body.MarketName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("marketName", "body"))
	}
	if body.Precision == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("precision", "body"))
	}
	if body.CoinType == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("coinType", "body"))
	}
	if body.Precision != nil {
		if *body.Precision < 6 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.precision", *body.Precision, 6, true))
		}
	}
	if body.Precision != nil {
		if *body.Precision > 18 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.precision", *body.Precision, 18, false))
		}
	}
	if body.CoinType != nil {
		if !(*body.CoinType == "stable_coin" || *body.CoinType == "coin") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.coinType", *body.CoinType, []interface{}{"stable_coin", "coin"}))
		}
	}
	return
}

// ValidateTokenDeleteRequestBody runs the validations defined on
// TokenDeleteRequestBody
func ValidateTokenDeleteRequestBody(body *TokenDeleteRequestBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("_id", "body"))
	}
	return
}
