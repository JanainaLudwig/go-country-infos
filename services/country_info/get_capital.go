package country_info

import (
	"encoding/xml"
	"log"
	"strings"
)

type CapitalResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"soap,attr"`
	Body    struct {
		Text                string `xml:",chardata"`
		CapitalCityResponse struct {
			Text              string `xml:",chardata"`
			M                 string `xml:"m,attr"`
			CapitalCityResult string `xml:"CapitalCityResult"`
		} `xml:"CapitalCityResponse"`
	} `xml:"Body"`
}


func (c * CountryInfo) GetCapital(countryCode string) error {
	payload := []byte(strings.TrimSpace(`
    <soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	  <soap:Body>
		<CapitalCity xmlns="http://www.oorsprong.org/websamples.countryinfo">
		  <sCountryISOCode>` + countryCode + `</sCountryISOCode>
		</CapitalCity>
	  </soap:Body>
	</soap:Envelope>`,
	))

	action, err := c.NewAction("CapitalCity", payload)
	if err != nil {
		return err
	}

	err = action.Do()
	if err != nil {
		return err
	}

	res := CapitalResponse{}
	err = action.Decode(&res)
	if err != nil {
		return err
	}

	log.Println(res.Body.CapitalCityResponse.CapitalCityResult)

	return nil
}

/*

func (c * CountryInfo) GetCapital(countryCode string) error {
	payload := []byte(strings.TrimSpace(`
    <soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	  <soap:Body>
		<CapitalCity xmlns="http://www.oorsprong.org/websamples.countryinfo">
		  <sCountryISOCode>` + countryCode + `</sCountryISOCode>
		</CapitalCity>
	  </soap:Body>
	</soap:Envelope>`,
	))

	req, err := http.NewRequest("POST", "http://webservices.oorsprong.org/websamples.countryinfo/CountryInfoService.wso", bytes.NewReader(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-type", "text/xml")
	req.Header.Set("SOAPAction", "CapitalCity")

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	log.Println(string(data))
	return nil
}

*/