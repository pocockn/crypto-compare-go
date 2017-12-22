package main

// echo similar to Ratpack, lightweight HTTP program

/*
* To import a package solely for its side-effects (initialization),
* the underscore import is used for the side-effect of
* registering the sqlite3 driver as a database driver
* in the init() function, without importing any other functions:
 */

import (
	"crypto-compare-go/api"
	"crypto-compare-go/handlers"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
)

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

	//This is a static html file that will contain our VueJS client code.
	// We can serve up static files using the 'File' function.
	e.GET("/", func(context echo.Context) error {
		return context.JSON(http.StatusOK, api.FetchCoinList())
	})

	e.GET("/coin", handlers.GetCoin)

	// Use the handlers within our handler package

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
