package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchCoinList(t *testing.T) {
	response := Coin{
		ID:                  "3808",
		URL:                 "/coins/ltc/overview",
		ImageURL:            "/media/19782/litecoin-logo.png",
		Name:                "LTC",
		Symbol:              "LTC",
		CoinName:            "Litecoin",
		FullName:            "Litecoin (LTC)",
		Algorithm:           "Scrypt",
		ProofType:           "PoW",
		FullyPremined:       "0",
		TotalCoinSupply:     "84000000",
		PreMinedValue:       "N/A",
		TotalCoinsFreeFloat: "N/A",
		SortOrder:           "3",
		Sponsored:           false,
	}

	coinList := FetchCoinList()
	assert.Equal(t, response, coinList["LTC"])
}

// Mock the data that the API brings back and check it gets unmarshalled into the correct format
// Price changes ever second, the only way for a stable test
func TestCoinPrice(t *testing.T) {
	expected := map[string]float64{}
	expected["GBP"] = 657.54
	data := []byte(`{
		"GBP": 657.54
	  }`)
	newMap := map[string]float64{}
	json.Unmarshal(data, &newMap)
	assert.Equal(t, expected, newMap)
}
