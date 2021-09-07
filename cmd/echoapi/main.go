package main

import (
	"context"
	_ "embed"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/signal"

	_ "github.com/lib/pq"

	"github.com/aleksvdim/go-grpc-gateway-template/internal/app/config"
	"github.com/aleksvdim/go-grpc-gateway-template/internal/app/server"
)

var configPath = flag.String("config", "configs/app.example.hcl", "application config")

func main() {
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())

	cfgBytes, err := ioutil.ReadFile(*configPath)
	if err != nil {
		log.Fatalln(err)
	}

	cfg, err := config.Parse(cfgBytes)
	if err != nil {
		log.Fatalln(err)
	}

	app, err := server.InitApp(ctx, cfg)
	if err != nil {
		log.Fatalln(err)
	}

	if err = app.Run(); err != nil {
		log.Fatalln(err)
	}

	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh)

	sig := <-shutdownCh

	cancel()

	log.Fatalf("exit reason: %s\n", sig)
}
