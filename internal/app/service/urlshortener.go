package service

import (
	"context"
	"fmt"

	"github.com/go-sink/sink/internal/app/datastruct"
	"github.com/go-sink/sink/internal/pkg/bijection"
)

type Repository interface {
	SetLink(ctx context.Context, link datastruct.Link) error
	GetLink(ctx context.Context, shortenedLink string) (datastruct.Link, error)
}

type URLEncoder interface {
	Encode(ctx context.Context, link string) (encodedLink string, err error)
	Decode(ctx context.Context, shortened string) (err error)
}

type encoder struct {
	algorithm bijection.EncodingAlgorithm
	repository Repository
	domain string
}

func NewEncoder(algorithm bijection.EncodingAlgorithm, repository Repository, domain string) *encoder {
	return &encoder{
		algorithm:  algorithm,
		repository: repository,
		domain:     domain,
	}
}

func (e *encoder) Encode(ctx context.Context, link string) (encodedLink string, err error) {
	for _, char := range link {
		encodedLink = encodedLink + e.algorithm.Encode(char)
	}
	encodedLink = e.domain + encodedLink

	newLink := datastruct.Link{
		Original:  link,
		Shortened: encodedLink,
	}

	setLinkErr := e.repository.SetLink(ctx, newLink)

	err = fmt.Errorf("problem setting link while encoding: %v", setLinkErr)

	return
}

func (e *encoder) Decode(ctx context.Context, link string) (err error) {
	return
}