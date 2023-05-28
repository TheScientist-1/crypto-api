package bitcoin

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BitcoinService struct {
	APIURL string
}

// NewBitcoinService creates a new instance of BitcoinService
func NewBitcoinService() *BitcoinService {
	return &BitcoinService{
		APIURL: "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=uah",
		
	}
}

// GetCurrentPrice executes a GET request to the API and returns the current bitcoin rate
func (s *BitcoinService) GetCurrentPrice() (float64, error) {
	resp, err := http.Get(s.APIURL)
	if err != nil {
		return 0, fmt.Errorf("GET request failed: %s", err.Error())
	}
	defer resp.Body.Close()

	var data struct {
		Price float64 `json:"price"`
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return 0, fmt.Errorf("error parsing JSON data: %s", err.Error())
	}

	return data.Price, nil
}
