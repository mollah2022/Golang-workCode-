package middleware

import (
	"log"
	"net/http"
)

func Hudai(next http.Handler) http.Handler {
    
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println("Ami MiddleWare : Ami Hudai Print Hobo...")

		next.ServeHTTP(w, r)
	})

}