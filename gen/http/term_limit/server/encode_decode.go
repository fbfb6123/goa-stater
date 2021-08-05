// Code generated by goa v3.4.3, DO NOT EDIT.
//
// term_limit HTTP server encoders and decoders
//
// Command:
// $ goa gen goa_starter/design

package server

import (
	"context"
	"net/http"
	"strconv"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeAddResponse returns an encoder for responses returned by the
// term_limit add endpoint.
func EncodeAddResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(int)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeAddRequest returns a decoder for requests sent to the term_limit add
// endpoint.
func DecodeAddRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			c2  int
			d   int
			err error

			params = mux.Vars(r)
		)
		{
			c2Raw := params["c"]
			v, err2 := strconv.ParseInt(c2Raw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("c2", c2Raw, "integer"))
			}
			c2 = int(v)
		}
		{
			dRaw := params["d"]
			v, err2 := strconv.ParseInt(dRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("d", dRaw, "integer"))
			}
			d = int(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewAddPayload(c2, d)

		return payload, nil
	}
}