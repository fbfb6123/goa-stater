// Code generated by goa v3.4.3, DO NOT EDIT.
//
// term_limit HTTP client encoders and decoders
//
// Command:
// $ goa gen goa_starter/design

package client

import (
	"bytes"
	"context"
	"fmt"
	termlimit "goa_starter/gen/term_limit"
	termlimitviews "goa_starter/gen/term_limit/views"
	"io/ioutil"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildGetRequest instantiates a HTTP request object with method and path set
// to call the "term_limit" service "get" endpoint
func (c *Client) BuildGetRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetTermLimitPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("term_limit", "get", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetRequest returns an encoder for requests sent to the term_limit get
// server.
func EncodeGetRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*termlimit.GetPayload)
		if !ok {
			return goahttp.ErrInvalidType("term_limit", "get", "*termlimit.GetPayload", v)
		}
		values := req.URL.Query()
		values.Add("muid", p.Muid)
		values.Add("media_id", fmt.Sprintf("%v", p.MediaID))
		req.URL.RawQuery = values.Encode()
		body := NewGetRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("term_limit", "get", err)
		}
		return nil
	}
}

// DecodeGetResponse returns a decoder for responses returned by the term_limit
// get endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeGetResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("term_limit", "get", err)
			}
			p := NewGetTermLimitResponseOK(&body)
			view := "default"
			vres := &termlimitviews.TermLimitResponse{Projected: p, View: view}
			if err = termlimitviews.ValidateTermLimitResponse(vres); err != nil {
				return nil, goahttp.ErrValidationError("term_limit", "get", err)
			}
			res := termlimit.NewTermLimitResponse(vres)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("term_limit", "get", resp.StatusCode, string(body))
		}
	}
}
