// +build integration
package repository

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"

	"github.com/go-sink/sink/internal/app/datastruct"
)

func TestRepository(t *testing.T) {
	linkRepository := setUpTestLinkRepository(t)

	const origTestValue = "orig"
	const shortTestValue = "short"

	t.Run("it writes a link to a database", func(t *testing.T) { //TODO: delete this
		link := datastruct.Link{Original: origTestValue, Shortened: shortTestValue}

		err := linkRepository.SetLink(link)
		if err != nil {
			t.Fatalf("couldnt write a link to a database: %v", err)
		}
	})

	t.Run("it gets corresponding link", func(t *testing.T) {
		want := datastruct.Link{ID: 5, Original: origTestValue, Shortened: shortTestValue}
		encodedLink := shortTestValue

		got, err := linkRepository.GetLink(encodedLink)
		assert.Nil(t, err)

		// because we don't know ID for sure
		got.ID = 5

		if got != want {
			assert.Equal(t, want, got)
		}
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
