// Code generated by goa v3.4.3, DO NOT EDIT.
//
// term_limit gRPC server
//
// Command:
// $ goa gen goa_starter/design

package server

import (
	"context"
	term_limitpb "goa_starter/gen/grpc/term_limit/pb"
	termlimit "goa_starter/gen/term_limit"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
)

// Server implements the term_limitpb.TermLimitServer interface.
type Server struct {
	GetH goagrpc.UnaryHandler
	term_limitpb.UnimplementedTermLimitServer
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the expr.
type ErrorNamer interface {
	ErrorName() string
}

// New instantiates the server struct with the term_limit service endpoints.
func New(e *termlimit.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		GetH: NewGetHandler(e.Get, uh),
	}
}

// NewGetHandler creates a gRPC handler which serves the "term_limit" service
// "get" endpoint.
func NewGetHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeGetRequest, EncodeGetResponse)
	}
	return h
}

// Get implements the "Get" method in term_limitpb.TermLimitServer interface.
func (s *Server) Get(ctx context.Context, message *term_limitpb.GetRequest) (*term_limitpb.GetResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "get")
	ctx = context.WithValue(ctx, goa.ServiceKey, "term_limit")
	resp, err := s.GetH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*term_limitpb.GetResponse), nil
}
