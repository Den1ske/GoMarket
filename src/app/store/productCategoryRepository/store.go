package productCategoryRepository

import (
	"github.com/jmoiron/sqlx"
)



type Store struct {
	db *sqlx.DB
	repository *productCategoryRepository
}

func New(db *sqlx.DB) *Store  {
	return &Store{
		db: db,
	}
}

func (s *Store) Category() ProductCategoryRepositoryInterface {
	if s.repository != nil {
		return s.repository
	}

	s.repository = &productCategoryRepository{
		store: s,
	}

	return s.repository
}