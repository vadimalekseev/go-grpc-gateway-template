package sinkapi

import (
	"context"

	gw "github.com/go-sink/sink/pkg/api/sink"
)

func (s SinkAPI) Sink(ctx context.Context, req *gw.SinkRequest) (*gw.SinkResponse, error) {
	encoded, err := s.encoder.Encode(ctx, req.Url)
	if err != nil {
		return nil, err
	}

	return &gw.SinkResponse{Url: encoded}, nil
}

func (s SinkAPI) Unsink(ctx context.Context, request *gw.UnsinkRequest) (*gw.UnsinkResponse, error) {
	panic("implement me")
}
