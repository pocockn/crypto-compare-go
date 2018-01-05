package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// FetchCoinList returns a list of coins from the Cryptocompare Api
func FetchCoinList() map[string]Coin {
	url := fmt.Sprintf("https://www.cryptocompare.com/api/data/coinlist/")

	fmt.Println("Requesting data from " + url)

	response, err := makeRequest(url)
	if err != nil {
		log.Fatal(err.Error())
	}

	body, err := readResponseBody(response)
	if err != nil {
		log.Fatal(err.Error())
	}

	apiModel := ResponseCryptoCompare{}

	// unmarshall json bytes into Response struct
	json.Unmarshal(body, &apiModel)

	return apiModel.Data
}

// FetchCoinPrice gets the price in GBP for one unit of the specified coin
func FetchCoinPrice(fsym string) map[string]float64 {
	url := fmt.Sprintf("https://min-api.cryptocompare.com/data/price?fsym=" + fsym + "&tsyms=GBP")

	fmt.Println("Requesting data from " + url)

	response, err := makeRequest(url)
	if err != nil {
		log.Fatal(err.Error())
	}

	body, err := readResponseBody(response)
	if err != nil {
		log.Fatal(err.Error())
	}

	priceMap := map[string]float64{}
	json.Unmarshal(body, &priceMap)

	return priceMap
}

// FetchTopCoins fetchs the top 10 coins via the Coin Market Cap API, used to display list on the home page
func FetchTopCoins() ResponseCoinMarketCap {
	url := fmt.Sprintf("https://api.coinmarketcap.com/v1/ticker/?limit=10")

	fmt.Println("Requesting data from " + url)

	response, err := makeRequest(url)
	if err != nil {
		log.Fatal(err.Error())
	}

	body, err := readResponseBody(response)
	if err != nil {
		log.Fatal(err.Error())
	}

	priceMap := ResponseCoinMarketCap{}

	err = json.Unmarshal(body, &priceMap.Coins)
	if err != nil {
		log.Fatal(err.Error())
	}

	return priceMap
}

// Make the HTTP GET request and return a response
func makeRequest(url string) (*http.Response, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Read the response body and return it if we get a 200
func readResponseBody(response *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		return body, nil
	}
	return nil, err
}
