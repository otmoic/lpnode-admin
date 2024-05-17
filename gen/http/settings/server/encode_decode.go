// Code generated by goa v3.11.0, DO NOT EDIT.
//
// settings HTTP server encoders and decoders
//
// Command:
// $ goa gen admin-panel/design

package server

import (
	settings "admin-panel/gen/settings"
	"context"
	"io"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeSettingsResponse returns an encoder for responses returned by the
// settings settings endpoint.
func EncodeSettingsResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*settings.SettingsResult)
		enc := encoder(ctx, w)
		body := NewSettingsResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeSettingsRequest returns a decoder for requests sent to the settings
// settings endpoint.
func DecodeSettingsRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body SettingsRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateSettingsRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewSettingsPayload(&body)

		return payload, nil
	}
}