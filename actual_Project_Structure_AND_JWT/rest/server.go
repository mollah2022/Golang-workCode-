package rest

import (
	"ecomerce/config"
	"ecomerce/rest/middleware"
	"fmt"
	"net/http"
	"strconv"
)

func Start(cnf config.Config) {

	manager := middleware.NewManager()
	manager.Use(
		middleware.PreFlight,
		middleware.Cors,
		middleware.Logger,
	)

    mux := http.NewServeMux()

	wrappedMux := manager.WrapMux(mux)
	InitRoutes(mux,manager)

	addr := ":"+strconv.Itoa(cnf.HttpPort)

	fmt.Println("Server Running on  ",addr)

	err:= http.ListenAndServe(addr,wrappedMux)

	if err != nil {
		fmt.Println("Error Starting the Server",err)
	}
}