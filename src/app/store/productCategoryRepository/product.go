package productCategoryRepository

import (
	"github.com/Den1ske/GoMarket/src/app/entity"
	"log"
	"strconv"
)

type ProductCategoryRepositoryInterface interface {
	List() (*[] entity.ProductCategory, error)
	GetByID(int) (*entity.ProductCategory, error)
	Create(*entity.ProductCategory) (*entity.ProductCategory, error)
}


type productCategoryRepository struct {
	store *Store
}

func (p *productCategoryRepository) List() (*[] entity.ProductCategory, error) {
	categories := []entity.ProductCategory{}
	var err = p.store.db.Select(&categories, "select * from product_categories")
	if err != nil {
		return nil, err
	}

	return &categories, nil
}

func (p *productCategoryRepository) GetByID(id int) (*entity.ProductCategory, error)  {

	product := entity.ProductCategory{}

	var err = p.store.db.Get(&product, "SELECT * FROM product_categories WHERE id="+strconv.Itoa(id))
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Print(product)
	return &product, nil
}

func (p *productCategoryRepository) Create(category *entity.ProductCategory) (*entity.ProductCategory, error)  {

	var _, err = p.store.db.NamedExec("INSERT INTO product_categories (title) VALUES (:title)", &category)

	if err != nil {
		log.Fatal(err.Error())
	}

	err = p.store.db.Get(&category.ID, " SELECT LAST_INSERT_ID()")

	return category, nil
}
