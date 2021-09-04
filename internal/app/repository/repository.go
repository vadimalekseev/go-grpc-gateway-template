package repository

import (
	"database/sql"

	"github.com/go-sink/sink/internal/app/datastruct"
)

type Repository struct {
	database *sql.DB
}

func New(db *sql.DB) Repository {
	return Repository{database: db}
}

func (r *Repository) GetLink(short string) (original datastruct.Link, err error) {
	originalRow := r.database.QueryRow("SELECT id, original, shortened from links where shortened = $1", short)
	err = originalRow.Scan(&original.ID, &original.Original, &original.Shortened)
	return
}

func (r *Repository) SetLink(link datastruct.Link) (err error) {
	_, err = r.database.Query("INSERT INTO links(original, shortened) VALUES ($1, $2)", link.Original, link.Shortened)
	return
}
