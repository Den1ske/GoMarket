package controller

import (
	"encoding/json"
	"github.com/Den1ske/GoMarket/src/app/store/repository"
	"log"
	"net/http"
	"github.com/jmoiron/sqlx"
)
type ProductController struct {
	DB *sqlx.DB
}
func (c *ProductController) List(w http.ResponseWriter, r *http.Request)  {

	pr := &repository.ProductRepository{
		DB: c.DB,
	}

	products, err := pr.List()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	}
	log.Print(products)
	if products != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(products)
	}
}