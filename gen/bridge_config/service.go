// Code generated by goa v3.11.0, DO NOT EDIT.
//
// bridgeConfig service
//
// Command:
// $ goa gen admin-panel/design

package bridgeconfig

import (
	"context"
)

// Service is the bridgeConfig service interface.
type Service interface {
	// 用于创建跨链配置
	BridgeCreate(context.Context, *BridgeItem) (res *BridgeCreateResult, err error)
	// BridgeList implements bridgeList.
	BridgeList(context.Context) (res *BridgeListResult, err error)
	// BridgeDelete implements bridgeDelete.
	BridgeDelete(context.Context, *DeleteBridgeFilter) (res *BridgeDeleteResult, err error)
	// BridgeTest implements bridgeTest.
	BridgeTest(context.Context, *BridgeTestPayload) (res *BridgeTestResult, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "bridgeConfig"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [4]string{"bridgeCreate", "bridgeList", "bridgeDelete", "bridgeTest"}

// BridgeCreateResult is the result type of the bridgeConfig service
// bridgeCreate method.
type BridgeCreateResult struct {
	Code *int64
	// 是否成功
	Result  *int64
	Message *string
}

// BridgeDeleteResult is the result type of the bridgeConfig service
// bridgeDelete method.
type BridgeDeleteResult struct {
	Code *int64
	// 是否删除成功
	Result  *int64
	Message *string
}

// BridgeItem is the payload type of the bridgeConfig service bridgeCreate
// method.
type BridgeItem struct {
	// bridge的Name ****
	BridgeName string
	// mongodb的主键,baseData中获取
	SrcChainID string
	// mongodb的主键,baseData中获取
	DstChainID string
	// mongodb的主键,tokenList中获取
	SrcTokenID string
	// mongodb的主键,tokenList中获取
	DstTokenID string
	// mongodb的主键,walletList 中获取
	WalletID string
	// mongodb的主键,walletList 中获取
	SrcWalletID string
	// amm安装时候的name
	AmmName     string
	EnableHedge bool
}

// BridgeListResult is the result type of the bridgeConfig service bridgeList
// method.
type BridgeListResult struct {
	Code *int64
	// 链的列表
	Result  []*ListBridgeItem
	Message *string
}

// BridgeTestPayload is the payload type of the bridgeConfig service bridgeTest
// method.
type BridgeTestPayload struct {
	ID *string
}

// BridgeTestResult is the result type of the bridgeConfig service bridgeTest
// method.
type BridgeTestResult struct {
	Code *int64
}

// DeleteBridgeFilter is the payload type of the bridgeConfig service
// bridgeDelete method.
type DeleteBridgeFilter struct {
	// Mongodb 的主键
	ID string
}

type ListBridgeItem struct {
	ID                *string
	DstChainID        *string
	DstTokenID        *string
	SrcChainID        *string
	SrcTokenID        *string
	AmmName           *string
	BridgeName        *string
	DstChainRawID     *int64
	DstClientURI      *string
	DstToken          *string
	LpReceiverAddress *string
	MsmqName          *string
	SrcChainRawID     *int64
	SrcToken          *string
	WalletName        *string
	WalletID          *string
	EnableHedge       *bool
}