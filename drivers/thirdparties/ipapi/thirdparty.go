package geolocation

import (
	"encoding/json"
	"laundro-api-ca/business/geolocation"
	"net/http"
)

type IpAPI struct {
	httpClient http.Client
}

func NewIpAPI() geolocation.Repository{
	return &IpAPI{
		httpClient: http.Client{},
	}
}

func (geo *IpAPI) GetLocationByIP() (geolocation.Domain, error){
	req, _ := http.NewRequest("GET", "https://ipapi.co/json/", nil)
	req.Header.Set("User-Agent", "ipapi.co/#go-v1.3")
	resp, err := geo.httpClient.Do(req)
	if err != nil {
		return geolocation.Domain{}, err
	}

	defer resp.Body.Close()

	data := Response{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return geolocation.Domain{}, err
	}
	return data.toDomain(), nil
}