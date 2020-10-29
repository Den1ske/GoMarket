package connection

import (
	"github.com/Den1ske/GoMarket/src/app/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)
var DB *sqlx.DB

func init() {

	var err error
	DB, err = sqlx.Open("mysql", config.Config.DBUser+":"+config.Config.DBPass+"@tcp("+config.Config.DBHost+")/"+config.Config.DBName)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func CloseDB() error{
	var err = DB.Close()
	if err != nil {
		return err
	}
	return nil
}