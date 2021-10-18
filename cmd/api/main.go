package main

import (
	"context"
	_ "embed"
	"flag"

	"github.com/rs/zerolog/log"

	_ "github.com/lib/pq"

	"github.com/aleksvdim/go-grpc-gateway-template/internal/app/config"
	"github.com/aleksvdim/go-grpc-gateway-template/internal/app/server"
)

const defaultConfig = "configs/app.example.hcl"

var configPath = flag.String("config", defaultConfig, "application config")

func main() {
	flag.Parse()
	ctx := context.Background()

	if *configPath == defaultConfig {
		log.Warn().Msgf(
			"App uses the default config file (%s). To provide your own config use -config flag.",
			defaultConfig,
		)
	}

	cfg, err := config.FromFile(*configPath)
	if err != nil {
		log.Fatal().Err(err).Msg("read config file error")
	}

	app, err := server.New(ctx, cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("app init error")
	}

	if err = app.Run(ctx); err != nil {
		log.Fatal().Err(err).Msg("app run error")
	}
}
