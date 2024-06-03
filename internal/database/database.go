package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"program_service/config"
)

func New(cfg *config.DBConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Database))

	if err != nil {
		return nil, err
	}

	return db, nil
}

func Release(db *sqlx.DB) error {
	return db.Close()
}
