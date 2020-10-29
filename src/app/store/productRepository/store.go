package productRepository

import (
	"github.com/jmoiron/sqlx"
)



type Store struct {
	db *sqlx.DB
	repository *productRepository
}

func New(db *sqlx.DB) *Store  {
	return &Store{
		db: db,
	}
}

func (s *Store) Product() ProductRepositoryInterface {
	if s.repository != nil {
		return s.repository
	}

	s.repository = &productRepository{
		store: s,
	}

	return s.repository
}