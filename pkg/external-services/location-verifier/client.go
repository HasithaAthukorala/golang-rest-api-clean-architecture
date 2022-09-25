package location_verifier

import (
	"encoding/json"
	"fmt"
	"golang-rest-api-clean-architecture/pkg/config"
	"golang-rest-api-clean-architecture/pkg/entities"
	"io/ioutil"
	"net/http"
	"net/url"
)

type LocationVerifier interface {
	Verify(ipAddress string) (bool, error)
}

type locationVerifier struct {
	httpClient *http.Client
	hostUrl    *url.URL
}

func New(cfg *config.Config) (LocationVerifier, error) {
	URL, err := url.Parse(cfg.LocationVerificationHostURL)
	if err != nil {
		return nil, fmt.Errorf("error occured while initializing location verifier: %v", err)
	}
	return &locationVerifier{
		hostUrl:    URL,
		httpClient: &http.Client{},
	}, nil
}

func (client *locationVerifier) Verify(ipAddress string) (bool, error) {
	endpoint := fmt.Sprintf("%s/json/%s", client.hostUrl, ipAddress)
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return false, err
	}
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return false, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("cannot get a valid response from ipapi.com for your ip address")
	}
	var ipResponse entities.IpResponse
	err = json.Unmarshal(respBody, &ipResponse)
	if err != nil {
		return false, err
	}
	if ipResponse.CountryCode == "CY" {
		return true, nil
	}
	return false, nil
}
