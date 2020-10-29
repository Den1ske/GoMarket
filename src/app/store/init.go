package store

import (
	"github.com/Den1ske/GoMarket/src/app/connection"
	"github.com/Den1ske/GoMarket/src/app/store/productCategoryRepository"
	"github.com/Den1ske/GoMarket/src/app/store/productRepository"
)

var GlobalStore struct{
	ProductStore *productRepository.Store
	ProductCategoryStore *productCategoryRepository.Store
}

func init() {
	DB := connection.DB
	GlobalStore.ProductStore = productRepository.New(DB)
	GlobalStore.ProductCategoryStore = productCategoryRepository.New(DB)
}