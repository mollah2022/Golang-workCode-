package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler 

type Manager struct{
	globalMiddleWares []Middleware
}

func NewManager() *Manager{
	return &Manager{
		globalMiddleWares: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(middlewares ...Middleware) {
	mngr.globalMiddleWares = append(mngr.globalMiddleWares, middlewares...)
}

func(mngr *Manager) With(handler http.Handler,middlewares ...Middleware) http.Handler {
	n := handler

	for _, middleware := range middlewares {
		n = middleware(n)
	}
	return n
} 

func(mngr *Manager) WrapMux(handler http.Handler) http.Handler {
	n := handler

	for _, middleware := range mngr.globalMiddleWares {
		n = middleware(n)
	}
	return n
} 