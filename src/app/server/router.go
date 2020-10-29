package server

import (
	"github.com/Den1ske/GoMarket/src/app/controllers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var Router chi.Router


func Routing()  {

	Router = chi.NewRouter()
	Router.Use(middleware.RequestID)
	Router.Use(middleware.RealIP)
	Router.Use(middleware.Logger)
	Router.Use(middleware.Recoverer)
	Router.Route("/products", func(r chi.Router) {
		r.Get("/", controllers.ProductController.List)
		r.Post("/", controllers.ProductController.Create)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", controllers.ProductController.GetByID)
		})
	})
}