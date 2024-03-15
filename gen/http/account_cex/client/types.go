// Code generated by goa v3.11.0, DO NOT EDIT.
//
// accountCex HTTP client types
//
// Command:
// $ goa gen admin-panel/design

package client

import (
	accountcex "admin-panel/gen/account_cex"
)

// WalletInfoResponseBody is the type of the "accountCex" service "walletInfo"
// endpoint HTTP response body.
type WalletInfoResponseBody struct {
	Code *int64 `form:"code,omitempty" json:"code,omitempty" xml:"code,omitempty"`
	// result
	Data    []*CexAccountBalanceResponseBody `form:"data,omitempty" json:"data,omitempty" xml:"data,omitempty"`
	Message *string                          `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// CexAccountBalanceResponseBody is used to define fields on response body
// types.
type CexAccountBalanceResponseBody struct {
	Asset  *string `form:"asset,omitempty" json:"asset,omitempty" xml:"asset,omitempty"`
	Total  *string `form:"total,omitempty" json:"total,omitempty" xml:"total,omitempty"`
	Free   *string `form:"free,omitempty" json:"free,omitempty" xml:"free,omitempty"`
	Locked *string `form:"locked,omitempty" json:"locked,omitempty" xml:"locked,omitempty"`
	Price  *string `form:"price,omitempty" json:"price,omitempty" xml:"price,omitempty"`
}

// NewWalletInfoResultOK builds a "accountCex" service "walletInfo" endpoint
// result from a HTTP "OK" response.
func NewWalletInfoResultOK(body *WalletInfoResponseBody) *accountcex.WalletInfoResult {
	v := &accountcex.WalletInfoResult{
		Code:    body.Code,
		Message: body.Message,
	}
	if body.Data != nil {
		v.Data = make([]*accountcex.CexAccountBalance, len(body.Data))
		for i, val := range body.Data {
			v.Data[i] = unmarshalCexAccountBalanceResponseBodyToAccountcexCexAccountBalance(val)
		}
	}

	return v
}
