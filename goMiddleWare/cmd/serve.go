package cmd

import (
	globalrouter "ecomerce/global_router"
	"ecomerce/middleware"
	"fmt"
	"net/http"
)

func Server () {

	mux := http.NewServeMux()

	manager := middleware.NewManager()

	manager.Use(middleware.Logger,middleware.Hudai)

	InitRoutes(mux,manager)

	fmt.Println("Server Running on : 8080")

	globalrouter := globalrouter.GlobalRouter(mux)

	err:= http.ListenAndServe(":8080",globalrouter)

	if err != nil {
		fmt.Println("Error Starting the Server",err)
	}

}