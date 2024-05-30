package main

import (
	"log"
	"net/http"

	ihttp "github.com/pinhob/ramengo-api/internal/http"
)

func main() {
	router := http.NewServeMux()

	ihttp.SetRoutes(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
