package cmd

import (
	globalrouter "ecomerce/global_router"
	"ecomerce/handlers"
	"fmt"
	"net/http"
)

func Serve () {

	mux := http.NewServeMux()

	mux.Handle("GET /products",http.HandlerFunc(handlers.GetProduct))

	mux.Handle("POST /products",http.HandlerFunc(handlers.CreateProduct))

	mux.Handle("GET /products/{id}",http.HandlerFunc(handlers.GetProductById))

	fmt.Println("Server Running on : 8080")

	globalrouter := globalrouter.GlobalRouter(mux)

	err:= http.ListenAndServe(":8080",globalrouter)

	if err != nil {
		fmt.Println("Error Starting the Server",err)
	}

}