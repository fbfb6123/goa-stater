// Code generated by goa v3.4.3, DO NOT EDIT.
//
// term_limit gRPC client
//
// Command:
// $ goa gen goa_starter/design

package client

import (
	"context"
	term_limitpb "goa_starter/gen/grpc/term_limit/pb"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli term_limitpb.TermLimitClient
	opts    []grpc.CallOption
}

// NewClient instantiates gRPC client for all the term_limit service servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: term_limitpb.NewTermLimitClient(cc),
		opts:    opts,
	}
}

// Add calls the "Add" function in term_limitpb.TermLimitClient interface.
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