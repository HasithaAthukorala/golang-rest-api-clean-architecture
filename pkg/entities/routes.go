package entities

import "net/http"

type ClientSet interface{} // TODO: add location verifier and service bus

type HandlerFunc func(clientSet ClientSet, w http.ResponseWriter, req *http.Request)

type Route struct {
	Path               string
	Insecure           bool
	LocationValidation bool
	Method             string
	HandlerFunc        HandlerFunc
}
