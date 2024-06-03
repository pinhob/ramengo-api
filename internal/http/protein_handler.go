package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pinhob/ramengo-api/protein"
)

var proteinService protein.Service

func ConfigureProteinsService() {
	proteinService = protein.Service{
		Repository: &protein.RepositoryData{
			MockData: protein.MockDBProtein,
		},
	}
}

func HandleProteins(w http.ResponseWriter, r *http.Request) {
	proteins, err := proteinService.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(fmt.Sprintf("error: %s", err))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(proteins)
}
