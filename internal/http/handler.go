package http

import (
	"net/http"
)

func ConfigureServices() {
	ConfigureBrothsService()
	ConfigureProteinsService()
	ConfigureDishService()
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("RamenGo API"))
}
