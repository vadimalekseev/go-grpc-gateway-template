package repository

import (
	"database/sql"
	"fmt"

	"github.com/go-sink/sink/internal/app/datasctruct"
)

type Repository struct {
	database *sql.DB
}

func New(db *sql.DB) Repository {
	return Repository{database: db}
}

func (r *Repository) GetLink(short string) (original datastruct.Link) {
	originalRow := r.database.QueryRow("SELECT * from links where shortened ==  ?", short)
	err := originalRow.Scan(&original)
	if err != nil {
		fmt.Printf("corresponding link was not found: %v", err)
	}

	return
}

func (r *Repository) SetLink(link datastruct.Link) (err error) {
	_, err = r.database.Query("INSERT INTO links VALUES (?, ?)", link.Original, link.Shortened)
	return
}
