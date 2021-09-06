package sinkapi

import (
	"context"

	gw "github.com/go-sink/sink/pkg/api/sink"
)

type URLEncoder interface {
	Encode(ctx context.Context, link string) (encodedLink string, err error)
	Decode(ctx context.Context, shortened string) (err error)
}

type SinkAPI struct {
	gw.UnsafeSinkServer

	encoder URLEncoder
}

func New(encoder URLEncoder) SinkAPI {
	return SinkAPI{
		encoder: encoder,
	}
}
