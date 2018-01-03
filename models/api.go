package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Response represents the response from the CryptoCompare API
type Response struct {
	Response string
	Message  string
	Data     map[string]Coin
}

type Price struct {
	PriceWithSymbol map[string]string
}

type PriceList struct {
	priceList []Price
}

// FetchCoinList returns a list of coins from the Cryptocompare Api
func FetchCoinList() map[string]Coin {
	url := fmt.Sprintf("https://www.cryptocompare.com/api/data/coinlist/")

	fmt.Println("Requesting data from " + url)

	response := Response{}

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
		json.Unmarshal(body, &response)
	}
	return response.Data
}

// FetchCoinPrice gets the price in GBP for one unit of the specified coin
func FetchCoinPrice(fsym string) map[string]float64 {
	url := fmt.Sprintf("https://min-api.cryptocompare.com/data/price?fsym=" + fsym + "&tsyms=GBP")

	fmt.Println("Requesting data from " + url)

	priceMap := map[string]float64{}

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
		json.Unmarshal(body, &priceMap)
	}

	return priceMap

}
