package http

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/pinhob/ramengo-api/dish"
	"github.com/pinhob/ramengo-api/types"
)

const (
	errorMsgOrderNotPlaced = "could not place order"
	errorMsgIdsAreRequired = "both brothId and proteinId are required"
)

var dishService dish.Service

func ConfigureDishService() {
	dishService = dish.Service{
		Repository: &dish.RepositoryData{
			MockData: dish.MockDBDishes,
		},
	}
}

func HandleOrders(w http.ResponseWriter, r *http.Request) {
	var ids types.OrderRequest

	err := json.NewDecoder(r.Body).Decode(&ids)
	if err != nil {
		errRes, _ := json.Marshal(types.ErrorObject{Error: errorMsgIdsAreRequired})
		w.WriteHeader(http.StatusForbidden)
		w.Write(errRes)
		return
	}

	brothId, err := strconv.Atoi(ids.BrothId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	proteinId, err := strconv.Atoi(ids.ProteinId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	dish, err := dishService.GetDishByCondiments(brothId, proteinId)
	if err != nil {
		errorResponse, _ := json.Marshal(types.ErrorObject{Error: errorMsgOrderNotPlaced})

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(errorResponse)
		return
	}

	orderId, err := getOrderId()
	if err != nil {
		errorResponse, _ := json.Marshal(types.ErrorObject{Error: errorMsgOrderNotPlaced})

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(errorResponse)
		return
	}

	order := types.OrderRespone{
		ID:          orderId,
		Description: dish.Name,
		Image:       dish.Image,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

func getOrderId() (string, error) {
	req, err := http.NewRequest("POST", "https://api.tech.redventures.com.br/orders/generate-id", &bytes.Buffer{})
	if err != nil {
		log.Printf("Error making a new request to generate id endpoint: %v", err)
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", "ZtVdh8XQ2U8pWI2gmZ7f796Vh8GllXoN7mr0djNf")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Requisition to generate id error: %v", err)
		return "", err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error reading requisiton response body: %v", err)
		return "", err
	}

	var orderIdObj map[string]string

	err = json.Unmarshal(resBody, &orderIdObj)
	if err != nil {
		log.Printf("Error unmarshalling the requisition response body: %v", err)
		return "", err
	}

	return orderIdObj["orderId"], nil
}
