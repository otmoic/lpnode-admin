// Code generated by goa v3.11.0, DO NOT EDIT.
//
// settings HTTP client encoders and decoders
//
// Command:
// $ goa gen admin-panel/design

package client

import (
	settings "admin-panel/gen/settings"
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildSettingsRequest instantiates a HTTP request object with method and path
// set to call the "settings" service "settings" endpoint
func (c *Client) BuildSettingsRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: SettingsSettingsPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("settings", "settings", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeSettingsRequest returns an encoder for requests sent to the settings
// settings server.
func EncodeSettingsRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*settings.SettingsPayload)
		if !ok {
			return goahttp.ErrInvalidType("settings", "settings", "*settings.SettingsPayload", v)
		}
		body := NewSettingsRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("settings", "settings", err)
		}
		return nil
	}
}

// DecodeSettingsResponse returns a decoder for responses returned by the
// settings settings endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeSettingsResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body SettingsResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("settings", "settings", err)
			}
			res := NewSettingsResultOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("settings", "settings", resp.StatusCode, string(body))
		}
	}
}