package database

import (
	"github.com/sirupsen/logrus"
	"golang-rest-api-clean-architecture/pkg/config"
	"golang-rest-api-clean-architecture/pkg/external-services/database/repositories"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

//go:generate moq -out client.fake.go . DbClient

type DbClient interface {
	GetCompanyRepository() repositories.CompanyRepository
}

type dbClient struct {
	db *gorm.DB
}

func New(cfg *config.Config) DbClient {
	dsn := cfg.DbDSN
	entry := logrus.WithFields(logrus.Fields{"log_type": "gorm"})
	gormLogger := logger.New(entry, logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		IgnoreRecordNotFoundError: true,
		LogLevel:                  logger.Warn,
	})
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 gormLogger,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		logrus.Fatal(err)
	}
	return &dbClient{
		db: gormDB,
	}
}

func (client *dbClient) GetCompanyRepository() repositories.CompanyRepository {
	return repositories.GetCompanyRepository(client.db)
}
