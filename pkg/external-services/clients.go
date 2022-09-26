package services

import (
	"golang-rest-api-clean-architecture/pkg/config"
	"golang-rest-api-clean-architecture/pkg/external-services/authenticator"
	"golang-rest-api-clean-architecture/pkg/external-services/database"
	location_verifier "golang-rest-api-clean-architecture/pkg/external-services/location-verifier"
	service_bus "golang-rest-api-clean-architecture/pkg/external-services/service-bus"
)

//go:generate moq -out clients.fake.go . ClientSet

type ClientSet interface {
	DbClient() database.DbClient
	AuthenticationClient() authenticator.AuthenticationClient
	LocationVerificationClient() location_verifier.LocationVerificationClient
	ServiceBusClient() service_bus.ServiceBusClient
}

type clientSet struct {
	dbClient                   database.DbClient
	authenticationClient       authenticator.AuthenticationClient
	locationVerificationClient location_verifier.LocationVerificationClient
	serviceBusClient           service_bus.ServiceBusClient
}

func NewClients(cfg *config.Config, stopCh <-chan struct{}) (ClientSet, error) {
	dbClient := database.New(cfg)
	authenticationClient := authenticator.New("Google")
	locationVerificationClient, err := location_verifier.New(cfg)
	if err != nil {
		return nil, err
	}
	serviceBusClient := service_bus.New(cfg, stopCh)
	return &clientSet{
		dbClient:                   dbClient,
		authenticationClient:       authenticationClient,
		locationVerificationClient: locationVerificationClient,
		serviceBusClient:           serviceBusClient,
	}, nil
}

func (clients *clientSet) DbClient() database.DbClient {
	return clients.dbClient
}

func (clients *clientSet) AuthenticationClient() authenticator.AuthenticationClient {
	return clients.authenticationClient
}

func (clients *clientSet) LocationVerificationClient() location_verifier.LocationVerificationClient {
	return clients.locationVerificationClient
}

func (clients *clientSet) ServiceBusClient() service_bus.ServiceBusClient {
	return clients.serviceBusClient
}
