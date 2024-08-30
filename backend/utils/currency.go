package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func GetExchangeAmount(from, to string, amount float64) (float64, error) {
	url := "https://tw.rter.info/capi.php"

	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("failed to get exchange rate api: %w", err)
	}
	defer resp.Body.Close()

	var data map[string]struct {
		Exrate float64 `json:"Exrate"`
		Date   string  `json:"Date"`
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return 0, fmt.Errorf("failed to decode response body: %w", err)
	}

	ex, ok := data[from+to]
	if ok {
		return ex.Exrate * amount, nil
	}

	ex, ok = data[to+from]
	if ok {
		return amount / ex.Exrate, nil
	}

	return 0, errors.New("exchange rate not found")
}
