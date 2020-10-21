package main

import (
	//"flag"
	//"log"
	//"github.com/BurntSushi/toml"
	"net/http"
	"database/sql"
    //"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"encoding/json"
)

var (
	configPath string
)

func init() {
	//flag.StringVar(&configPath, "config-path", "configs/server.toml", "Config file path")
}
type request struct {

}
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			err.Error();
			return
		}
		products, err := db()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err.Error())
		}
		log.Print(products)
		if products != nil {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(products);
		}
	})
	http.ListenAndServe(":80", nil)
}

type Product struct{
    ID int `json:"id"`
    Title string `json:"title"`
    Article sql.NullString `json:"article"`
	Price int `json:"price"`
	CategoryID sql.NullInt64 `json:"category_id"`
}

func db() ([]Product, error) {
    db, err := sql.Open("mysql", "root:root@tcp(db:3306)/go_market")
     
    if err != nil {
		return nil, err;
    } 
    defer db.Close()
    rows, err := db.Query("select * from go_market.products")
    if err != nil {
		return nil, err;
    }
    defer rows.Close()
    products := []Product{}
     
    for rows.Next(){
        p := Product{}
        err := rows.Scan(&p.ID, &p.Title, &p.Article, &p.Price, &p.CategoryID)
        if err != nil{
			return nil, err;
        }
        products = append(products, p)
	}
	return products, nil
}