package cmd

import (
	"ecomerce/middleware"
	"fmt"
	"net/http"
)

func Server () {

	manager := middleware.NewManager()

	manager.Use(
		middleware.PreFlight,
		middleware.Cors,
		middleware.Logger,
	)

    mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	InitRoutes(mux,manager)

	fmt.Println("Server Running on : 8080")

	err:= http.ListenAndServe(":8080",wrappedMux)

	if err != nil {
		fmt.Println("Error Starting the Server",err)
	}

}