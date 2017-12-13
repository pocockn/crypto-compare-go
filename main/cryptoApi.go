package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Coin represents a currency from the CrytoCompare API
type Coin struct {
	ID                  string `json:"Id"`
	URL                 string `json:"Url"`
	ImageURL            string `json:"ImageUrl"`
	Name                string `json:"Name"`
	Symbol              string `json:"Symbol"`
	CoinName            string `json:"CoinName"`
	FullName            string `json:"FullName"`
	Algorithm           string `json:"Algorithm"`
	ProofType           string `json:"ProofType"`
	FullyPremined       string `json:"FullyPremined"`
	TotalCoinSupply     string `json:"TotalCoinSupply"`
	PreMinedValue       string `json:"PreMinedValue"`
	TotalCoinsFreeFloat string `json:"TotalCoinsFreeFloat"`
	SortOrder           string `json:"SortOrder"`
	Sponsored           bool   `json:"Sponsored"`
}

// Response represents the response from the CryptoCompare API
type Response struct {
	Response string
	Message  string
	Data     map[string]Coin
}

func main() {
	url := fmt.Sprintf("https://www.cryptocompare.com/api/data/coinlist/")

	fmt.Println("Requesting data from " + url)

	resp := Response{}

	// getting the data using http
	request, err := http.Get(url)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Read the response body using ioutil
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer request.Body.Close()

	if request.StatusCode == http.StatusOK {
		// unmarshall json bytes into Response struct
		json.Unmarshal(body, &resp)

		fmt.Println(resp.Data)
	}
}
