package sqlstore

import (
	"database/sql"
	"github.com/Den1ske/GoMarket/src/internal/app/store"
	_ "github.com/go-sql-driver/mysql" //
)

// Store ...
type Store struct {
	db *sql.DB
	productRepository *ProductRepository
}

// New ...
func New(db *sql.DB) *Store{
	return &Store {
		db: db,
	}
}

// Product ...
func (s *Store) Product() store.ProductRepository {

	if s.productRepository != nil {
		return s.productRepository
	}
	
	s.productRepository = &ProductRepository {
		store: s,
	} 

	return s.productRepository
}