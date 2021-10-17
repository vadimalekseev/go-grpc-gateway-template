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
}

// New returns a new Server.
func New(ctx context.Context, cfg config.Config) (*Server, error) {
	s := &Server{
		httpAddr: cfg.App.HTTPAddr,
		grpcAddr: cfg.App.GRPCAddr,
	}

	db, err := setUpDb(cfg.Database)
	if err != nil {
		return nil, fmt.Errorf("set up database: %s", err)
	}

	repo := repository.New(db)
	echoAPI := echoapi.New(repo)

	registrar := handlers.NewRegistrar(echoAPI)

	err = s.init(ctx, cfg.App, registrar)
	if err != nil {
		return nil, fmt.Errorf("init server: %s", err)
	}

	return s, err
}
