package product

import "ecomerce/rest/middleware"

type Handler struct {
	middlewares *middleware.Middlewares
}

func NewHandler(middlewares *middleware.Middlewares) *Handler {
	return &Handler{
		middlewares: middlewares,
	}
}