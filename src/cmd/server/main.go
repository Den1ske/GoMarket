package main

import (
	"net/http"
	"encoding/json"
	"log"
	"github.com/Den1ske/GoMarket/src/app/connection"
)


func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		products, err := db()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err.Error())
		}
		log.Print(products)
		if products != nil {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(products)
		}
	})
	http.ListenAndServe(":80", nil)
}

type Product struct {
	ID         int			`json:"id"`
	Title      string		`json:"title"`
	Article    *string		`json:"article"`
	Price      int			`json:"price"`
	Category_ID *int64		`json:"category_id"`
}

func db() ([]Product, error) {
	products := []Product{}
	db, err := connection.Connect()
	if err != nil {
		return nil, err
	}
	if err = db.Select(&products, "select * from products"); err != nil {
		return nil, err
	}

	defer func() {
		err = db.Close()
	}()

	return products, err
}
