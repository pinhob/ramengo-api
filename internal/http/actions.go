package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

var brothService broth.Service

func HandleBroths(w http.ResponseWriter, r *http.Request) {
	broths, err := brothService.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(fmt.Sprintf("error: %s", err))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(broths)
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

	req, err := http.NewRequest("POST", "https://api.tech.redventures.com.br/orders/generate-id", &bytes.Buffer{})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", "ZtVdh8XQ2U8pWI2gmZ7f796Vh8GllXoN7mr0djNf")

	client := &http.Client{}
	res, errRes := client.Do(req)
	if errRes != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	var resJson map[string]string
	errUnmarshall := json.Unmarshal(resBody, &resJson)
	if errUnmarshall != nil {

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(fmt.Sprintf("erro aqui %v %v", err, string(resBody)))
		return
	}

	order := types.OrderRespone{
		ID:          resJson["orderId"],
		Description: dish.Name,
		Image:       dish.Image,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}
