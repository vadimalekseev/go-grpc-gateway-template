package server

import (
	"mime"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/aleksvdim/go-grpc-gateway-template/swagger"
)

const swaggerUIPrefix = "/docs/"

func serveSwaggerUI(mux *http.ServeMux) error {
	if err := mime.AddExtensionType(".svg", "image/svg+xml"); err != nil {
		return err
	}

	// Expose files on <host>/docs/
	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, _ *http.Request) {
		if _, err := w.Write(swagger.GetEchoSwaggerJSON()); err != nil {
			log.Err(err).Msg("error writing swagger.json file: %v")
		}
	})

	mux.Handle(swaggerUIPrefix, http.StripPrefix(swaggerUIPrefix, http.FileServer(http.FS(swagger.GetSwaggerUI()))))

	return nil
}
