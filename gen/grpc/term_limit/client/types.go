// Code generated by goa v3.4.3, DO NOT EDIT.
//
// term_limit gRPC client types
//
// Command:
// $ goa gen goa_starter/design

package client

import (
	term_limitpb "goa_starter/gen/grpc/term_limit/pb"
	termlimit "goa_starter/gen/term_limit"
)

// NewAddRequest builds the gRPC request type from the payload of the "add"
// endpoint of the "term_limit" service.
func NewAddRequest(payload *termlimit.AddPayload) *term_limitpb.AddRequest {
	message := &term_limitpb.AddRequest{
		C: int32(payload.C),
		D: int32(payload.D),
	}
	return message
}

// NewAddResult builds the result type of the "add" endpoint of the
// "term_limit" service from the gRPC response type.
func NewAddResult(message *term_limitpb.AddResponse) int {
	result := int(message.Field)
	return result
}
