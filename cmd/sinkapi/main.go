package main

import (
	"context"
	_ "embed"
	"flag"
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

func serveSwaggerUI(mux *http.ServeMux) error {
	if err := mime.AddExtensionType(".svg", "image/svg+xml"); err != nil {
		return err
	}

	// Expose files in third_party/swagger-ui/ on <host>/swagger-ui
	fileServer := http.FileServer(http.FS(sink.SwaggerUI))
	mux.Handle("/third_party/swagger-ui/", fileServer)

	return nil
}

func serveGRPC(ctx context.Context, mux *http.ServeMux) error {
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
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := http.NewServeMux()

	//if err := serveGRPC(ctx, mux); err != nil {
	//	return err
	//}
	if err := serveSwaggerUI(mux); err != nil {
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
