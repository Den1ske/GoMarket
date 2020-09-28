package model

import (
	"database/sql"
)

// Product ...
type Product struct {
	ID int64
	Title string
	Article string
	Price float64
	CategoryID sql.NullInt64
}