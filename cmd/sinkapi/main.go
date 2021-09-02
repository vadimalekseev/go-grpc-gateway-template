package main

import (
	"context"
	_ "embed"
	"flag"
	"io/fs"
	"log"
	"mime"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"github.com/go-sink/sink"
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

type sinktest struct {
	gw.UnimplementedSinkServer
}

func (s *sinktest) Sink(_ context.Context, req *gw.SinkRequest) (*gw.SinkResponse, error){
	return &gw.SinkResponse{Url: req.Url}, nil
}

func serveGRPC(mux *http.ServeMux) error {
	lis, err := net.Listen("tcp", ":5555")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	serveMux := runtime.NewServeMux()
	grpcServer := grpc.NewServer()
	gw.RegisterSinkServer(grpcServer, &sinktest{})
	err = gw.RegisterSinkHandlerFromEndpoint(context.TODO(), serveMux, ":5555", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		return err
	}
	mux.Handle("/", serveMux)

	err = grpcServer.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}

func run() error {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := http.NewServeMux()
	go func(mux *http.ServeMux) {
		if err := serveGRPC(mux); err != nil {
			log.Fatalln(err)
		}
	}(mux)

	if err := serveSwaggerUI(mux); err != nil {
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
