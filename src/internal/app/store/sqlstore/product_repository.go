package sqlstore

import (
	"strconv"
	"github.com/Den1ske/GoMarket/src/internal/app/model"
)

// ProductRepository ...
type ProductRepository struct {
	store *Store
	perPage int
}

// List ...
func (r *ProductRepository) List(page int) ([] *model.Product, error) {
	var products [] *model.Product
	lowLevel := (page - 1) * r.perPage
	highLevel := page * r.perPage

	rows, err := r.store.db.Query("SELECT * FROM product LIMIT " + string(lowLevel) + "," + string(highLevel))
	if err != nil {
		return nil, err
	}
	i := 0

	for rows.Next() {
		products[i] = &model.Product{}

		if err := rows.Scan(
			&products[i].ID,
			&products[i].Title,
			&products[i].Article,
			&products[i].Price,
			&products[i].CategoryID,
			); err != nil {
				return nil, err
			}
		i++
		
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil	
}

// Create ...
func (r *ProductRepository) Create(p *model.Product) (*model.Product, error) {

	stmt, err := r.store.db.Prepare("INSERT INTO products (title, article, price, category_id) VALUES( ?, ?, ?, ? )")
	if err != nil {
		return nil, err
	}
	res, err := stmt.Exec(p.Title, p.Article, p.Price, p.CategoryID)
	if err != nil {
		return nil, err;
	}
	p.ID, err = res.LastInsertId()
	if err != nil {
		return nil, err;
	}
	return p, nil
}

// FindByID ...
func (r *ProductRepository) FindByID(id int64) (*model.Product, error) {
	p := &model.Product{}
	if err := r.store.db.QueryRow(
		"SELECT id, title, article, price, category_id FROM products WHERE id = " + strconv.FormatInt(id, 10),
	).Scan(
		&p.ID,
		&p.Title,
		&p.Article,
		&p.Price,
		&p.CategoryID,
	); err != nil {
		return nil, err
	}

	return p, nil
}