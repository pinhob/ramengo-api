package main

import (
	"log"
	"net/http"

	ihttp "github.com/pinhob/ramengo-api/internal/http"
)

func main() {
	router := http.NewServeMux()

	ihttp.SetRoutes(router)
	log.Fatal(http.ListenAndServe(":8080", CorsMiddleware(router)))
}

func CorsMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS, GET, DELETE, PUT")
		w.Header().Add("Access-Control-Allow-Headers", "content-type, x-api-key")
		w.Header().Add("Access-Control-Max-Age", "86400")

		next.ServeHTTP(w, r)
	}
}
