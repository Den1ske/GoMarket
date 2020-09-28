package model

import (
	"database/sql"
	"testing"
)

// TestProduct ...
func TestProduct(t *testing.T) *Product {
	t.Helper()

	return &Product{
		Title:    "Продукты",
		Article: "334334HR",
		Price: 3000.50,
		CategoryID: sql.NullInt64{Int64: 0},
	}
}