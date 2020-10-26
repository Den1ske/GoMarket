package server

import (
	"github.com/Den1ske/GoMarket/src/app/controller"
	"log"
	"net/http"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	DB *sqlx.DB
}

func (s *Server) Router() {
	pc := &controller.ProductController{DB:s.DB}
	http.HandleFunc("/", pc.List)
}

func (s *Server) ServeHTTP()  {
	s.Router()
	var err = http.ListenAndServe(":80", nil)
	if err != nil  {
		log.Fatalln(err)
	}
}