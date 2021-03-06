// Code generated by goa v3.4.3, DO NOT EDIT.
//
// goa_starter-calc gRPC server encoders and decoders
//
// Command:
// $ goa gen goa_starter/design

package server

import (
	"context"
	goastartercalc "goa_starter/gen/goa_starter_calc"
	goa_starter_calcpb "goa_starter/gen/grpc/goa_starter_calc/pb"

	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc/metadata"
)

// EncodeAddResponse encodes responses from the "goa_starter-calc" service
// "add" endpoint.
func EncodeAddResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(int)
	if !ok {
		return nil, goagrpc.ErrInvalidType("goa_starter-calc", "add", "int", v)
	}
	resp := NewAddResponse(result)
	return resp, nil
}

// DecodeAddRequest decodes requests sent to "goa_starter-calc" service "add"
// endpoint.
func DecodeAddRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *goa_starter_calcpb.AddRequest
		ok      bool
	)
	{
		if message, ok = v.(*goa_starter_calcpb.AddRequest); !ok {
			return nil, goagrpc.ErrInvalidType("goa_starter-calc", "add", "*goa_starter_calcpb.AddRequest", v)
		}
	}
	var payload *goastartercalc.AddPayload
	{
		payload = NewAddPayload(message)
	}
	return payload, nil
}
