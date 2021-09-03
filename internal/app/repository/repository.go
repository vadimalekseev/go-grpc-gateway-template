package repository

import (
	"database/sql"
	"fmt"
)

type Repository struct {
	database *sql.DB
}

func New(db *sql.DB) Repository {
	return Repository{database: db}
}

func (r *Repository) GetLink(short string) (original string) {
	originalRow := r.database.QueryRow("SELECT original from links where shortened ==  ?", short)
	err := originalRow.Scan(&original)
	if err != nil {
		fmt.Printf("corresponding link was not found: %v", err)
	}

	return
}

func (r *Repository) SetLink(original, short string) {
	_, err := r.database.Query("INSERT INTO links VALUES (?, ?)", original, short)
	if err != nil {
		fmt.Printf("error inserting new shortened link: %v", err)
	}
}
