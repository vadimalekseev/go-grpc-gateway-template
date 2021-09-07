package main

import (
	"context"
	_ "embed"
	"flag"
	"os"
	"os/signal"

	"github.com/rs/zerolog/log"

	_ "github.com/lib/pq"

	"github.com/aleksvdim/go-grpc-gateway-template/internal/app/config"
	"github.com/aleksvdim/go-grpc-gateway-template/internal/app/server"
)

var configPath = flag.String("config", "configs/app.example.hcl", "application config")

func main() {
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())

	cfg, err := config.FromFile(*configPath)
	if err != nil {
		log.Fatal().Err(err).Msg("read config file error")
	}

	app, err := server.InitApp(ctx, cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("app init error")
	}

	if err = app.Run(); err != nil {
		log.Fatal().Err(err).Msg("app run error")
	}

	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh)

	sig := <-shutdownCh

	cancel()

	log.Fatal().Msgf("exit reason: %s", sig)
}
