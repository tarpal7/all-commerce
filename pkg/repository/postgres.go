package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable = "users"
	allcommerceListsTable = "allcommerce_lists"
	allcommerceItemsTable = "allcommerce_items"
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
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBname, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	return db, nil
}

