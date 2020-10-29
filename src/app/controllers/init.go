package controllers

import (
	"github.com/Den1ske/GoMarket/src/app/controllers/productCategoryController"
	"github.com/Den1ske/GoMarket/src/app/controllers/productController"
	"github.com/Den1ske/GoMarket/src/app/store"
)

var ProductController productController.ProductControllerInterface
var ProductCategoryController productCategoryController.ProductCategoryControllerInterface

func init() {
	ProductController = productController.New(store.GlobalStore.ProductStore).Product()
	ProductCategoryController = productCategoryController.New(store.GlobalStore.ProductCategoryStore).Category()
}
