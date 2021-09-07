package swagger

import (
	"embed"
	"io/fs"
)

// swaggerUI built files generating with `make generate` or `make download-swagger`.
//go:embed swagger-ui
var swaggerUI embed.FS

//go:embed echo/v1/echo.swagger.json
var echoSwaggerJSON []byte

// GetSwaggerUI returns a file system in which the Swagger UI is embedded.
func GetSwaggerUI() fs.FS {
	return swaggerUI
}

// GetEchoSwaggerJSON returns swagger.json byte slice of Echo API service.
func GetEchoSwaggerJSON() []byte {
	return echoSwaggerJSON
}
