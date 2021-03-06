// Code generated by goa v3.4.3, DO NOT EDIT.
//
// term_limit gRPC server types
//
// Command:
// $ goa gen goa_starter/design

package server

import (
	term_limitpb "goa_starter/gen/grpc/term_limit/pb"
	termlimit "goa_starter/gen/term_limit"
	termlimitviews "goa_starter/gen/term_limit/views"
)

// NewGetPayload builds the payload of the "get" endpoint of the "term_limit"
// service from the gRPC request type.
func NewGetPayload(message *term_limitpb.GetRequest) *termlimit.GetPayload {
	v := &termlimit.GetPayload{
		ID:      int(message.Id),
		Muid:    message.Muid,
		MediaID: message.MediaId,
	}
	return v
}

// NewGetResponse builds the gRPC response type from the result of the "get"
// endpoint of the "term_limit" service.
func NewGetResponse(result *termlimitviews.TermLimitResponseView) *term_limitpb.GetResponse {
	message := &term_limitpb.GetResponse{}
	if result.ID != nil {
		message.Id = *result.ID
	}
	if result.OfferID != nil {
		message.OfferId = *result.OfferID
	}
	if result.Status != nil {
		message.Status = *result.Status
	}
	return message
}
