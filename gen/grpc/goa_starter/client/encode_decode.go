// Code generated by goa v3.4.3, DO NOT EDIT.
//
// goa_starter gRPC client encoders and decoders
//
// Command:
// $ goa gen goa_starter/design

package client

import (
	"context"
	goastarter "goa_starter/gen/goa_starter"
	goa_starterpb "goa_starter/gen/grpc/goa_starter/pb"

	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildAddFunc builds the remote method to invoke for "goa_starter" service
// "add" endpoint.
func BuildAddFunc(grpccli goa_starterpb.GoaStarterClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Add(ctx, reqpb.(*goa_starterpb.AddRequest), opts...)
		}
		return grpccli.Add(ctx, &goa_starterpb.AddRequest{}, opts...)
	}
}

// EncodeAddRequest encodes requests sent to goa_starter add endpoint.
func EncodeAddRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*goastarter.AddPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("goa_starter", "add", "*goastarter.AddPayload", v)
	}
	return NewAddRequest(payload), nil
}

// DecodeAddResponse decodes responses from the goa_starter add endpoint.
func DecodeAddResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*goa_starterpb.AddResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("goa_starter", "add", "*goa_starterpb.AddResponse", v)
	}
	res := NewAddResult(message)
	return res, nil
}
