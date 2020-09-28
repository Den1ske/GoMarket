package teststore

import (
	"github.com/Den1ske/GoMarket/src/internal/app/model"
)

// ProductRepository ...
type ProductRepository struct {
	store *Store
	products map[string]*model.Product
}

// Create ...
func (r *ProductRepository) Create(p *model.Product) error  {
	
}