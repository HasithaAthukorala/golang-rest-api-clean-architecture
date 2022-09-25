package services

import (
	"golang-rest-api-clean-architecture/pkg/config"
	"golang-rest-api-clean-architecture/pkg/external-services/database"
)

type ClientSet interface {
	DbClient() database.DbClient
} // TODO: add location verifier and service bus

type clientSet struct {
	dbClient database.DbClient
}

func NewClients(cfg *config.Config) (ClientSet, error) {
	dbClient := database.New(cfg)
	return &clientSet{
		dbClient: dbClient,
	}, nil
}

func (clients *clientSet) DbClient() database.DbClient {
	return clients.dbClient
}
