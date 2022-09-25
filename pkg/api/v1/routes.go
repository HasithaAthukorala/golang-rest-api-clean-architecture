package v1

import (
	"golang-rest-api-clean-architecture/pkg/api"
	"net/http"
)

func Build() []api.Route {
	return []api.Route{
		{Path: "/companies", LocationValidation: false, Insecure: true, Method: http.MethodGet, HandlerFunc: GetCompanies},
	}
}
