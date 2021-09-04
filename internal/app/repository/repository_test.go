package repository_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/hashicorp/hcl/v2/hclsimple"

	_ "github.com/lib/pq"

	datastruct "github.com/go-sink/sink/internal/app/datasctruct"
	"github.com/go-sink/sink/internal/app/repository"
)

type Config struct {
	DB Database `hcl:"database,block"`
}

type Database struct {
	Host string `hcl:"host"`
	Port int `hcl:"port"`
	User string `hcl:"user"`
	Password string `hcl:"password"`
	Dbname string `hcl:"dbname"`
}



func TestRepository(t *testing.T){
	linkRepository := setUpLinkRepository(t)


	t.Run("it writes a link to a database", func(t *testing.T) { //TODO: delete this
		link := datastruct.NewLink("stubborn.fuk", "lmao.no")

		err := linkRepository.SetLink(link)
		if err != nil {
			t.Fatalf("couldnt write a link to a database: %v", err)
		}

	})

	t.Run("it gets corresponding link", func(t *testing.T) {
		want := datastruct.NewLink("stubborn.fuk", "lmao.no")
		encodedLink := "lmao.no"

		got := linkRepository.GetLink(encodedLink)

		if got != want {
			t.Errorf("wrong link pair, got: %v, want: %v", got, want)
		}
	})

}

func setUpLinkRepository(t testing.TB) (linkRepository repository.Repository){
	t.Helper()

	var config Config
	err := hclsimple.DecodeFile("../../../config.hcl", nil, &config)
	if err != nil {
		t.Fatalf("Failed to load configuration: %s", err)
	}
	fmt.Printf("Configuration is %#v", config)

	dbConfig := config.DB

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Dbname)
	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		t.Fatalf("could not establish db connection: %v", err)
	}
	return repository.New(conn)
}
