package product

type Product struct {
	ID         int			`json:"id"`
	Title      string		`json:"title"`
	Article    *string		`json:"article"`
	Price      int			`json:"price"`
	Category_ID *int64		`json:"category_id"`
}