package repository

import (
	"github.com/Den1ske/GoMarket/src/app/entity"
	"github.com/jmoiron/sqlx"
)


type ProductRepository struct {
	DB *sqlx.DB
}

func (p *ProductRepository) List() ([] entity.Product, error) {
	products := []entity.Product{}
	var err = p.DB.Select(&products, "select * from products")
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *ProductRepository) GetByID(id int) (entity.Product, error)  {
	var product = entity.Product{
		ID:          id ,
		Title:       "test",
		Article:     nil,
		Price:       2242114,
		Category_ID: nil,
	}
	return product, nil
}
