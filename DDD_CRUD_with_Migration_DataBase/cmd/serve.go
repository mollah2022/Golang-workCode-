package cmd

import (
	"ecomerce/config"
	"ecomerce/infra/db"
	"ecomerce/product"
	"ecomerce/repo"
	"ecomerce/rest"
	productHandler "ecomerce/rest/handlers/product"
	userHandler "ecomerce/rest/handlers/user"
	"ecomerce/rest/middleware"
	"ecomerce/user"
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
   
	//repos
	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)
    
	//domain
	usrSvc := user.NewService(userRepo)
	prdctSvc := product.NewService(productRepo)


   middlewares := middleware.NewMiddlewares(cnf)

	productHandler := productHandler.NewHandler(middlewares,prdctSvc)
	userHandler := userHandler.NewHandler(cnf,usrSvc)

	server := rest.NewServer(cnf,productHandler,userHandler)
	server.Start()

}