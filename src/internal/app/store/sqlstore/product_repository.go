package sqlstore

import (
	"strconv"
	"github.com/Den1ske/GoMarket/src/internal/app/model"
)

// ProductRepository ...
type ProductRepository struct {
	store *Store
}

// Create ...
func (r *ProductRepository) Create(p *model.Product) error {

	stmt, err := r.store.db.Prepare("INSERT INTO products (title, article, price, category_id) VALUES( ?, ?, ?, ? )")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(p.Title, p.Article, p.Price, p.CategoryID); err != nil {
		return err;
	}
	return nil;
	// id, err := res.LastInsertId()
	// if err != nil {
	// 	return err;
	// }
	// p.ID = id
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