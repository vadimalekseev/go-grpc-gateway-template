package main

import (
	"context"
	_ "embed"
	"flag"
	"io/fs"
	"log"
	"mime"
	"net/http"

	"github.com/go-sink/sink"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	gw "github.com/go-sink/sink/pkg/api/sink"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcAddr = flag.String("grpc-addr", ":9090", "gRPC server addr")
	httpAddr = flag.String("http-addr", ":8081", "http server addr")
)

const swaggerUIPrefix = "/docs/"

func enableSwaggerSupport(mux *http.ServeMux) error {
	if err := mime.AddExtensionType(".svg", "image/svg+xml"); err != nil {
		return err
	}

	// Expose files on <host>/docs
	swaggerUIFS, err := fs.Sub(sink.SwaggerUI, sink.SwaggerUIPath)
	if err != nil {
		return err
	}

	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, _ *http.Request) {
		if _, err = w.Write(sink.SinkSwaggerJSON); err != nil {
			log.Printf("error writing swagger.json file: %v", err)
		}
	})
	mux.Handle(swaggerUIPrefix, http.StripPrefix(swaggerUIPrefix, http.FileServer(http.FS(swaggerUIFS))))

	return nil
}

func enableGRPCSupport(ctx context.Context, mux *http.ServeMux) error {
	// TODO: Make sure the gRPC server is running properly and accessible
	gatewayMux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterSinkHandlerFromEndpoint(ctx, gatewayMux, *grpcAddr, opts)
	if err != nil {
		return err
	}

	mux.Handle("/", gatewayMux)
	return nil
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := http.NewServeMux()

	if err := enableGRPCSupport(ctx, mux); err != nil {
		return err
	}
	if err := enableSwaggerSupport(mux); err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(*httpAddr, mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
