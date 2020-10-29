package productRepository

import (
	"github.com/Den1ske/GoMarket/src/app/entity"
	"log"
	"strconv"
)

type ProductRepositoryInterface interface {
	List() (*[] entity.Product, error)
	GetByID(int) (*entity.Product, error)
	Create(*entity.Product) (*entity.Product, error)
	GetByArticle(string) (*entity.Product, error)
}


type productRepository struct {
	store *Store
}

func (p *productRepository) List() (*[] entity.Product, error) {
	products := []entity.Product{}
	var err = p.store.db.Select(&products, "select * from products")
	if err != nil {
		return nil, err
	}

	return &products, nil
}

func (p *productRepository) GetByID(id int) (*entity.Product, error)  {

	product := entity.Product{}

	var err = p.store.db.Get(&product, "SELECT * FROM products WHERE id="+strconv.Itoa(id))
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Print(product)
	return &product, nil
}

func (p *productRepository) GetByArticle(article string) (*entity.Product, error)  {

	product := entity.Product{}

	var err = p.store.db.Get(&product, "SELECT * FROM products WHERE article="+article)
	if err != nil {
		//log.Fatal(err.Error())
		return nil, err
	}
	log.Print(product)
	return &product, nil
}

func (p *productRepository) Create(product *entity.Product) (*entity.Product, error)  {

	var _, err = p.store.db.NamedExec("INSERT INTO products (title, article, price, category_id) VALUES (:title, :article, :price, :category_id)", &product)

	if err != nil {
		log.Fatal(err.Error())
	}

	err = p.store.db.Get(&product.ID, " SELECT LAST_INSERT_ID()")

	return product, nil
}
