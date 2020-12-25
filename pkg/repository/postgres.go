package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Config ...
type Config struct {
	Host	 string
	Port	 string
	Username string
	Password string
	DBname 	 string
	SSLMode  string	
}

// NewPostgresDB ...	
func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBname, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err2 := db.Ping()
	if err2 != nil {
		return nil, err2
	}

	return db, nil
}

