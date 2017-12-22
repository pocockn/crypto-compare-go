package handlers

import (
	"net/http"

	"github.com/labstack/echo"
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
