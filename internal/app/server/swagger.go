package server

import (
	"io/fs"
	"log"
	"mime"
	"net/http"

	"github.com/go-sink/sink"
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
