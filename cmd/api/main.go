package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pinhob/ramengo-api/broth"
	"github.com/pinhob/ramengo-api/dish"
	"github.com/pinhob/ramengo-api/protein"
	"github.com/pinhob/ramengo-api/types"
)

func main() {
	http.HandleFunc("GET /broths", handleBroths)
	http.HandleFunc("GET /proteins", handleProteins)
	http.HandleFunc("POST /orders", handleOrders)
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

func handleOrders(w http.ResponseWriter, r *http.Request) {
	var ids types.OrderRequest

	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	brothId, err := strconv.Atoi(ids.BrothId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("error: brothId must be a number")
		return
	}

	proteinId, err := strconv.Atoi(ids.ProteinId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("error: proteinId must be a number")
		return
	}

	err, dish := dish.GetDish(brothId, proteinId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	order := types.OrderRespone{
		ID:          123,
		Description: dish.Name,
		Image:       dish.Image,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}
