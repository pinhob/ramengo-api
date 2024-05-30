package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pinhob/ramengo-api/broth"
	"github.com/pinhob/ramengo-api/dish"
	"github.com/pinhob/ramengo-api/protein"
	"github.com/pinhob/ramengo-api/types"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("RamenGo API"))
}

func HandleBroths(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(broth.BrothsTable)
}

func HandleProteins(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(protein.ProteinsTable)
}

func HandleOrders(w http.ResponseWriter, r *http.Request) {
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

	dish, err := dish.GetDish(brothId, proteinId)
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
