package handlers

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	echov1 "github.com/aleksvdim/go-grpc-gateway-template/pkg/echo/v1"
)

// Registrar is registrar of gRPC and gRPC-Gateway handlers.
type Registrar struct {
	echoServer echov1.EchoServiceServer
}

// NewRegistrar returns new registrar instance.
func NewRegistrar(echo echov1.EchoServiceServer) Registrar {
	return Registrar{
		echoServer: echo,
	}
}

// RegisterHandlers registers all handlers for grpcServer and runtime.ServeMux.
// Moved separately to avoid mixing business logic and infrastructure.
func (r Registrar) RegisterHandlers(ctx context.Context, grpcServer *grpc.Server, wgMux *runtime.ServeMux) error {
	// register gRPC
	echov1.RegisterEchoServiceServer(grpcServer, r.echoServer)
	// register gRPC Gateway
	err := echov1.RegisterEchoServiceHandlerServer(ctx, wgMux, r.echoServer)
	if err != nil {
		return err
	}

	return nil
}
