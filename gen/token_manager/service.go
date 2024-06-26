// Code generated by goa v3.11.0, DO NOT EDIT.
//
// tokenManager service
//
// Command:
// $ goa gen admin-panel/design

package tokenmanager

import (
	"context"
)

// used to manage all tokens
type Service interface {
	// TokenList implements tokenList.
	TokenList(context.Context) (res *TokenListResult, err error)
	// TokenCreate implements tokenCreate.
	TokenCreate(context.Context, *TokenItem) (res *TokenCreateResult, err error)
	// TokenDelete implements tokenDelete.
	TokenDelete(context.Context, *DeleteTokenFilter) (res *TokenDeleteResult, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "tokenManager"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [3]string{"tokenList", "tokenCreate", "tokenDelete"}

// DeleteTokenFilter is the payload type of the tokenManager service
// tokenDelete method.
type DeleteTokenFilter struct {
	// mongodb primary key
	ID string
}

// TokenCreateResult is the result type of the tokenManager service tokenCreate
// method.
type TokenCreateResult struct {
	Code    *int64
	Result  *int64
	Message *string
}

// TokenDeleteResult is the result type of the tokenManager service tokenDelete
// method.
type TokenDeleteResult struct {
	Code    *int64
	Result  *int64
	Message *string
}

// TokenItem is the payload type of the tokenManager service tokenCreate method.
type TokenItem struct {
	ID      *string
	TokenID *string
	// chain id
	ChainID int64
	// token address
	Address string
	// name in token contract
	TokenName *string
	// corresponding name in cex
	MarketName string
	Precision  int64
	ChainType  *string
	CoinType   string
}

// TokenListResult is the result type of the tokenManager service tokenList
// method.
type TokenListResult struct {
	Code *int64
	// list
	Result  []*TokenItem
	Message *string
}
