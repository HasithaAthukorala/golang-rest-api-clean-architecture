package v1

import (
	"golang-rest-api-clean-architecture/pkg/entities"
	"net/http"
)

func GetCompanies(clientSet entities.ClientSet, w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}
