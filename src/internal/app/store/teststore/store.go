package teststore

import (
	"github.com/Den1ske/GoMarket/src/internal/app/store"
)

// Store ...
type Store struct {
	productRepository *ProductRepository
}

// New ...
func New() *Store{
	
}

// Product ...
func (s *Store) Product() *store.ProductRepository {

	if s.productRepository != nil {
		return s.productRepository
	}

	s.productRepository = &ProductRepository {
		store: s,
	} 

	return s.productRepository
}