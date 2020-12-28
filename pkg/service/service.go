package service

import (
	"github.com/tarpal7/all-commerce"
	"github.com/tarpal7/all-commerce/pkg/repository"
)
// Authorization ...
type Authorization interface {
	CreateUser(user allcommerce.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

// CommerceList ...
type CommerceList interface {

}

// CommerceItem ...
type CommerceItem interface {

}

// Service ...
type Service struct {
	Authorization
	CommerceList
	CommerceItem
}

// NewService ...
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
