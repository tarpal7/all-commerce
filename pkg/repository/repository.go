package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/tarpal7/all-commerce"
)
// Authorization ...
type Authorization interface {
	CreateUser(user allcommerce.User) (int, error)
	GetUser(username, password string) (allcommerce.User, error)
}

// CommerceList ...
type CommerceList interface {

}

// CommerceItem ...
type CommerceItem interface {

}

// Repository ...
type Repository struct {
	Authorization
	CommerceList
	CommerceItem
}

// NewRepository ...
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
