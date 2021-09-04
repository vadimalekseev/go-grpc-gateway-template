package repository_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"

	"github.com/go-sink/sink/internal/app/datastruct"
	"github.com/go-sink/sink/internal/app/repository"
)

func TestRepository(t *testing.T) {
	linkRepository := setUpLinkRepository(t)

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
		want := datastruct.Link{Original: origTestValue, Shortened: shortTestValue}
		encodedLink := shortTestValue

		got := linkRepository.GetLink(encodedLink)

		if got != want {
			assert.Equal(t, got, want)
			//t.Errorf("wrong link pair, got: %v, want: %v", got, want)
		}
	})

}

func setUpLinkRepository(t testing.TB) (linkRepository repository.Repository) {
	t.Helper()

	DSN, ok := os.LookupEnv("TEST_DSN")
	if !ok {
		fmt.Println("TEST_DSN environment variable is required")
	}

	conn, err := sql.Open("postgres", DSN)
	if err != nil {
		t.Fatalf("could not establish db connection: %v", err)
	}

	return repository.New(conn)
}
