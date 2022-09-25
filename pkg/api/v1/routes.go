package v1

import (
	"golang-rest-api-clean-architecture/pkg/entities"
	"net/http"
)

func Build() []entities.Route {
	return []entities.Route{
		{Path: "/companies", LocationValidation: false, Insecure: true, Method: http.MethodGet, HandlerFunc: GetCompanies},
	}
}
