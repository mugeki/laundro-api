package geolocation_test

import (
	"bytes"
	"io/ioutil"
	geoBusiness "laundro-api-ca/business/geolocation"
	geolocation "laundro-api-ca/drivers/thirdparties/ipapi"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

var (
	geoRepo geoBusiness.Repository
	geoClient *MockClient
	MockDo func(req *http.Request) (*http.Response, error)
)

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return MockDo(req)
}

func TestMain(m *testing.M){
	geoRepo = geolocation.NewIpAPI()
	geoClient = &MockClient{}
	m.Run()
}

func TestGetLocationByIP(t *testing.T) {
	geoResponse := `{
		"ip": "8.8.8.8",
		"version": "IPv4",
		"city": "Mountain View",
		"region": "California",
		"region_code": "CA",
		"country": "US",
		"country_name": "United States",
		"country_code": "US",
		"country_code_iso3": "USA",
		"country_capital": "Washington",
		"country_tld": ".us",
		"continent_code": "NA",
		"in_eu": false,
		"postal": "Sign up to access",
		"latitude": "Sign up to access",
		"longitude": "Sign up to access",
		"timezone": "America/Los_Angeles",
		"utc_offset": "-0700",
		"country_calling_code": "+1",
		"currency": "USD",
		"currency_name": "Dollar",
		"languages": "en-US,es-US,haw,fr",
		"country_area": 9629091.0,
		"country_population": 327167434.0,
		"message": "Please message us at ipapi.co/trial for full access",
		"asn": "AS15169",
		"org": "GOOGLE"
	}`

	res := ioutil.NopCloser(bytes.NewReader([]byte(geoResponse)))
	geoClient.DoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body: res,
		}, nil
	}

	resp, err := geoRepo.GetLocationByIP("8.8.8.8")
	
	assert.Nil(t, err)
	assert.Equal(t, "Mountain View", resp.City)
}