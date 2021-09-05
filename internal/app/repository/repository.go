package repository

import (
	"context"
	"database/sql"

	"github.com/go-sink/sink/internal/app/datastruct"
)

type Repository struct {
	database *sql.DB
}

func New(db *sql.DB) Repository {
	return Repository{database: db}
}

func (r *Repository) GetLink(ctx context.Context, short string) (original datastruct.Link, err error) {
	originalRow := r.database.QueryRowContext(ctx, "SELECT id, original, shortened from links where shortened = $1", short)
	err = originalRow.Scan(&original.ID, &original.Original, &original.Shortened)
	return
}

func (r *Repository) SetLink(ctx context.Context, link datastruct.Link) (err error) {
	_, err = r.database.QueryContext(ctx, "INSERT INTO links(original, shortened) VALUES ($1, $2)", link.Original, link.Shortened)
	return
}
