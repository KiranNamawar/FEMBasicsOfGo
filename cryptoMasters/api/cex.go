package api

import (
	"cryptomaster/types"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"fmt"
)

const apiEndpoint = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*types.Rate, error) {
	if len(currency) != 3 {
		return nil, fmt.Errorf("[%s] min 3 characters required for currency", currency)
	}
	res, err := http.Get(fmt.Sprintf(apiEndpoint, strings.ToUpper(currency)))
	if err != nil {
		return nil, err
	}
	var response CEXResponse
	if res.StatusCode == http.StatusOK {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("status code: %d", res.StatusCode)
	}
	rate := types.Rate{Currency: currency, Price: response.Bid}
	return &rate, nil
}
