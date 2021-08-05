package main

import (
	"context"
	goastarter "goa_starter/gen/goa_starter"
	goa_starterpb "goa_starter/gen/grpc/goa_starter/pb"
	goastartersvr "goa_starter/gen/grpc/goa_starter/server"
	term_limitpb "goa_starter/gen/grpc/term_limit/pb"
	termlimitsvr "goa_starter/gen/grpc/term_limit/server"
	log "goa_starter/gen/log"
	termlimit "goa_starter/gen/term_limit"
	"net"
	"net/url"
	"sync"

	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcmdlwr "goa.design/goa/v3/grpc/middleware"
	"goa.design/goa/v3/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// handleGRPCServer starts configures and starts a gRPC server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleGRPCServer(ctx context.Context, u *url.URL, goaStarterEndpoints *goastarter.Endpoints, termLimitEndpoints *termlimit.Endpoints, wg *sync.WaitGroup, errc chan error, logger *log.Logger, debug bool) {

	// Setup goa log adapter.
	var (
		adapter middleware.Logger
	)
	{
		adapter = logger
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to gRPC requests and
	// responses.
	var (
		goaStarterServer *goastartersvr.Server
		termLimitServer  *termlimitsvr.Server
	)
	{
		goaStarterServer = goastartersvr.New(goaStarterEndpoints, nil)
		termLimitServer = termlimitsvr.New(termLimitEndpoints, nil)
	}

	// Initialize gRPC server with the middleware.
	srv := grpc.NewServer(
		grpcmiddleware.WithUnaryServerChain(
			grpcmdlwr.UnaryRequestID(),
			grpcmdlwr.UnaryServerLog(adapter),
		),
	)

	// Register the servers.
	goa_starterpb.RegisterGoaStarterServer(srv, goaStarterServer)
	term_limitpb.RegisterTermLimitServer(srv, termLimitServer)

	for svc, info := range srv.GetServiceInfo() {
		for _, m := range info.Methods {
			logger.Infof("serving gRPC method %s", svc+"/"+m.Name)
		}
	}

	// Register the server reflection service on the server.
	// See https://grpc.github.io/grpc/core/md_doc_server-reflection.html.
	reflection.Register(srv)

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start gRPC server in a separate goroutine.
		go func() {
			lis, err := net.Listen("tcp", u.Host)
			if err != nil {
				errc <- err
			}
			logger.Infof("gRPC server listening on %q", u.Host)
			errc <- srv.Serve(lis)
		}()

		<-ctx.Done()
		logger.Infof("shutting down gRPC server at %q", u.Host)
		srv.Stop()
	}()
}
