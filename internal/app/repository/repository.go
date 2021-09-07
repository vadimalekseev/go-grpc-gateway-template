package repository

import (
	"context"
	"database/sql"

	"github.com/aleksvdim/go-grpc-gateway-template/internal/app/datastruct"
)

// Repository is repository.
type Repository struct {
	database *sql.DB
}

// New returns new repository instance.
func New(db *sql.DB) Repository {
	return Repository{database: db}
}

// Echo returns message from database.
func (r Repository) Echo(ctx context.Context, message string) (echo datastruct.Echo, err error) {
	originalRow := r.database.QueryRowContext(ctx, "SELECT $1", message)
	err = originalRow.Scan(&echo.Message)
	return
}
