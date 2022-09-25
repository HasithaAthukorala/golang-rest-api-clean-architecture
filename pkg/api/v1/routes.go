package v1

import (
	"golang-rest-api-clean-architecture/pkg/api"
	"net/http"
)

func Build() []api.Route {
	return []api.Route{
		{Path: "/companies", Insecure: true, Method: http.MethodGet, HandlerFunc: GetCompanies},

		// query parameters can be passed to filter the org, ex: /company?name=ABC, /company?country=ABC
		{Path: "/company", Insecure: true, Method: http.MethodGet, HandlerFunc: GetCompany},
		{Path: "/company", Insecure: false, Method: http.MethodPost, HandlerFunc: AddCompany},

		{Path: "/company/{id}", Insecure: true, Method: http.MethodGet, HandlerFunc: GetCompanyById},
		{Path: "/company/{id}", Insecure: false, Method: http.MethodDelete, HandlerFunc: RemoveCompany},
	}
}
