// Code generated by goa v3.4.3, DO NOT EDIT.
//
// goa_starter-calc gRPC client
//
// Command:
// $ goa gen goa_starter/design

package client

import (
	"context"
	goa_starter_calcpb "goa_starter/gen/grpc/goa_starter_calc/pb"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli goa_starter_calcpb.GoaStarterCalcClient
	opts    []grpc.CallOption
}

// NewClient instantiates gRPC client for all the goa_starter-calc service
// servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: goa_starter_calcpb.NewGoaStarterCalcClient(cc),
		opts:    opts,
	}
}

// Add calls the "Add" function in goa_starter_calcpb.GoaStarterCalcClient
// interface.
func (c *Client) Add() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildAddFunc(c.grpccli, c.opts...),
			EncodeAddRequest,
			DecodeAddResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}
