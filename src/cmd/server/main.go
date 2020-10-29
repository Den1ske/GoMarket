package main

import (
	"github.com/Den1ske/GoMarket/src/app/connection"
	"github.com/Den1ske/GoMarket/src/app/server"

	"log"
)

func main() {


	err := server.ServeHTTP()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = connection.CloseDB()
		if err != nil {
			log.Fatal(err)
		}
	}()
}
