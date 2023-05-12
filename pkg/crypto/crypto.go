package crypto

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func GetCryptoPrice(name string) (float64, error) {
	var crypto map[string]string
	url := fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%sUSDT", name)
	res, err := http.Get(url)
	if res.StatusCode != 200 {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}
	err = json.Unmarshal(body, &crypto)
	if err != nil {
		return 0, err
	}

	price, err := strconv.ParseFloat(crypto["price"], 64)
	if err != nil {
		return 0, err
	}

	return price, nil
}
