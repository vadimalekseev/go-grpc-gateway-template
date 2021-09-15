package swagger

import (
	"embed"
	"io/fs"
)

// swaggerUI built files generating with `make generate` or `make download-swagger`.
//go:embed swagger-ui
var swaggerUI embed.FS

const swaggerUISubPath = "swagger-ui"

//go:embed echo/v1/echo.swagger.json
var echoSwaggerJSON []byte

// GetSwaggerUI returns a file system in which the Swagger UI is embedded.
func GetSwaggerUI() fs.FS {
	swaggerFS, err := fs.Sub(swaggerUI, swaggerUISubPath)
	if err != nil {
		panic(err) // the application won't compile without go:embed folder
	}
	return swaggerFS
}

// GetEchoSwaggerJSON returns swagger.json byte slice of Echo API service.
func GetEchoSwaggerJSON() []byte {
	return echoSwaggerJSON
}
