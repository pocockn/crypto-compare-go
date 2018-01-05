package main

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/crypto-compare-go/api"
	"github.com/crypto-compare-go/handlers"
	"github.com/crypto-compare-go/persistance"
	"github.com/crypto-compare-go/wallet"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	persistance.InitDB("crypto_compare")
	err := persistance.CreateSchema(&wallet.Wallet{})

	if err != nil {
		panic(err)
	}

	// Create a new instance of Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("public/*.html")),
	}
	e.Renderer = renderer

	// CORS restricted
	// allow request from any localhost address
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5000"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.GET("/", func(context echo.Context) error {
		return context.Render(http.StatusOK, "index.html", api.FetchCoinList())
	})

	e.POST("/createWallet", handlers.CreateWallet)

	e.GET("/deposit/:id", func(context echo.Context) error {
		return context.Render(http.StatusOK, "deposit.html", api.FetchCoinList())
	})
	e.POST("/deposit/:id", handlers.DepositCoin)

	e.GET("/withdraw/:id", func(context echo.Context) error {
		wallet, err := wallet.GetWallet(context.Param("id"))
		if err != nil {
			panic(err)
		}
		return context.Render(http.StatusOK, "withdraw.html", wallet)
	})
	e.POST("/withdraw/:id", handlers.WithdrawCoin)
	e.POST("/delete/:id", handlers.DeleteWallet)

	e.GET("get-price/:symbol", handlers.GetPrice)

	e.GET("/wallets", func(context echo.Context) error {
		log.Printf("Returning all wallets")
		wallets, err := wallet.AllWallets()
		if err != nil {
			panic(err)
		}
		return context.Render(http.StatusOK, "wallets.html", wallets)
	})
	e.GET("/wallet/:id", handlers.GetWallet)

	e.GET("/coins", func(context echo.Context) error {
		return context.Render(http.StatusOK, "home.html", api.FetchTopCoins())
	})

	// Fetchs a list of coins from the cryptocompare API
	e.GET("/allCoins", func(context echo.Context) error {
		return context.JSON(http.StatusOK, api.FetchCoinList())
	})

	e.GET("/coin", handlers.GetCoin)

	// Start the web server
	e.Start(":8000")

}
