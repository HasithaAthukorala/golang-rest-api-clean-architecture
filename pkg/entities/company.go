package entities

type Company struct {
	Id      int    `json:"id" valid:"Required" schema:"id"`
	Name    string `json:"name" valid:"Required" schema:"name"`
	Country string `json:"country" valid:"Required" schema:"country"`
	Code    string `json:"code" valid:"Required" schema:"code"`
	Website string `json:"website" valid:"Required" schema:"website"`
	Phone   string `json:"phone" valid:"Required" schema:"phone"`
}

type IpResponse struct {
	CountryCode string `json:"countryCode" valid:"Required"`
}
