package sink

import (
	"embed"
)

// SwaggerUI built files.
//go:embed third_party/swagger-ui
var SwaggerUI embed.FS

// SinkSwaggerJSON embedded sink.swagger.json from sink.proto
//go:embed pkg/api/sink/sink.swagger.json
var SinkSwaggerJSON string
