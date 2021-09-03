package repository_test

import (
	"github.com/go-sink/sink/internal/app/repository"
	"database/sql"
	"database/sql/driver"
	"testing"
)


func TestRepository(t *testing.T){
	t.Run("it gets corresponding link", func(t *testing.T) {
		want := "somelink.dom"
		encodedLink := []int{1,2,3,45,6,7,8,8}
		sql.Open("postgres")
		linkRepository := repository.New()
		got := linkRepository.Get(encodedLink)

		if got != want {
			t.Errorf("wrong link, got: %v, want: %v", got, want)
		}
	})
}
