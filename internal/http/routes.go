package http

import (
	"net/http"
)

func SetRoutes(r *http.ServeMux) {
	r.HandleFunc("GET /broths", HandleBroths)
	r.HandleFunc("GET /proteins", HandleProteins)
	r.HandleFunc("POST /orders", HandleOrders)
	r.HandleFunc("GET /", HandleRoot)
}
