package api

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/EvegeniyNekrasov/go/cryptomasters/datatypes"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"


// we are reciving a pointer to datatypes.Rate so we can 
// return a nil if there is an error
func GetRate(currency string) (*datatypes.Rate, error) {
	res, err := http.Get(fmt.Sprintf(apiUrl, strings.ToUpper(currency)))
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusOK {
		data, err := io .ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		json := string(data)

		fmt.Println(json)
	} else {
		return nil, fmt.Errorf("status code received: %v", res.StatusCode)
	}
	rate := datatypes.Rate{Currency: currency, Price: 20}

	return &rate, nil
}