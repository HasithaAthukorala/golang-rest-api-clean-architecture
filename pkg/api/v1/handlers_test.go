package v1

import (
	"encoding/json"
	"golang-rest-api-clean-architecture/pkg/entities"
	services "golang-rest-api-clean-architecture/pkg/external-services"
	"golang-rest-api-clean-architecture/pkg/external-services/authenticator"
	"golang-rest-api-clean-architecture/pkg/external-services/database"
	"golang-rest-api-clean-architecture/pkg/external-services/database/repositories"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var clientSet services.ClientSet
var testCompany entities.Company

func init() {
	testCompany = entities.Company{
		Id:      0,
		Name:    "Test Company",
		Country: "India",
		Code:    "tindco",
		Website: "test-company.com",
		Phone:   "+72323422342",
	}
	clientSet = &services.ClientSetMock{
		AuthenticationClientFunc: func() authenticator.AuthenticationClient {
			return &authenticator.AuthenticationClientMock{
				AuthenticateFunc: func(token string) bool {
					return true
				},
			}
		},
		DbClientFunc: func() database.DbClient {
			return &database.DbClientMock{
				GetCompanyRepositoryFunc: func() repositories.CompanyRepository {
					return &repositories.CompanyRepositoryMock{
						AddCompanyFunc: nil,
						GetCompaniesFunc: func() []entities.Company {
							return []entities.Company{
								testCompany,
							}
						},
						GetCompanyByCodeFunc: func(code string) entities.Company {
							return entities.Company{Code: code}
						},
						GetCompanyByCountryFunc: func(country string) entities.Company {
							return entities.Company{Country: country}
						},
						GetCompanyByIdFunc: func(id int) entities.Company {
							return entities.Company{Id: id}
						},
						GetCompanyByNameFunc: func(name string) entities.Company {
							return entities.Company{Name: name}
						},
						GetCompanyByPhoneFunc: func(phone string) entities.Company {
							return entities.Company{Phone: phone}
						},
						GetCompanyByWebsiteFunc: func(website string) entities.Company {
							return entities.Company{Website: website}
						},
						GetLastAddedCompanyFunc: func() entities.Company {
							return entities.Company{Name: "LastAdded"}
						},
						RemoveCompanyFunc: nil,
					}
				},
			}
		},
		LocationVerificationClientFunc: nil,
		ServiceBusClientFunc:           nil,
	}
}

type GetCompaniesResp struct {
	Data []entities.Company `json:"data"`
}

func TestGetCompanies(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/companies", nil)
	w := httptest.NewRecorder()
	GetCompanies(clientSet, w, req)
	res := w.Result()
	defer func() {
		_ = res.Body.Close()
	}()
	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("unexpected error while reading the res body: %v", err)
	}

	var response GetCompaniesResp
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		t.Errorf("unexpected error while unmarshalling the response body to a company: %v", err)
	}

	if len(response.Data) != 1 {
		t.Errorf("expected a list of companies of length 1, but received: %v", response.Data)
	}

	if !reflect.DeepEqual(response.Data[0], testCompany) {
		t.Errorf("error, received = %v, expected %v", response.Data[0], testCompany)
	}
}

type GetCompanyResp struct {
	Data entities.Company `json:"data"`
}

func TestGetCompanyWithoutParameters(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/company", nil)
	w := httptest.NewRecorder()
	GetCompany(clientSet, w, req)
	res := w.Result()
	defer func() {
		_ = res.Body.Close()
	}()
	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("unexpected error while reading the res body: %v", err)
	}

	var response GetCompanyResp
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		t.Errorf("unexpected error while unmarshalling the response body to a company: %v", err)
	}

	if response.Data.Name != "LastAdded" {
		t.Errorf("expected the name of the company as `LastAdded`, but received: %s", response.Data.Name)
	}
}

func TestGetCompanyWithName(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/company?name=NAME", nil)
	w := httptest.NewRecorder()
	GetCompany(clientSet, w, req)
	res := w.Result()
	defer func() {
		_ = res.Body.Close()
	}()
	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("unexpected error while reading the res body: %v", err)
	}

	var response GetCompanyResp
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		t.Errorf("unexpected error while unmarshalling the response body to a company: %v", err)
	}

	if response.Data.Name != "NAME" {
		t.Errorf("expected the name of the company as `NAME`, but received: %s", response.Data.Name)
	}
}

func TestGetCompanyWithCountry(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/company?country=COUNTRY", nil)
	w := httptest.NewRecorder()
	GetCompany(clientSet, w, req)
	res := w.Result()
	defer func() {
		_ = res.Body.Close()
	}()
	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("unexpected error while reading the res body: %v", err)
	}

	var response GetCompanyResp
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		t.Errorf("unexpected error while unmarshalling the response body to a company: %v", err)
	}

	if response.Data.Country != "COUNTRY" {
		t.Errorf("expected the name of the company as `COUNTRY`, but received: %s", response.Data.Country)
	}
}

func TestGetCompanyWithCode(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/company?code=CODE", nil)
	w := httptest.NewRecorder()
	GetCompany(clientSet, w, req)
	res := w.Result()
	defer func() {
		_ = res.Body.Close()
	}()
	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("unexpected error while reading the res body: %v", err)
	}

	var response GetCompanyResp
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		t.Errorf("unexpected error while unmarshalling the response body to a company: %v", err)
	}

	if response.Data.Code != "CODE" {
		t.Errorf("expected the name of the company as `CODE`, but received: %s", response.Data.Code)
	}
}

func TestGetCompanyWithWebsite(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/company?website=WEBSITE", nil)
	w := httptest.NewRecorder()
	GetCompany(clientSet, w, req)
	res := w.Result()
	defer func() {
		_ = res.Body.Close()
	}()
	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("unexpected error while reading the res body: %v", err)
	}

	var response GetCompanyResp
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		t.Errorf("unexpected error while unmarshalling the response body to a company: %v", err)
	}

	if response.Data.Website != "WEBSITE" {
		t.Errorf("expected the name of the company as `WEBSITE`, but received: %s", response.Data.Website)
	}
}

func TestGetCompanyWithPhone(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/company?phone=PHONE", nil)
	w := httptest.NewRecorder()
	GetCompany(clientSet, w, req)
	res := w.Result()
	defer func() {
		_ = res.Body.Close()
	}()
	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("unexpected error while reading the res body: %v", err)
	}

	var response GetCompanyResp
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		t.Errorf("unexpected error while unmarshalling the response body to a company: %v", err)
	}

	if response.Data.Phone != "PHONE" {
		t.Errorf("expected the name of the company as `PHONE`, but received: %s", response.Data.Phone)
	}
}
