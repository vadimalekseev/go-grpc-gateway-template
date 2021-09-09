package server

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/grpc"

	"github.com/aleksvdim/go-grpc-gateway-template/internal/app/config"
	"github.com/aleksvdim/go-grpc-gateway-template/internal/app/handlers"
	"github.com/aleksvdim/go-grpc-gateway-template/internal/app/handlers/echoapi"
	"github.com/aleksvdim/go-grpc-gateway-template/internal/app/repository"
)

// Server contains application dependencies.
type Server struct {
	httpAddr, grpcAddr string
	grpcServer         *grpc.Server
	httpServer         *http.Server
	registrar          handlers.Registrar
}

// InitApp initializes handlers and transport.
func InitApp(ctx context.Context, config config.Config) (*Server, error) {
	s := &Server{
		httpAddr: config.App.HTTPAddr,
		grpcAddr: config.App.GRPCAddr,
	}

	db, err := setUpDb(config.Database)
	if err != nil {
		return nil, fmt.Errorf("error set up database: %s", err)
	}

	repo := repository.New(db)
	echoAPI := echoapi.New(repo)
	s.registrar = handlers.NewRegistrar(echoAPI)

	if err = s.initTransport(ctx); err != nil {
		return nil, fmt.Errorf("error initializing transport: %w", err)
	}

	return s, nil
}
