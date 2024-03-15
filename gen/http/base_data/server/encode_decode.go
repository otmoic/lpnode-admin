// Code generated by goa v3.11.0, DO NOT EDIT.
//
// baseData HTTP server encoders and decoders
//
// Command:
// $ goa gen admin-panel/design

package server

import (
	basedata "admin-panel/gen/base_data"
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
)

// EncodeChainDataListResponse returns an encoder for responses returned by the
// baseData chainDataList endpoint.
func EncodeChainDataListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*basedata.ChainDataListResult)
		enc := encoder(ctx, w)
		body := NewChainDataListResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeRunTimeEnvResponse returns an encoder for responses returned by the
// baseData runTimeEnv endpoint.
func EncodeRunTimeEnvResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*basedata.RunTimeEnvResult)
		enc := encoder(ctx, w)
		body := NewRunTimeEnvResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// marshalBasedataChainDataItemToChainDataItemResponseBody builds a value of
// type *ChainDataItemResponseBody from a value of type *basedata.ChainDataItem.
func marshalBasedataChainDataItemToChainDataItemResponseBody(v *basedata.ChainDataItem) *ChainDataItemResponseBody {
	if v == nil {
		return nil
	}
	res := &ChainDataItemResponseBody{
		ID:        v.ID,
		ChainID:   v.ChainID,
		Name:      v.Name,
		ChainName: v.ChainName,
		TokenName: v.TokenName,
	}

	return res
}
