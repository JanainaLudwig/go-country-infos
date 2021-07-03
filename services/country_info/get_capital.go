package country_info

import (
	"encoding/xml"
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

func (c * CountryInfo) GetCapital(countryCode string) (string, error) {
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
		return "", err
	}

	res := CapitalResponse{}
	err = action.Do(&res)
	if err != nil {
		return "", err
	}

	return res.Body.CapitalCityResponse.CapitalCityResult, nil
}
