// Code generated by goa v3.11.0, DO NOT EDIT.
//
// chainConfig HTTP client encoders and decoders
//
// Command:
// $ goa gen admin-panel/design

package client

import (
	chainconfig "admin-panel/gen/chain_config"
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildSetChainListRequest instantiates a HTTP request object with method and
// path set to call the "chainConfig" service "setChainList" endpoint
func (c *Client) BuildSetChainListRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: SetChainListChainConfigPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("chainConfig", "setChainList", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeSetChainListRequest returns an encoder for requests sent to the
// chainConfig setChainList server.
func EncodeSetChainListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*chainconfig.SetChainListPayload)
		if !ok {
			return goahttp.ErrInvalidType("chainConfig", "setChainList", "*chainconfig.SetChainListPayload", v)
		}
		body := NewSetChainListRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("chainConfig", "setChainList", err)
		}
		return nil
	}
}

// DecodeSetChainListResponse returns a decoder for responses returned by the
// chainConfig setChainList endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeSetChainListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body SetChainListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chainConfig", "setChainList", err)
			}
			res := NewSetChainListResultOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("chainConfig", "setChainList", resp.StatusCode, string(body))
		}
	}
}

// BuildDelChainListRequest instantiates a HTTP request object with method and
// path set to call the "chainConfig" service "delChainList" endpoint
func (c *Client) BuildDelChainListRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DelChainListChainConfigPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("chainConfig", "delChainList", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDelChainListRequest returns an encoder for requests sent to the
// chainConfig delChainList server.
func EncodeDelChainListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*chainconfig.DelChainListPayload)
		if !ok {
			return goahttp.ErrInvalidType("chainConfig", "delChainList", "*chainconfig.DelChainListPayload", v)
		}
		body := NewDelChainListRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("chainConfig", "delChainList", err)
		}
		return nil
	}
}

// DecodeDelChainListResponse returns a decoder for responses returned by the
// chainConfig delChainList endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeDelChainListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body DelChainListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chainConfig", "delChainList", err)
			}
			res := NewDelChainListResultOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("chainConfig", "delChainList", resp.StatusCode, string(body))
		}
	}
}

// BuildChainListRequest instantiates a HTTP request object with method and
// path set to call the "chainConfig" service "chainList" endpoint
func (c *Client) BuildChainListRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ChainListChainConfigPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("chainConfig", "chainList", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeChainListResponse returns a decoder for responses returned by the
// chainConfig chainList endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeChainListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body ChainListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chainConfig", "chainList", err)
			}
			res := NewChainListResultOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("chainConfig", "chainList", resp.StatusCode, string(body))
		}
	}
}

// BuildSetChainGasUsdRequest instantiates a HTTP request object with method
// and path set to call the "chainConfig" service "setChainGasUsd" endpoint
func (c *Client) BuildSetChainGasUsdRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: SetChainGasUsdChainConfigPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("chainConfig", "setChainGasUsd", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeSetChainGasUsdRequest returns an encoder for requests sent to the
// chainConfig setChainGasUsd server.
func EncodeSetChainGasUsdRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*chainconfig.SetChainGasUsdPayload)
		if !ok {
			return goahttp.ErrInvalidType("chainConfig", "setChainGasUsd", "*chainconfig.SetChainGasUsdPayload", v)
		}
		body := NewSetChainGasUsdRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("chainConfig", "setChainGasUsd", err)
		}
		return nil
	}
}

// DecodeSetChainGasUsdResponse returns a decoder for responses returned by the
// chainConfig setChainGasUsd endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeSetChainGasUsdResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body SetChainGasUsdResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chainConfig", "setChainGasUsd", err)
			}
			res := NewSetChainGasUsdResultOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("chainConfig", "setChainGasUsd", resp.StatusCode, string(body))
		}
	}
}

// BuildSetChainClientConfigRequest instantiates a HTTP request object with
// method and path set to call the "chainConfig" service "setChainClientConfig"
// endpoint
func (c *Client) BuildSetChainClientConfigRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: SetChainClientConfigChainConfigPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("chainConfig", "setChainClientConfig", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeSetChainClientConfigRequest returns an encoder for requests sent to
// the chainConfig setChainClientConfig server.
func EncodeSetChainClientConfigRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*chainconfig.SetChainClientConfigPayload)
		if !ok {
			return goahttp.ErrInvalidType("chainConfig", "setChainClientConfig", "*chainconfig.SetChainClientConfigPayload", v)
		}
		body := NewSetChainClientConfigRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("chainConfig", "setChainClientConfig", err)
		}
		return nil
	}
}

// DecodeSetChainClientConfigResponse returns a decoder for responses returned
// by the chainConfig setChainClientConfig endpoint. restoreBody controls
// whether the response body should be restored after having been read.
func DecodeSetChainClientConfigResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body SetChainClientConfigResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chainConfig", "setChainClientConfig", err)
			}
			res := NewSetChainClientConfigResultOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("chainConfig", "setChainClientConfig", resp.StatusCode, string(body))
		}
	}
}

// marshalChainconfigChainDataItemToChainDataItemRequestBody builds a value of
// type *ChainDataItemRequestBody from a value of type
// *chainconfig.ChainDataItem.
func marshalChainconfigChainDataItemToChainDataItemRequestBody(v *chainconfig.ChainDataItem) *ChainDataItemRequestBody {
	if v == nil {
		return nil
	}
	res := &ChainDataItemRequestBody{
		ChainID:   v.ChainID,
		ChainName: v.ChainName,
		Name:      v.Name,
		TokenName: v.TokenName,
	}

	return res
}

// marshalChainDataItemRequestBodyToChainconfigChainDataItem builds a value of
// type *chainconfig.ChainDataItem from a value of type
// *ChainDataItemRequestBody.
func marshalChainDataItemRequestBodyToChainconfigChainDataItem(v *ChainDataItemRequestBody) *chainconfig.ChainDataItem {
	if v == nil {
		return nil
	}
	res := &chainconfig.ChainDataItem{
		ChainID:   v.ChainID,
		ChainName: v.ChainName,
		Name:      v.Name,
		TokenName: v.TokenName,
	}

	return res
}

// unmarshalChainDataItemResponseBodyToChainconfigChainDataItem builds a value
// of type *chainconfig.ChainDataItem from a value of type
// *ChainDataItemResponseBody.
func unmarshalChainDataItemResponseBodyToChainconfigChainDataItem(v *ChainDataItemResponseBody) *chainconfig.ChainDataItem {
	if v == nil {
		return nil
	}
	res := &chainconfig.ChainDataItem{
		ChainID:   v.ChainID,
		ChainName: v.ChainName,
		Name:      v.Name,
		TokenName: v.TokenName,
	}

	return res
}