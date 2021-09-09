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

var configPath = flag.String("config", "configs/app.example.hcl", "application config")

func main() {
	flag.Parse()
	ctx := context.Background()

	cfg, err := config.FromFile(*configPath)
	if err != nil {
		log.Fatal().Err(err).Msg("read config file error")
	}

	app, err := server.InitApp(ctx, cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("app init error")
	}

	if err = app.Run(ctx); err != nil {
		log.Fatal().Err(err).Msg("app run error")
	}
}
