package controllers

import (
	"github.com/Den1ske/GoMarket/src/app/controllers/productController"
	"github.com/Den1ske/GoMarket/src/app/store"
)

var ProductController productController.ProductControllerInterface

func init() {
	ProductController = productController.New(store.GlobalStore.ProductStore).Product()
}