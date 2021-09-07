package echoapi

import (
	"context"
	"fmt"

	echov1 "github.com/aleksvdim/go-grpc-gateway-template/pkg/echo/v1"
)

// Echo returns echo from repository.
func (e EchoAPI) Echo(ctx context.Context, request *echov1.EchoRequest) (*echov1.EchoResponse, error) {
	repoEcho, err := e.repo.Echo(ctx, request.Message)
	if err != nil {
		return nil, fmt.Errorf("error getting echo from server: %s", err)
	}

	return &echov1.EchoResponse{Message: repoEcho.Message}, nil
}
