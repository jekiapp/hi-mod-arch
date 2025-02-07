package db

import (
	"database/sql"
)

type DbConfig struct {
	Host string
	// ... etc
}

func InitDatabase(cfg DbConfig) (*sql.DB, error) {
	return &sql.DB{}, nil
}
