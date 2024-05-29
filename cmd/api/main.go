package main

import (
	"encoding/json"
	"net/http"

	"github.com/pinhob/ramengo-api/broth"
	"github.com/pinhob/ramengo-api/protein"
)

func main() {
	http.HandleFunc("GET /broths", handleBroths)
	http.HandleFunc("GET /proteins", handleProteins)
	http.HandleFunc("GET /", handleRoot)

	http.ListenAndServe(":8080", nil)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("RamenGo API"))
}

func handleBroths(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(broth.BrothsTable)
}

func handleProteins(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(protein.ProteinsTable)
}
