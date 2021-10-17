package server

import (
	"database/sql"
	"fmt"

	"github.com/aleksvdim/go-grpc-gateway-template/internal/app/config"
)

func setUpDb(dbCfg config.Database) (*sql.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s database=%s sslmode=%s", dbCfg.User, dbCfg.Password, dbCfg.Database, dbCfg.SSLMode)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("db ping: %w", err)
	}

	return db, nil
}
