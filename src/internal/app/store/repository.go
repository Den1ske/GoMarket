package store

import (
	"github.com/Den1ske/GoMarket/src/internal/app/model"
)

// ProductRepository ...
type ProductRepository interface {
	Create(*model.Product) error
	FindByID(int64) (*model.Product, error)
}