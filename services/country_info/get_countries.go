package country_info

import (
	"encoding/xml"
	"log"
	"strings"
)

type CountriesResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soap    string   `xml:"soap,attr"`
	Body    struct {
		Text                             string `xml:",chardata"`
		ListOfCountryNamesByCodeResponse struct {
			Text                           string `xml:",chardata"`
			Xmlns                          string `xml:"xmlns,attr"`
			ListOfCountryNamesByCodeResult struct {
				Text                string `xml:",chardata"`
				TCountryCodeAndName []struct {
					Text     string `xml:",chardata"`
					SISOCode string `xml:"sISOCode"`
					SName    string `xml:"sName"`
				} `xml:"tCountryCodeAndName"`
			} `xml:"ListOfCountryNamesByCodeResult"`
		} `xml:"ListOfCountryNamesByCodeResponse"`
	} `xml:"Body"`
}

type Country struct {
	Name string
	Code string
}

func (c * CountryInfo) GetCountries() (string, error) {
	payload := []byte(strings.TrimSpace(`
    <soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	  <soap:Body>
		<ListOfCountryNamesByCode xmlns="http://www.oorsprong.org/websamples.countryinfo">
		</ListOfCountryNamesByCode>
	  </soap:Body>
	</soap:Envelope>`,
	))

	action, err := c.NewAction("ListOfCountryNamesByCode", payload)
	if err != nil {
		return "", err
	}

	res := CountriesResponse{}
	err = action.Do(&res)
	if err != nil {
		return "", err
	}

	var countries []Country
	for _, val := range res.Body.ListOfCountryNamesByCodeResponse.ListOfCountryNamesByCodeResult.TCountryCodeAndName {
		countries = append(countries, Country{
			Name: val.SName,
			Code: val.SISOCode,
		})
	}
	log.Println(countries)
	return "", nil
}
