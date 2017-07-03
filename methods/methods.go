package methods

import (
	"encoding/xml"
)

type GetCitiesByCountry struct {
	XMLName     xml.Name `xml:"http://www.webserviceX.NET GetCitiesByCountry"`
	CountryName string   `xml:"CountryName"`
}
