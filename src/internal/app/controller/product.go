package controller

import (
	"net/http"
	"io"
	"github.com/Den1ske/GoMarket/src/internal/app/store/sqlstore"
)

type ProductController struct{
	repository *sqlstore.ProductRepository
}


// List ...
func (p *ProductController) ListProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `Hello`)
	}
}