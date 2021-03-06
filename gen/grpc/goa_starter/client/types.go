// Code generated by goa v3.4.3, DO NOT EDIT.
//
// goa_starter gRPC client types
//
// Command:
// $ goa gen goa_starter/design

package client

import (
	goastarter "goa_starter/gen/goa_starter"
	goa_starterpb "goa_starter/gen/grpc/goa_starter/pb"
)

// NewAddRequest builds the gRPC request type from the payload of the "add"
// endpoint of the "goa_starter" service.
func NewAddRequest(payload *goastarter.AddPayload) *goa_starterpb.AddRequest {
	message := &goa_starterpb.AddRequest{
		A: int32(payload.A),
		B: int32(payload.B),
	}
	return message
}

// NewAddResult builds the result type of the "add" endpoint of the
// "goa_starter" service from the gRPC response type.
func NewAddResult(message *goa_starterpb.AddResponse) int {
	result := int(message.Field)
	return result
}
