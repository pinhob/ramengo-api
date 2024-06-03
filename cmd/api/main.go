package main

import (
	"log"
	"net/http"

	ihttp "github.com/pinhob/ramengo-api/internal/http"
)

func main() {
	router := http.NewServeMux()

	ihttp.SetRoutes(router)
	log.Fatal(http.ListenAndServe(":8080", ValidateHeaderMiddleware(router)))
}

func ValidateHeaderMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("x-api-key")

		if apiKey == "" {
			http.Error(w, "x-api-key header missing", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func CorsMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS, GET, DELETE, PUT")
		w.Header().Add("Access-Control-Allow-Headers", "content-type, x-api-key")
		w.Header().Add("Access-Control-Max-Age", "86400")

		if r.Method == "OPTIONS" {
			http.Error(w, "No Content", http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	}
}
