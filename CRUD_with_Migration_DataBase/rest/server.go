package rest

import (
	"ecomerce/config"
	"ecomerce/rest/handlers/product"
	"ecomerce/rest/handlers/user"
	"ecomerce/rest/middleware"
	"fmt"
	"net/http"
	"strconv"
)

type Server struct {
	cnf *config.Config
    productHandler *product.Handler
	userHandler   *user.Handler
}

func NewServer(cnf *config.Config,productHandler *product.Handler,userHandler   *user.Handler) *Server{
	return &Server{
	    cnf: cnf,
		productHandler: productHandler,
		userHandler: userHandler,
	}
}

func (server *Server) Start() {

	manager := middleware.NewManager()
	manager.Use(
		middleware.PreFlight,
		middleware.Cors,
		middleware.Logger,
	)

    mux := http.NewServeMux()

	wrappedMux := manager.WrapMux(mux)
	// InitRoutes(mux,manager)
	server.productHandler.RegisterRoutes(mux,manager)
	server.userHandler.RegisterRoutes(mux,manager)

	addr := ":"+strconv.Itoa(server.cnf.HttpPort)

	fmt.Println("Server Running on  ",addr)

	err:= http.ListenAndServe(addr,wrappedMux)

	if err != nil {
		fmt.Println("Error Starting the Server",err)
	}
}