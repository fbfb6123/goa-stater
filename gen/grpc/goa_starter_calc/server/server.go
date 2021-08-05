// Code generated by goa v3.4.3, DO NOT EDIT.
//
// goa_starter-calc gRPC server
//
// Command:
// $ goa gen goa_starter/design

package server

import (
	"context"
	goastartercalc "goa_starter/gen/goa_starter_calc"
	goa_starter_calcpb "goa_starter/gen/grpc/goa_starter_calc/pb"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
)

// Server implements the goa_starter_calcpb.GoaStarterCalcServer interface.
type Server struct {
	AddH goagrpc.UnaryHandler
	goa_starter_calcpb.UnimplementedGoaStarterCalcServer
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the expr.
type ErrorNamer interface {
	ErrorName() string
}

// New instantiates the server struct with the goa_starter-calc service
// endpoints.
func New(e *goastartercalc.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		AddH: NewAddHandler(e.Add, uh),
	}
}

// NewAddHandler creates a gRPC handler which serves the "goa_starter-calc"
// service "add" endpoint.
func NewAddHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeAddRequest, EncodeAddResponse)
	}
	return h
}

// Add implements the "Add" method in goa_starter_calcpb.GoaStarterCalcServer
// interface.
func (s *Server) Add(ctx context.Context, message *goa_starter_calcpb.AddRequest) (*goa_starter_calcpb.AddResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "add")
	ctx = context.WithValue(ctx, goa.ServiceKey, "goa_starter-calc")
	resp, err := s.AddH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*goa_starter_calcpb.AddResponse), nil
}
