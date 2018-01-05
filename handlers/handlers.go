package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/pocockn/crypto-compare-go/api"
	"github.com/pocockn/crypto-compare-go/persistance"
	"github.com/pocockn/crypto-compare-go/wallet"

	"github.com/labstack/echo"
)

// GetCoin handler that queries the Crypto Compare API for a specific coin
func GetCoin(context echo.Context) error {
	fsym := context.QueryParam("fsym")
	tsym := context.QueryParam("tsym")
	return context.String(http.StatusOK, fsym+tsym)
}

// CreateWallet initialises a user wallet based on one coin and some initial units
func CreateWallet(c echo.Context) error {
	coinMap := make(map[string]int)
	coin := c.FormValue("coin")
	units, err := strconv.Atoi(c.FormValue("units"))
	if err != nil {
		panic(err)
	}

	coinMap[coin] = units
	wallet := &wallet.Wallet{
		ID:        rand.Int(),
		CoinsHeld: coinMap,
	}

	errDb := persistance.DB.Insert(wallet)
	if errDb != nil {
		panic(errDb)
	}
	return c.Redirect(301, "/wallets")
}

// GetWallet returns us a wallet from the DB via it's ID
func GetWallet(c echo.Context) error {
	id := c.Param("id")
	wallet, err := wallet.GetWallet(id)
	if err != nil {
		panic(err)
	}

	return c.Render(http.StatusOK, "wallet.html", wallet)
}

// DepositCoin adds a new coin to the specified wallet
func DepositCoin(c echo.Context) error {
	id := c.Param("id")
	coin := c.FormValue("coin")
	units, err := strconv.Atoi(c.FormValue("units"))

	if err != nil {
		panic(err)
	}

	wallet, err := wallet.GetWallet(id)
	wallet.CoinsHeld[coin] = units
	err = persistance.DB.Update(wallet)

	if err != nil {
		panic(err)
	}

	return c.Redirect(301, "/wallet/"+id)
}

// WithdrawCoin withdraws units based on a coins symbol
func WithdrawCoin(c echo.Context) error {
	id := c.Param("id")
	coin := c.FormValue("coin")
	units, err := strconv.Atoi(c.FormValue("units"))

	if err != nil {
		panic(err)
	}

	wallet, err := wallet.GetWallet(id)

	err = wallet.Withdraw(coin, units)

	if err != nil {
		return c.String(500, "Balance cannot be negative")
	}

	err = persistance.DB.Update(wallet)

	if err != nil {
		panic(err)
	}

	return c.Redirect(301, "/wallet/"+id)

}

// DeleteWallet deletes a wallet based on it's ID
func DeleteWallet(c echo.Context) error {
	id := c.Param("id")
	err := wallet.DeleteWallet(id)

	if err != nil {
		panic(err)
	}

	return c.Redirect(301, "/wallets")

}

func GetPrice(c echo.Context) error {
	symbol := c.Param("symbol")

	fmt.Println(symbol)

	price := api.FetchCoinPrice(symbol)

	fmt.Println(price)

	return c.JSON(200, price)
}
