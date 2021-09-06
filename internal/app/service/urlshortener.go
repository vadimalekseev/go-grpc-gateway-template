package service

import (
	"context"
	"fmt"

	"github.com/go-sink/sink/internal/app/datastruct"
)

type Repository interface {
	SetLink(ctx context.Context, link datastruct.Link) error
	GetLink(ctx context.Context, shortenedLink string) (datastruct.Link, error)
}

type EncodingAlgorithm interface {
	Encode(rune) string
	Decode(string) rune
}

type encoder struct {
	algorithm EncodingAlgorithm
	repository Repository
	domain string
}

func NewEncoder(algorithm EncodingAlgorithm, repository Repository, domain string) *encoder {
	return &encoder{
		algorithm:  algorithm,
		repository: repository,
		domain:     domain,
	}
}

func (e *encoder) Encode(ctx context.Context, link string) (string, error) {
	var encodedLink string
	for _, char := range link {
		encodedLink = encodedLink + e.algorithm.Encode(char)
	}
	encodedLink = e.domain + encodedLink

	newLink := datastruct.Link{
		Original:  link,
		Shortened: encodedLink,
	}

	if err := e.repository.SetLink(ctx, newLink); err != nil {
		return encodedLink, fmt.Errorf("error encoding link while writing to repository: %v", err)
	}

	return encodedLink, nil
}

func (e *encoder) Decode(ctx context.Context, link string) (err error) {
	return
}