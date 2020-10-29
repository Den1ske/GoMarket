package productCategoryController

import "github.com/Den1ske/GoMarket/src/app/store/productCategoryRepository"

type Controller struct {
	store *productCategoryRepository.Store
	productCategoryController *productCategoryController
}

func New(store *productCategoryRepository.Store) *Controller  {
	return &Controller{
		store: store,
	}
}

func (c *Controller) Category() ProductCategoryControllerInterface {
	if c.productCategoryController != nil {
		return c.productCategoryController
	}

	c.productCategoryController = &productCategoryController{
		productCategoryStore: c.store,
	}

	return c.productCategoryController
}