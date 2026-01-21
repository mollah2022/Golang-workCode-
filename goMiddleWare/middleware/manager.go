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

func(mngr *Manager) With(next http.Handler,middlewares ...Middleware) http.Handler {
	n := next

	for _, middleware := range middlewares {
		n = middleware(n)
	}

	for _, globalMiddleware := range mngr.globalMiddleWares {
		n = globalMiddleware(n)
	}

	return n

} 