package api

import (
	services "golang-rest-api-clean-architecture/pkg/external-services"
	"net/http"
)

type HandlerFunc func(clientSet services.ClientSet, w http.ResponseWriter, req *http.Request)

type Route struct {
	Path               string
	Insecure           bool
	LocationValidation bool
	Method             string
	HandlerFunc        HandlerFunc
}
