package main

import (
	"github.com/Den1ske/GoMarket/src/app/connection"
	"github.com/Den1ske/GoMarket/src/app/server"
	"log"
)


func main() {

	db, err := connection.Connect()
	if err != nil {
		log.Fatal(err)
		return
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	s := server.Server{DB:db}
	s.ServeHTTP()
}


//func db() ([]Product, error) {
//	products := []Product{}
//	db, err := connection.Connect()
//	if err != nil {
//		return nil, err
//	}
//	if err = db.Select(&products, "select * from products"); err != nil {
//		return nil, err
//	}
//
//	defer func() {
//		err = db.Close()
//	}()
//
//	return products, err
//}
