package api

import (
	"crypto-compare-go/models"
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
	Data     map[string]models.Coin
}

func getCoinList() {
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
