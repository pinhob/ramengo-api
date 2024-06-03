package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pinhob/ramengo-api/broth"
)

var brothService broth.Service

func ConfigureBrothsService() {
	brothService = broth.Service{
		Repository: &broth.RepositoryData{
			MockData: broth.MockDBTable,
		},
	}
}

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
