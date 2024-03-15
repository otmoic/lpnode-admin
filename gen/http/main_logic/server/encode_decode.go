// Code generated by goa v3.11.0, DO NOT EDIT.
//
// mainLogic HTTP server encoders and decoders
//
// Command:
// $ goa gen admin-panel/design

package server

import (
	mainlogic "admin-panel/gen/main_logic"
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
)

// EncodeMainLogicResponse returns an encoder for responses returned by the
// mainLogic mainLogic endpoint.
func EncodeMainLogicResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// EncodeMainLogicLinkResponse returns an encoder for responses returned by the
// mainLogic mainLogicLink endpoint.
func EncodeMainLogicLinkResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*mainlogic.MainLogicLinkResult)
		enc := encoder(ctx, w)
		body := NewMainLogicLinkResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}
