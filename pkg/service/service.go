package service

import "github.com/tarpal7/all-commerce/pkg/repository"
// Authorization ...
type Authorization interface {

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
	return &Service{}
}
