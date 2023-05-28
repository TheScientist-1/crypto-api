package bitcoin

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// BitcoinService відповідає за взаємодію з API криптовалютної біржі
type BitcoinService struct {
	APIURL string
	// Інші поля, якщо потрібно
}

// NewBitcoinService створює новий екземпляр BitcoinService
func NewBitcoinService() *BitcoinService {
	return &BitcoinService{
		APIURL: "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=uah",
		
	}
}

// GetCurrentPrice виконує GET-запит до API та повертає поточний курс біткоїна
func (s *BitcoinService) GetCurrentPrice() (float64, error) {
	resp, err := http.Get(s.APIURL)
	if err != nil {
		return 0, fmt.Errorf("помилка під час виконання GET-запиту: %s", err.Error())
	}
	defer resp.Body.Close()

	var data struct {
		Price float64 `json:"price"`
		// Інші поля, якщо є
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return 0, fmt.Errorf("помилка під час розпарсування JSON-даних: %s", err.Error())
	}

	return data.Price, nil
}
