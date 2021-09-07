package echoapi

import (
	"context"

	"github.com/aleksvdim/go-grpc-gateway-template/internal/app/datastruct"
)

// Repo contains required repository methods for Echo API.
type Repo interface {
	Echo(ctx context.Context, message string) (datastruct.Echo, error)
}

// EchoAPI contains Echo API dependencies.
type EchoAPI struct {
	repo Repo
}

// New returns new EchoAPI instance.
func New(repo Repo) EchoAPI {
	return EchoAPI{
		repo: repo,
	}
}
