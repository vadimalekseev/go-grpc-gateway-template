package server

import (
	"mime"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/aleksvdim/go-grpc-gateway-template/swagger"
)

const swaggerUIPrefix = "/docs/"
const swaggerJSONPath = "/docs/swagger.json"

func serveSwaggerUI(mux *http.ServeMux) error {
	if err := mime.AddExtensionType(".svg", "image/svg+xml"); err != nil {
		return err
	}

	echoAPISwagger := swagger.GetEchoSwaggerJSON()
	// Expose files on <host>/docs/
	mux.HandleFunc(swaggerJSONPath, func(w http.ResponseWriter, _ *http.Request) {
		if _, err := w.Write(echoAPISwagger); err != nil {
			log.Err(err).Msg("write swagger.json file")
		}
	})

	mux.Handle(swaggerUIPrefix, http.StripPrefix(swaggerUIPrefix, http.FileServer(http.FS(swagger.GetSwaggerUI()))))

	return nil
}
