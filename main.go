package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-pg/pg/orm"

	"github.com/pocockn/crypto-compare-go/handlers"
	"github.com/pocockn/crypto-compare-go/models"

	"github.com/pocockn/crypto-compare-go/api"
	"github.com/pocockn/crypto-compare-go/wallet"

	"math/rand"

	"github.com/go-pg/pg"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	models.InitDB()
	err := createSchema(models.DB)

	if err != nil {
		panic(err)
	}

	// Create a new instance of Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

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
		var wallets []wallet.Wallet
		err = models.DB.Model(&wallets).Select()
		if err != nil {
			panic(err)
		}
		return context.JSON(200, wallets)
	})

	e.GET("/coin", handlers.GetCoin)

	// Start the web server
	e.Start(":8000")

}

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
	db := pg.Connect(&pg.Options{
		Database: "crypto_compare",
		User:     "pocockn",
		Password: "only8deb",
	})
	wallet := &wallet.Wallet{
		ID:        rand.Int(),
		CoinsHeld: coinMap,
	}
	errDb := db.Insert(wallet)
	if errDb != nil {
		panic(errDb)
	}
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

func createSchema(db *pg.DB) error {
	err := db.CreateTable(&wallet.Wallet{}, &orm.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		return err
	}
	log.Println("Database schema created")
	return nil
}
