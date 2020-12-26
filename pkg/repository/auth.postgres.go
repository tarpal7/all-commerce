package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	allcommerce "github.com/tarpal7/all-commerce"
)

// AuthPostgres ...
type AuthPostgres struct {
	db *sqlx.DB	
}

// NewAuthPostgres ...	
func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

// CreateUser ...
func (r *AuthPostgres) CreateUser(user allcommerce.User) (int, error) {
	
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash) values ($1, $2) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

// GetUser ...
func (r *AuthPostgres) GetUser(username, password string) (allcommerce.User, error) {
	var user allcommerce.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
