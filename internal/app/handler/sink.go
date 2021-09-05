package handler

import (
	"context"
	"flag"
	"io/fs"
	"log"
	"mime"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/go-sink/sink"
	"github.com/go-sink/sink/internal/app/service"
	gw "github.com/go-sink/sink/pkg/api/sink"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcAddr = flag.String("grpc-addr", ":9090", "gRPC server addr")
)

const swaggerUIPrefix = "/docs/"

type server struct {
	gw.UnimplementedSinkServer
	encoder service.URLEncoder
}

func NewServer(encoder service.URLEncoder) *server {
	return &server{
		UnimplementedSinkServer: gw.UnimplementedSinkServer{},
		encoder:                 encoder,
	}
}

func (s *server) Sink(ctx context.Context ,req *gw.SinkRequest) (*gw.SinkResponse, error) {
	encoded, err := s.encoder.Encode(ctx, req.Url)
	if err != nil {
		return nil, err
	}

	return &gw.SinkResponse{Url: encoded}, nil
}

func ServeSwaggerUI(mux *http.ServeMux) error {
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



func ServeGRPC(mux *http.ServeMux, server *server) error {
	lis, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	serveMux := runtime.NewServeMux()
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	gw.RegisterSinkServer(grpcServer, server)
	err = gw.RegisterSinkHandlerFromEndpoint(context.TODO(), serveMux, *grpcAddr, []grpc.DialOption{grpc.WithInsecure()})
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