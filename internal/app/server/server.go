package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
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
		return nil, fmt.Errorf("error initializing transport: %v", err)
	}

	return s, nil
}

// Run starts the application.
func (s Server) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
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

	errWg.Go(func() error {
		swaggerAddr := s.httpAddr + swaggerUIPrefix
		log.Info().Msgf("App started. HTTP: %s, Swagger UI: %s, gRPC: %s\n", s.httpAddr, swaggerAddr, s.grpcAddr)
		return nil
	})

	errWg.Go(func() error {
		shutdownCh := make(chan os.Signal, 1)
		signal.Notify(shutdownCh)
		sig := <-shutdownCh

		s.Stop(ctx)
		cancel()

		log.Fatal().Msgf("exit reason: %s", sig)

		return nil
	})

	return errWg.Wait()
}

// Stop the gRPC and HTTP servers.
func (s Server) Stop(ctx context.Context) {
	s.grpcServer.Stop()
	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		log.Err(err).Msg("error shutting down the http server")
	}
}
