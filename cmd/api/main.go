package main

import (
	"log"
	"net/http"

	internal_http "github.com/pinhob/ramengo-api/internal/http"
	"github.com/pinhob/ramengo-api/internal/middlewares"
)

func main() {
	router := http.NewServeMux()

	middlewares := middlewares.MiddlewareChain(
		middlewares.RequestLoggerMiddleware,
		middlewares.CorsMiddleware,
		middlewares.ValidateHeaderMiddleware,
	)

	internal_http.ConfigureServices()
	internal_http.SetRoutes(router)
	log.Fatal(http.ListenAndServe(":8080", middlewares(router)))
}
