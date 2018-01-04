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

func TestTopCoinList(t *testing.T) {
	jsonData := []byte(`
		[
			{
				"id": "bitcoin",
				"name": "Bitcoin",
				"symbol": "BTC",
				"rank": "1",
				"price_usd": "14832.6",
				"price_btc": "1.0",
				"24h_volume_usd": "19144300000.0",
				"market_cap_usd": "248910117556",
				"available_supply": "16781287.0",
				"total_supply": "16781287.0",
				"max_supply": "21000000.0",
				"percent_change_1h": "-1.11",
				"percent_change_24h": "-3.55",
				"percent_change_7d": "2.05",
				"last_updated": "1515055160"
			}
		]
	`)

	responseMarket := ResponseMarket{}

	response := CoinMarketCapCoin{
		ID:               "bitcoin",
		Name:             "Bitcoin",
		Symbol:           "BTC",
		Rank:             "1",
		PriceUsd:         "14832.6",
		PriceBtc:         "1.0",
		Two4HVolumeUsd:   "19144300000.0",
		MarketCapUsd:     "248910117556",
		AvailableSupply:  "16781287.0",
		TotalSupply:      "16781287.0",
		MaxSupply:        "21000000.0",
		PercentChange1H:  "-1.11",
		PercentChange24H: "-3.55",
		PercentChange7D:  "2.05",
		LastUpdated:      "1515055160",
	}

	json.Unmarshal(jsonData, &responseMarket.Coins)
	assert.Equal(t, response, responseMarket.Coins[0])
}
