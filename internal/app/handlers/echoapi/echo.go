package echoapi

import (
	"context"
	"fmt"

	"github.com/aleksvdim/go-grpc-gateway-template/pkg/api/echo"
)

func (e EchoAPI) Echo(ctx context.Context, request *echo.EchoRequest) (*echo.EchoResponse, error) {
	repoEcho, err := e.repo.Echo(ctx, request.Message)
	if err != nil {
		return nil, fmt.Errorf("error getting echo from server: %s", err)
	}

	return &echo.EchoResponse{Message: repoEcho.Message}, nil
}
