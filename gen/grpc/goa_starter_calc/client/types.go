// Code generated by goa v3.4.3, DO NOT EDIT.
//
// goa_starter-calc gRPC client types
//
// Command:
// $ goa gen goa_starter/design

package client

import (
	goastartercalc "goa_starter/gen/goa_starter_calc"
	goa_starter_calcpb "goa_starter/gen/grpc/goa_starter_calc/pb"
)

// NewAddRequest builds the gRPC request type from the payload of the "add"
// endpoint of the "goa_starter-calc" service.
func NewAddRequest(payload *goastartercalc.AddPayload) *goa_starter_calcpb.AddRequest {
	message := &goa_starter_calcpb.AddRequest{
		C: int32(payload.C),
		D: int32(payload.D),
	}
	return message
}

// NewAddResult builds the result type of the "add" endpoint of the
// "goa_starter-calc" service from the gRPC response type.
func NewAddResult(message *goa_starter_calcpb.AddResponse) int {
	result := int(message.Field)
	return result
}
