package v1

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"golang-rest-api-clean-architecture/pkg/entities"
	services "golang-rest-api-clean-architecture/pkg/external-services"
	"golang-rest-api-clean-architecture/pkg/internal"
	"net/http"
	"strconv"
)

func GetCompanies(clients services.ClientSet, w http.ResponseWriter, req *http.Request) {
	companies := clients.DbClient().GetCompanyRepository().GetCompanies()
	w.Header().Set("Content-Type", "application/json")
	internal.RespondWithJSON(w, http.StatusOK, companies)
}

func GetCompany(clients services.ClientSet, w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	if len(name) > 0 {
		GetCompanyByName(clients, w, req)
		return
	}

	code := req.URL.Query().Get("code")
	if len(code) > 0 {
		GetCompanyByCode(clients, w, req)
		return
	}

	website := req.URL.Query().Get("website")
	if len(website) > 0 {
		GetCompanyByWebsite(clients, w, req)
		return
	}

	phone := req.URL.Query().Get("phone")
	if len(phone) > 0 {
		GetCompanyByPhone(clients, w, req)
		return
	}

	country := req.URL.Query().Get("country")
	if len(country) > 0 {
		GetCompanyByCountry(clients, w, req)
		return
	}

	company := clients.DbClient().GetCompanyRepository().GetLastAddedCompany()
	w.Header().Set("Content-Type", "application/json")
	internal.RespondWithJSON(w, http.StatusOK, company)
}

func GetCompanyById(clients services.ClientSet, w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		internal.RespondWithJSON(w, http.StatusBadRequest, err)
		return
	}
	companies := clients.DbClient().GetCompanyRepository().GetCompanyById(id)
	w.Header().Set("Content-Type", "application/json")
	internal.RespondWithJSON(w, http.StatusOK, companies)
}

func GetCompanyByName(clients services.ClientSet, w http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	companies := clients.DbClient().GetCompanyRepository().GetCompanyByName(name)
	w.Header().Set("Content-Type", "application/json")
	internal.RespondWithJSON(w, http.StatusOK, companies)
}

func GetCompanyByCode(clients services.ClientSet, w http.ResponseWriter, req *http.Request) {
	code := req.URL.Query().Get("code")
	companies := clients.DbClient().GetCompanyRepository().GetCompanyByCode(code)
	w.Header().Set("Content-Type", "application/json")
	internal.RespondWithJSON(w, http.StatusOK, companies)
}

func GetCompanyByWebsite(clients services.ClientSet, w http.ResponseWriter, req *http.Request) {
	website := req.URL.Query().Get("website")
	companies := clients.DbClient().GetCompanyRepository().GetCompanyByWebsite(website)
	w.Header().Set("Content-Type", "application/json")
	internal.RespondWithJSON(w, http.StatusOK, companies)
}

func GetCompanyByPhone(clients services.ClientSet, w http.ResponseWriter, req *http.Request) {
	phone := req.URL.Query().Get("phone")
	companies := clients.DbClient().GetCompanyRepository().GetCompanyByPhone(phone)
	w.Header().Set("Content-Type", "application/json")
	internal.RespondWithJSON(w, http.StatusOK, companies)
}

func GetCompanyByCountry(clients services.ClientSet, w http.ResponseWriter, req *http.Request) {
	country := req.URL.Query().Get("country")
	companies := clients.DbClient().GetCompanyRepository().GetCompanyByCountry(country)
	w.Header().Set("Content-Type", "application/json")
	internal.RespondWithJSON(w, http.StatusOK, companies)
}

func AddCompany(clients services.ClientSet, w http.ResponseWriter, req *http.Request) {
	company := &entities.Company{}
	if err := json.NewDecoder(req.Body).Decode(&company); err != nil {
		internal.RespondWithJSON(w, http.StatusBadRequest, fmt.Errorf("unable to add the company, invalid request payload received: %v", err))
		return
	}
	if err := clients.DbClient().GetCompanyRepository().AddCompany(company); err != nil {
		internal.RespondWithJSON(w, http.StatusBadRequest, fmt.Errorf("unable to add the company, something went wrong: %v", err))
		return
	}
	clients.ServiceBusClient().Publish(company) // if there is an error while publishing to the service bus, it will just be logged
	internal.RespondWithJSON(w, http.StatusAccepted, nil)
}

func RemoveCompany(clients services.ClientSet, w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	idStr := params["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		internal.RespondWithJSON(w, http.StatusBadRequest, err)
		return
	}
	if err := clients.DbClient().GetCompanyRepository().RemoveCompany(id); err != nil {
		internal.RespondWithJSON(w, http.StatusBadRequest, fmt.Errorf("unable to remove the company, something went wrong: %v", err))
		return
	}
	internal.RespondWithJSON(w, http.StatusOK, nil)
}
