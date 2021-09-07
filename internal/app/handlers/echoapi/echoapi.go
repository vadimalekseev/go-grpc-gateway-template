package echoapi

import (
	"context"

	"github.com/aleksvdim/go-grpc-gateway-template/internal/app/datastruct"
	"github.com/aleksvdim/go-grpc-gateway-template/pkg/api/echo"
)

type Repo interface {
	Echo(ctx context.Context, message string) (datastruct.Echo, error)
}

type EchoAPI struct {
	repo Repo

	echo.UnimplementedEchoServer
}

func New(repo Repo) EchoAPI {
	return EchoAPI{
		repo: repo,
	}
}
