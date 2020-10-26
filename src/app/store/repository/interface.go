package repository

import (
	"github.com/Den1ske/GoMarket/src/app/entity"
)

// Repository ...
type Repository interface {
	List() ([] entity.Product, error)
	GetByID(int) (entity.Product, error)
}