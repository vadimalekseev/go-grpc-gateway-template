package server

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/http"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/aleksvdim/go-grpc-gateway-template/internal/app/config"
	"github.com/aleksvdim/go-grpc-gateway-template/internal/app/handlers"
	"github.com/aleksvdim/go-grpc-gateway-template/internal/app/handlers/echoapi"
	"github.com/aleksvdim/go-grpc-gateway-template/internal/app/repository"
)

type Server struct {
	httpAddr, grpcAddr string
	grpcServer         *grpc.Server
	httpServer         *http.Server
	registrar          handlers.Registrar
}

func InitApp(ctx context.Context, config config.Config) (*Server, error) {
	s := &Server{
		httpAddr: config.App.HTTPAddr,
		grpcAddr: config.App.GRPCAddr,
	}

	// set up database
	dbcfg := config.Database

	dsn := fmt.Sprintf("user=%s password=%s database=%s sslmode=%s", dbcfg.User, dbcfg.Password, dbcfg.Database, dbcfg.SSLMode)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %s", err)
	}

	repo := repository.New(db)

	echoAPI := echoapi.New(repo)

	s.registrar = handlers.NewRegistrar(echoAPI)

	if err = s.initTransport(ctx); err != nil {
		return nil, fmt.Errorf("error initializing transport: %v", err)
	}

	return s, nil
}

func (s Server) Run() error {
	errWg := errgroup.Group{}

	errWg.Go(func() error {
		lis, err := net.Listen("tcp", s.grpcAddr)
		if err != nil {
			return err
		}

		return s.grpcServer.Serve(lis)
	})

	errWg.Go(func() error {
		l, err := net.Listen("tcp", s.httpAddr)
		if err != nil {
			return err
		}

		return s.httpServer.Serve(l)
	})

	return errWg.Wait()
}
