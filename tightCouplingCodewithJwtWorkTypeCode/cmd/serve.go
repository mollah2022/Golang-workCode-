package cmd

import (
	"ecomerce/config"
	"ecomerce/rest"
	"ecomerce/rest/handlers/product"
	"ecomerce/rest/handlers/user"
	"ecomerce/rest/middleware"
)

func Serve () {

	cnf := config.GetConfig()

   middlewares := middleware.NewMiddlewares(cnf)

	productHandler := product.NewHandler(middlewares)
	userHandler := user.NewHandler()

	server := rest.NewServer(cnf,productHandler,userHandler)
	server.Start()

}