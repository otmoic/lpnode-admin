// Code generated by goa v3.11.0, DO NOT EDIT.
//
// accountDex service
//
// Command:
// $ goa gen admin-panel/design

package accountdex

import (
	"context"
)

// Service is the accountDex service interface.
type Service interface {
	// WalletInfo implements walletInfo.
	WalletInfo(context.Context, *WalletInfoPayload) (res *WalletInfoResult, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "accountDex"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"walletInfo"}

type DexAccountBalance struct {
	Token     *string
	TokenName *string
	Amount    *string
	Free      *string
	Locked    *string
	Price     *string
}

// WalletInfoPayload is the payload type of the accountDex service walletInfo
// method.
type WalletInfoPayload struct {
	// 链的Id
	ChainID *int64
}

// WalletInfoResult is the result type of the accountDex service walletInfo
// method.
type WalletInfoResult struct {
	Code *int64
	// 是否成功
	Data    []*DexAccountBalance
	Message *string
}