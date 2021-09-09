package echoapi

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"

	echov1 "github.com/aleksvdim/go-grpc-gateway-template/pkg/echo/v1"
)

// Echo returns echo from repository.
func (e EchoAPI) Echo(ctx context.Context, request *echov1.EchoRequest) (*echov1.EchoResponse, error) {
	repoEcho, err := e.repo.Echo(ctx, request.Message)
	if err != nil {
		log.Err(err).Msg("error executing Echo repository method")
		return nil, fmt.Errorf("error getting echo from server: %w", err)
	}

	return &echov1.EchoResponse{Message: repoEcho.Message}, nil
}
