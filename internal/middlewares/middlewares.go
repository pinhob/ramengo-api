package middlewares

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pinhob/ramengo-api/types"
)

const (
	errMsgApiKey = "x-api-key header missing"
)

func ValidateHeaderMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("x-api-key")

		if apiKey == "" {
			errRes, _ := json.Marshal(types.ErrorObject{Error: errMsgApiKey})
			w.WriteHeader(http.StatusForbidden)
			w.Write(errRes)
			// http.Error(w, , http.StatusForbidden)
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

func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("method %s, path: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}

type Middleware func(http.Handler) http.HandlerFunc

func MiddlewareChain(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}

		return next.ServeHTTP
	}
}
