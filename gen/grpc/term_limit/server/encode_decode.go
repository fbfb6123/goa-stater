// Code generated by goa v3.4.3, DO NOT EDIT.
//
// term_limit gRPC server encoders and decoders
//
// Command:
// $ goa gen goa_starter/design

package server

import (
	"context"
	term_limitpb "goa_starter/gen/grpc/term_limit/pb"
	termlimit "goa_starter/gen/term_limit"
	termlimitviews "goa_starter/gen/term_limit/views"

	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc/metadata"
)

// EncodeGetResponse encodes responses from the "term_limit" service "get"
// endpoint.
func EncodeGetResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	vres, ok := v.(*termlimitviews.TermLimitResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("term_limit", "get", "*termlimitviews.TermLimitResponse", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewGetResponse(result)
	return resp, nil
}

// DecodeGetRequest decodes requests sent to "term_limit" service "get"
// endpoint.
func DecodeGetRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *term_limitpb.GetRequest
		ok      bool
	)
	{
		if message, ok = v.(*term_limitpb.GetRequest); !ok {
			return nil, goagrpc.ErrInvalidType("term_limit", "get", "*term_limitpb.GetRequest", v)
		}
	}
	var payload *termlimit.GetPayload
	{
		payload = NewGetPayload(message)
	}
	return payload, nil
}
