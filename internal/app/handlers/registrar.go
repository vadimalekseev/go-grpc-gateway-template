package handlers

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	pb "github.com/go-sink/sink/pkg/api/sink"
)

type Registrar struct {
	sinkServer pb.SinkServer
}

func NewRegistrar(sink pb.SinkServer) Registrar {
	return Registrar{
		sinkServer: sink,
	}
}

func (r Registrar) RegisterHandlers(ctx context.Context, grpcServer *grpc.Server, wgMux *runtime.ServeMux) error {
	pb.RegisterSinkServer(grpcServer, r.sinkServer)
	err := pb.RegisterSinkHandlerServer(ctx, wgMux, r.sinkServer)
	if err != nil {
		return err
	}

	return nil
}
