package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/liboshi/golang/methods"
	"github.com/liboshi/golang/soap"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	v := &soap.Envelope{}
	v.Body = &soap.Body{}
	v.Body.Content = &methods.GetCitiesByCountry{CountryName: "China"}
	output, err := xml.Marshal(v)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	rawReqBody := io.MultiReader(strings.NewReader(xml.Header), bytes.NewReader(output))

	url := "http://www.webservicex.net/globalweather.asmx"
	action := "GetCitiesByCountry"
	req, err := http.NewRequest("POST", url, rawReqBody)
	if err != nil {
		return
	}
	req.Header.Set(`Host`, "www.webservicex.net")
	req.Header.Set(`Content-Type`, `text/xml; charset="utf-8"`)
	soapAction := fmt.Sprintf("%s/%s", "http://www.webserviceX.NET", action)
	req.Header.Set(`SOAPAction`, soapAction)

	client := &http.Client{}
	resp, _ := client.Do(req)

	_, errr := io.Copy(os.Stdout, resp.Body)
	if errr != nil {
		fmt.Println(errr)
	}
}
