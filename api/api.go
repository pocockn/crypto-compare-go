package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pocockn/crypto-compare-go/models"
)

// Response represents the response from the CryptoCompare API
type Response struct {
	Response string
	Message  string
	Data     map[string]models.Coin
}

// FetchCoinList returns a list of coins from the Cryptocompare Api
func FetchCoinList() map[string]models.Coin {
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
