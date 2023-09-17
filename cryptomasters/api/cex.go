package api

import (
	"encoding/json"
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
	if len(currency) != 3 {
		return nil, fmt.Errorf("3 characteres required; %d received", len(currency))
	}
	res, err := http.Get(fmt.Sprintf(apiUrl, strings.ToUpper(currency)))
	if err != nil {
		return nil, err
	}
	var response CEXResponse
	if res.StatusCode == http.StatusOK {
		data, err := io .ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		
		err = json.Unmarshal(data, &response)
		if err != nil {
			return nil, err
		}

		
	} else {
		return nil, fmt.Errorf("status code received: %v", res.StatusCode)
	}
	rate := datatypes.Rate{Currency: currency, Price: response.Bid}

	return &rate, nil
}