package main

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"crypto-compare-go/handlers"
	"crypto-compare-go/models"

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
	models.InitDB("crypto_compare")

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
		return context.File("public/index.html")
	})

	e.POST("/createWallet", handlers.CreateWallet)

	e.GET("/deposit/:id", func(context echo.Context) error {
		return context.Render(http.StatusOK, "deposit.html", nil)
	})
	e.POST("/deposit/:id", handlers.DepositCoin)

	e.GET("/withdraw/:id", func(context echo.Context) error {
		return context.Render(http.StatusOK, "withdraw.html", nil)
	})
	e.POST("/withdraw/:id", handlers.WithdrawCoin)

	e.GET("/wallets", func(context echo.Context) error {
		log.Printf("Returning all wallets")
		wallets, err := models.AllWallets()
		if err != nil {
			panic(err)
		}
		return context.Render(http.StatusOK, "wallets.html", wallets)
	})
	e.GET("/wallet/:id", handlers.GetWallet)

	// Fetchs a list of coins from the cryptocompare API
	e.GET("/allCoins", func(context echo.Context) error {
		return context.JSON(http.StatusOK, models.FetchCoinList())
	})

	e.GET("/coin", handlers.GetCoin)

	// Start the web server
	e.Start(":8000")

}
