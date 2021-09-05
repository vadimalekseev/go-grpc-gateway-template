package main

import (
	"context"
	"database/sql"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/go-sink/sink/internal/app/handler"
	"github.com/go-sink/sink/internal/app/repository"
	"github.com/go-sink/sink/internal/app/service"
	"github.com/go-sink/sink/internal/pkg/bijection"
)

var httpAddr = flag.String("http-addr", ":8081", "http server addr")

func setUpTempdb(driver string) *sql.DB { //TODO:change it
	DSN, ok := os.LookupEnv("DSN")
	if !ok {
		fmt.Println("DSN environment variable is required")
	}

	conn, err := sql.Open(driver, DSN)
	if err != nil {

	}

	return conn
}

var (
	algorithm = bijection.NewNumberSystemConverter()
	linkRepository = repository.New(setUpTempdb("postgres"))
	encoder = service.NewEncoder(algorithm, &linkRepository, "somedomain")
	server = handler.NewServer(encoder)
)

func run() error {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := http.NewServeMux()
	go func(mux *http.ServeMux) {
		if err := handler.ServeGRPC(mux, server); err != nil {
			log.Fatalln(err)
		}
	}(mux)

	if err := handler.ServeSwaggerUI(mux); err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(*httpAddr, mux)
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
