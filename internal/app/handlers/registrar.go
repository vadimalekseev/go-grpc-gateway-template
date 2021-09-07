package handlers

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"github.com/aleksvdim/go-grpc-gateway-template/pkg/api/echo"
)

type Registrar struct {
	echoServer echo.EchoServer
}

func NewRegistrar(echo echo.EchoServer) Registrar {
	return Registrar{
		echoServer: echo,
	}
}

func (r Registrar) RegisterHandlers(ctx context.Context, grpcServer *grpc.Server, wgMux *runtime.ServeMux) error {
	// register gRPC
	echo.RegisterEchoServer(grpcServer, r.echoServer)
	// register gRPC Gateway
	err := echo.RegisterEchoHandlerServer(ctx, wgMux, r.echoServer)
	if err != nil {
		return err
	}

	return nil
}
