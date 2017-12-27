package main

// echo similar to Ratpack, lightweight HTTP program

/*
* To import a package solely for its side-effects (initialization),
* the underscore import is used for the side-effect of
* registering the sqlite3 driver as a database driver
* in the init() function, without importing any other functions:
 */

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/pocockn/crypto-compare-go/handlers"

	"github.com/pocockn/crypto-compare-go/api"
	"github.com/pocockn/crypto-compare-go/wallet"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
)

// CreateWallet initialises a user wallet based on one coin and some initial units
func createWallet(c echo.Context) error {
	coinMap := make(map[string]int)
	coin := c.QueryParam("coin")
	units, err := strconv.Atoi(c.QueryParam("units"))
	if err != nil {
		fmt.Println("error creating wallet")
	}
	coinMap[coin] = units
	btcWallet := wallet.NewWallet(coinMap)
	return c.JSON(http.StatusCreated, btcWallet)
}

func depositFunds(c echo.Context) error {
	// for now lets just create a new wallet
	// Next step is to find the wallet based on the ID and deposit to it
	coinMap := make(map[string]int)
	coin := c.QueryParam("coin")
	units, err := strconv.Atoi(c.QueryParam("units"))
	if err != nil {
		fmt.Println("error creating wallet")
	}
	coinMap[coin] = units
	btcWallet := wallet.NewWallet(coinMap)
	btcWallet.Deposit(coin, units)
	return c.JSON(http.StatusCreated, btcWallet)
}

func withdrawFunds(c echo.Context) error {
	// for now lets just create a new wallet with a base amount
	// And withdraw from that amount
	coinMap := make(map[string]int)
	coin := "BTC"
	units := 100
	coinMap[coin] = units
	btcWallet := wallet.NewWallet(coinMap)
	coinFromQuery := c.QueryParam("coin")
	units, err := strconv.Atoi(c.QueryParam("units"))
	if err != nil {
		return c.Render(http.StatusBadRequest, "Bad request", units)
	}
	btcWallet.Withdraw(coinFromQuery, units)
	return c.JSON(http.StatusOK, btcWallet)
}

func main() {

	// Create a new instance of Echo
	// e := is short hand for var e =
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS default
	// Allows requests from any origin with GET,HEAD,PUT,POST,DELETE
	// e.Use(middleware.CORS())

	// CORS restricted
	// allow request from any localhost address
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5000"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// Fetchs a list of coins from the cryptocompare API
	e.GET("/", func(context echo.Context) error {
		return context.JSON(http.StatusOK, api.FetchCoinList())
	})

	e.GET("/html", func(context echo.Context) error {
		return context.File("public/index.html")
	})

	e.GET("/createWallet", createWallet)

	e.POST("/deposit", depositFunds)

	e.POST("/withdraw", withdrawFunds)

	e.GET("/wallet", func(context echo.Context) error {
		return context.JSON(http.StatusOK, nil)
	})

	e.GET("/coin", handlers.GetCoin)

	// Similar to Ratpack handler, route takes a pattern and then
	// a handler function as param
	// e.GET("/tasks", handlers.GetTasks(db))

	// // In GOLang the type comes after variable
	// // eg context is of type echo.Context
	// e.PUT("/tasks", handlers.PutTasks(db))

	// e.DELETE("/tasks/:id", handlers.DeleteTask(db))

	// Start the web server
	e.Start(":8000")

}
