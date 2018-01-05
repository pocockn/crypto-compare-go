package api

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock the data that the API brings back and check it gets unmarshalled into the correct format
// Price changes every second, the only way for a stable test

func TestFetchCoinList(t *testing.T) {
	jsonData := []byte(` {
		"Response": "Success",
		"Message": "Coin list succesfully returned! This api is moving to https://min-api.cryptocompare.com/data/all/coinlist, please change the path.",
		"BaseImageUrl": "https://www.cryptocompare.com",
		"BaseLinkUrl": "https://www.cryptocompare.com",
		"DefaultWatchlist": {
		"CoinIs": "1182,7605,5038,24854,3807,3808,202330,5324,5031,20131",
		"Sponsored": ""
		},
		"Data": {
			"42": {
			"Id": "4321",
			"Url": "/coins/42/overview",
			"ImageUrl": "/media/12318415/42.png",
			"Name": "42",
			"Symbol": "42",
			"CoinName": "42 Coin",
			"FullName": "42 Coin (42)",
			"Algorithm": "Scrypt",
			"ProofType": "PoW/PoS",
			"FullyPremined": "0",
			"TotalCoinSupply": "42",
			"PreMinedValue": "N/A",
			"TotalCoinsFreeFloat": "N/A",
			"SortOrder": "34",
			"Sponsored": false
			}
		}
		}`)

	response := ResponseCryptoCompare{}

	expected := Coin{
		ID:                  "4321",
		URL:                 "/coins/42/overview",
		ImageURL:            "/media/12318415/42.png",
		Name:                "42",
		Symbol:              "42",
		CoinName:            "42 Coin",
		FullName:            "42 Coin (42)",
		Algorithm:           "Scrypt",
		ProofType:           "PoW/PoS",
		FullyPremined:       "0",
		TotalCoinSupply:     "42",
		PreMinedValue:       "N/A",
		TotalCoinsFreeFloat: "N/A",
		SortOrder:           "34",
		Sponsored:           false,
	}

	json.Unmarshal(jsonData, &response)

	assert.Equal(t, response.Data["42"], expected)
}

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

	responseMarket := ResponseCoinMarketCap{}

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
