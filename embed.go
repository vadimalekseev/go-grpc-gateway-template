package template

import (
	"embed"
)

// SwaggerUIPath is the path to swagger-ui built files in embed.FS.
const SwaggerUIPath = "third_party/swagger-ui"

// ThirdParty built files.
//go:embed third_party
var ThirdParty embed.FS

// EchoAPISwaggerJSON embedded echo.swagger.json from echo.proto
//go:embed pkg/api/echo/echo.swagger.json
var EchoAPISwaggerJSON []byte
