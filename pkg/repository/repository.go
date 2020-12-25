package repository

import "github.com/jmoiron/sqlx"

// Authorization ...
type Authorization interface {

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
	return &Repository{}
}
