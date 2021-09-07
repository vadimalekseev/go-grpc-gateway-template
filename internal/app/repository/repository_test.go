// +build integration

package repository

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"

	"github.com/aleksvdim/go-grpc-gateway-template/internal/app/datastruct"
)

func TestRepository(t *testing.T) {
	linkRepository := setUpTestLinkRepository(t)

	const echoValue = "hello!"

	t.Run("get echo from db", func(t *testing.T) {
		expected := datastruct.Echo{Message: echoValue}

		echo, err := linkRepository.Echo(context.Background(), echoValue)
		assert.Nil(t, err)
		assert.Equal(t, expected, echo)
	})
}

func setUpTestLinkRepository(t testing.TB) (linkRepository Repository) {
	t.Helper()

	DSN, ok := os.LookupEnv("TEST_DSN")
	if !ok {
		fmt.Println("TEST_DSN environment variable is required")
	}

	conn, err := sql.Open("postgres", DSN)
	if err != nil {
		t.Fatalf("could not establish db connection: %v", err)
	}

	return New(conn)
}
