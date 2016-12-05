package methods

import (
	"encoding/xml"
)

type GetCitiesByCountry struct {
	XMLName     xml.Name `xml:"http://www.webservicex.net GetCitiesByCountry"`
	CountryName string   `xml:"CountryName"`
}
