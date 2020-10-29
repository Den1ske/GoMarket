package productController

import "github.com/Den1ske/GoMarket/src/app/store/productRepository"

type Controller struct {
	store *productRepository.Store
	productController *productController
}

func New(store *productRepository.Store) *Controller  {
	return &Controller{
		store: store,
	}
}

func (c *Controller) Product() ProductControllerInterface {
	if c.productController != nil {
		return c.productController
	}

	c.productController = &productController{
		productStore: c.store,
	}

	return c.productController
}