package cmd

import (
	"ecomerce/config"
	"ecomerce/infra/db"
	"ecomerce/repo"
	"ecomerce/rest"
	"ecomerce/rest/handlers/product"
	"ecomerce/rest/handlers/user"
	"ecomerce/rest/middleware"
	"fmt"
	"os"
)

func Serve () {

	cnf := config.GetConfig()

	dbCon,err := db.NewConnection(cnf.DB)

	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbCon,"./migrations")

	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}

	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

   middlewares := middleware.NewMiddlewares(cnf)

	productHandler := product.NewHandler(middlewares,productRepo)
	userHandler := user.NewHandler(cnf,userRepo)

	server := rest.NewServer(cnf,productHandler,userHandler)
	server.Start()

}