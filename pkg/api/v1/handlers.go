package v1

import (
	"encoding/json"
	services "golang-rest-api-clean-architecture/pkg/external-services"
	"net/http"
)

func GetCompanies(clients services.ClientSet, w http.ResponseWriter, req *http.Request) {
	companies := clients.DbClient().GetCompanies()
	w.Header().Set("Content-Type", "application/json")
	RespondwithJSON(w, 200, companies)
}

func RespondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	t := "data"
	if code > 300 {
		t = "errors"
	}
	response, _ := json.Marshal(map[string]interface{}{
		t: payload,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}
