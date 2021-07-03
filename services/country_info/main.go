package country_info

import (
	"bytes"
	"encoding/xml"
	"net/http"
)

type CountryInfo struct {
	endpoint string
}

type CountryInfoAction struct {
	request *http.Request
	Response *http.Response
}

func NewCountryInfoClient() *CountryInfo {
	c := CountryInfo{
		endpoint: "http://webservices.oorsprong.org/websamples.countryinfo/CountryInfoService.wso",
	}

	return &c
}

func (c * CountryInfo) NewAction(action string, payload []byte) (*CountryInfoAction, error) {
	req, err := http.NewRequest("POST", c.endpoint, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-type", "text/xml")
	req.Header.Set("SOAPAction", action)

	return &CountryInfoAction{
		request: req,
	}, err
}

func (c *CountryInfoAction) Do() error {
	client := http.Client{}

	res, err := client.Do(c.request)
	if err != nil {
		return err
	}

	c.Response = res

	return err
}

func (c *CountryInfoAction) Decode(dest interface{}) error {
	err := xml.NewDecoder(c.Response.Body).Decode(dest)

	return err
}
