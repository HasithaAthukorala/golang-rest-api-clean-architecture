package repositories

import (
	"golang-rest-api-clean-architecture/pkg/entities"
	"gorm.io/gorm"
)

//go:generate moq -out company_repository.fake.go . CompanyRepository

type CompanyRepository interface {
	GetCompanies() []entities.Company
	GetCompanyById(id int) entities.Company
	GetCompanyByName(name string) entities.Company
	GetCompanyByCode(code string) entities.Company
	GetCompanyByWebsite(website string) entities.Company
	GetCompanyByPhone(phone string) entities.Company
	GetCompanyByCountry(phone string) entities.Company
	AddCompany(company *entities.Company) error
	RemoveCompany(id int) error
	GetLastAddedCompany() entities.Company
}

type repository struct {
	db *gorm.DB
}

func GetCompanyRepository(db *gorm.DB) CompanyRepository {
	return &repository{db: db}
}

func (repo *repository) GetCompanies() []entities.Company {
	var companies []entities.Company
	repo.db.Find(&companies)
	return companies
}

func (repo *repository) GetCompanyById(id int) entities.Company {
	company := entities.Company{
		Id: id,
	}
	repo.db.Find(&company)
	return company
}

func (repo *repository) GetCompanyByName(name string) entities.Company {
	company := entities.Company{
		Name: name,
	}
	repo.db.Find(&company)
	return company
}

func (repo *repository) GetCompanyByCode(code string) entities.Company {
	company := entities.Company{
		Code: code,
	}
	repo.db.Find(&company)
	return company
}

func (repo *repository) GetCompanyByWebsite(website string) entities.Company {
	company := entities.Company{
		Website: website,
	}
	repo.db.Find(&company)
	return company
}

func (repo *repository) GetCompanyByPhone(phone string) entities.Company {
	company := entities.Company{
		Phone: phone,
	}
	repo.db.Find(&company)
	return company
}

func (repo *repository) GetCompanyByCountry(country string) entities.Company {
	company := entities.Company{
		Country: country,
	}
	repo.db.Find(&company)
	return company
}

func (repo *repository) AddCompany(company *entities.Company) error {
	if dbc := repo.db.Create(&company); dbc.Error != nil {
		return dbc.Error
	}
	return nil
}

func (repo *repository) RemoveCompany(id int) error {
	company := entities.Company{
		Id: id,
	}
	if dbc := repo.db.Delete(&company); dbc.Error != nil {
		return dbc.Error
	}
	return nil
}

func (repo *repository) GetLastAddedCompany() entities.Company {
	var company entities.Company
	repo.db.Last(&company)
	return company
}
