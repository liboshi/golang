package soap

import (
	"fmt"
	"io"
	"net/http"
)

func NewRequest(url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set(`Content-Type`, `text/xml; charset="utf-8"`)
	soapAction := fmt.Sprintf("%s/%s", "", "")
	req.Header.Set(`SOAPAction`, soapAction)

	return req, nil
}
