package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/aleksvdim/go-grpc-gateway-template/internal/app/config"
	"github.com/aleksvdim/go-grpc-gateway-template/internal/app/handlers"
)

const timeout = time.Second * 15

func (s *Server) init(ctx context.Context, cfg config.App, registrar handlers.Registrar) (err error) {
	gwMux := runtime.NewServeMux()
	httpMux := http.NewServeMux()

	grpcServer := grpc.NewServer(grpc.ConnectionTimeout(timeout))
	if err = registrar.RegisterHandlers(ctx, grpcServer, gwMux); err != nil {
		return fmt.Errorf("register handlers: %w", err)
	}

	if cfg.UseGRPCReflect {
		reflection.Register(grpcServer)
	}

	httpMux.Handle("/", gwMux)
	httpMux.Handle(cfg.MetricsAddr, promhttp.Handler())
	serveHealth(httpMux)
	if err = serveSwaggerUI(httpMux); err != nil {
		return fmt.Errorf("serve Swagger UI: %w", err)
	}

	s.httpServer = &http.Server{
		Handler:        httpMux,
		ReadTimeout:    timeout,
		WriteTimeout:   timeout,
		MaxHeaderBytes: 2 << 11,
	}
	s.grpcServer = grpcServer

	return nil
}
