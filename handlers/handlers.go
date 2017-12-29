package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/pocockn/crypto-compare-go/models"
)

// GetCoin handler that queries the Crypto Compare API for a specific coin
// Takes a crypto symbol and currency symbol
// EG: fsym=BTC tsym=USD
func GetCoin(context echo.Context) error {
	// get the crypto symbol and fiat symbol that is sent through the url
	fsym := context.QueryParam("fsym")
	tsym := context.QueryParam("tsym")
	return context.String(http.StatusOK, fsym+tsym)
}

// CreateWallet initialises a user wallet based on one coin and some initial units
func CreateWallet(c echo.Context) error {
	coinMap := make(map[string]int)
	coin := c.QueryParam("coin")
	units, err := strconv.Atoi(c.QueryParam("units"))
	if err != nil {
		fmt.Println("error creating wallet")
	}
	coinMap[coin] = units
	btcWallet := models.NewWallet(coinMap)
	return c.JSON(http.StatusCreated, btcWallet)
}
