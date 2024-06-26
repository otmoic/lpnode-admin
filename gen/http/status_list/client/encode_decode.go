// Code generated by goa v3.11.0, DO NOT EDIT.
//
// statusList HTTP client encoders and decoders
//
// Command:
// $ goa gen admin-panel/design

package client

import (
	statuslist "admin-panel/gen/status_list"
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildStatListRequest instantiates a HTTP request object with method and path
// set to call the "statusList" service "statList" endpoint
func (c *Client) BuildStatListRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: StatListStatusListPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("statusList", "statList", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeStatListResponse returns a decoder for responses returned by the
// statusList statList endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeStatListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body StatListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("statusList", "statList", err)
			}
			res := NewStatListResultOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("statusList", "statList", resp.StatusCode, string(body))
		}
	}
}

// unmarshalStatusListItemResponseBodyToStatuslistStatusListItem builds a value
// of type *statuslist.StatusListItem from a value of type
// *StatusListItemResponseBody.
func unmarshalStatusListItemResponseBodyToStatuslistStatusListItem(v *StatusListItemResponseBody) *statuslist.StatusListItem {
	if v == nil {
		return nil
	}
	res := &statuslist.StatusListItem{
		InstallType: v.InstallType,
		StatusKey:   v.StatusKey,
		StatusBody:  v.StatusBody,
		Name:        v.Name,
		ErrMessage:  v.ErrMessage,
	}

	return res
}
