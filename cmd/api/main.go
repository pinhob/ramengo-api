package main

import (
	"log"
	"net/http"

	ihttp "github.com/pinhob/ramengo-api/internal/http"
	"github.com/pinhob/ramengo-api/internal/middlewares"
)

func main() {
	router := http.NewServeMux()
	middlewares := middlewares.MiddlewareChain(
		middlewares.RequestLoggerMiddleware,
		middlewares.CorsMiddleware,
		middlewares.ValidateHeaderMiddleware,
	)

	ihttp.ConfigureServices()
	ihttp.SetRoutes(router)
	log.Fatal(http.ListenAndServe(":8080", middlewares(router)))
}
