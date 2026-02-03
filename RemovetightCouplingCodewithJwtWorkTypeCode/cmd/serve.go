package cmd

import (
	"ecomerce/config"
	"ecomerce/repo"
	"ecomerce/rest"
	"ecomerce/rest/handlers/product"
	"ecomerce/rest/handlers/user"
	"ecomerce/rest/middleware"
)

func Serve () {

	cnf := config.GetConfig()

	 productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo()

   middlewares := middleware.NewMiddlewares(cnf)

	productHandler := product.NewHandler(middlewares,productRepo)
	userHandler := user.NewHandler(cnf,userRepo)

	server := rest.NewServer(cnf,productHandler,userHandler)
	server.Start()

}