package cmd

import (
	"ecomerce/handlers"
	"ecomerce/middleware"
	"net/http"
)

func InitRoutes(mux *http.ServeMux, manager *middleware.Manager) {
   
	 mux.Handle(
		"GET /home",
		manager.With(
		http.HandlerFunc(handlers.Test),
	 ),
   )

	 mux.Handle(
		"GET /products",
		manager.With(
		http.HandlerFunc(handlers.GetProduct),
		),
	)

	 mux.Handle(
		"POST /products",
		manager.With(
		http.HandlerFunc(handlers.CreateProduct),
		),
	)

	 mux.Handle(
		"GET /products/{id}",
		manager.With(
		http.HandlerFunc(handlers.GetProductById),
		),
	)

}