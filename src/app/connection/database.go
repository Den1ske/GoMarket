package connection

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)


func Connect() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", "root:root@tcp(db:3306)/go_market")
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return db, nil
}