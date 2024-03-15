// Code generated by goa v3.11.0, DO NOT EDIT.
//
// accountDex HTTP server types
//
// Command:
// $ goa gen admin-panel/design

package server

import (
	accountdex "admin-panel/gen/account_dex"
)

// WalletInfoRequestBody is the type of the "accountDex" service "walletInfo"
// endpoint HTTP request body.
type WalletInfoRequestBody struct {
	// chain Id
	ChainID *int64 `form:"chainId,omitempty" json:"chainId,omitempty" xml:"chainId,omitempty"`
}

// WalletInfoResponseBody is the type of the "accountDex" service "walletInfo"
// endpoint HTTP response body.
type WalletInfoResponseBody struct {
	Code *int64 `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// result
	Data    []*DexAccountBalanceResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	Message *string                          `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// DexAccountBalanceResponseBody is used to define fields on response body
// types.
type DexAccountBalanceResponseBody struct {
	Token     *string `form:"token,omitempty" json:"token,omitempty" xml:"token,omitempty"`
	TokenName *string `form:"tokenName,omitempty" json:"tokenName,omitempty" xml:"tokenName,omitempty"`
	Amount    *string `form:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
	Free      *string `form:"free,omitempty" json:"free,omitempty" xml:"free,omitempty"`
	Locked    *string `form:"locked,omitempty" json:"locked,omitempty" xml:"locked,omitempty"`
	Price     *string `form:"price,omitempty" json:"price,omitempty" xml:"price,omitempty"`
}

// NewWalletInfoResponseBody builds the HTTP response body from the result of
// the "walletInfo" endpoint of the "accountDex" service.
func NewWalletInfoResponseBody(res *accountdex.WalletInfoResult) *WalletInfoResponseBody {
	body := &WalletInfoResponseBody{
		Code:    res.Code,
		Message: res.Message,
	}
	if res.Data != nil {
		body.Data = make([]*DexAccountBalanceResponseBody, len(res.Data))
		for i, val := range res.Data {
			body.Data[i] = marshalAccountdexDexAccountBalanceToDexAccountBalanceResponseBody(val)
		}
	}
	return body
}

// NewWalletInfoPayload builds a accountDex service walletInfo endpoint payload.
func NewWalletInfoPayload(body *WalletInfoRequestBody) *accountdex.WalletInfoPayload {
	v := &accountdex.WalletInfoPayload{
		ChainID: body.ChainID,
	}

	return v
}
