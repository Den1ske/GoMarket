package sqlstore_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/Den1ske/GoMarket/src/internal/app/store/sqlstore"
	"github.com/Den1ske/GoMarket/src/internal/app/model"
)

func TestProductRepository_Create(t *testing.T)  {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("products")

	s := sqlstore.New(db)
	p := model.TestProduct(t)
	assert.NoError(t, s.Product().Create(p))
	assert.NotNil(t, p)
}

// func TestProductRepository_FindByID(t *testing.T)  {
// 	db, teardown := sqlstore.TestDB(t, databaseURL)
// 	defer teardown("products")
	
// 	s := sqlstore.New(db)

// 	var id int64 = 1

// 	_, err := s.Product().FindByID(id)

// 	assert.Error(t, err)

// 	p1 := s.Product().Create(&model.Product{
// 		Title: "product",
// 		Article: "00001TEST",
// 		Price: 200.50,
// 		CategoryID: 0,
// 	})
// 	p2, err := s.Product().FindByID(p1.ID)

// 	assert.NoError(t, err)
// 	assert.NotNil(t, p2)
// }