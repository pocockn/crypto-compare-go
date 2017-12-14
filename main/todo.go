package main

// echo similar to Ratpack, lightweight HTTP program

/*
* To import a package solely for its side-effects (initialization),
* the underscore import is used for the side-effect of
* registering the sqlite3 driver as a database driver
* in the init() function, without importing any other functions:
 */

import (
	"database/sql"

	"crypto-compare-go/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db := initDB("storage.db")
	migrate(db)

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
	e.File("/", "public/index.html")

	// Use the handlers within our handler package

	// Similar to Ratpack handler, route takes a pattern and then
	// a handler function as param
	e.GET("/tasks", handlers.GetTasks(db))

	// In GOLang the type comes after variable
	// eg context is of type echo.Context
	e.PUT("/tasks", handlers.PutTasks(db))

	e.DELETE("/tasks/:id", handlers.DeleteTask(db))

	// Start the web server
	e.Start(":8000")

}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	// check if any db errors then exit
	if err != nil {
		panic(err)
	}

	// if we dont get any errors
	// but still no db connection, exit as well
	if db == nil {
		panic("db nill")
	}

	return db
}

func migrate(db *sql.DB) {
	sql := `
		CREATE TABLE IF NOT EXISTS tasks(
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name VARCHAR NOT NULL
	);
	`

	// _ is an ignored value, we only want to see if theres an error
	// so we ignore the first return value, which is the result and store error
	// in a variable.
	_, err := db.Exec(sql)

	// Exit if something went wrong
	if err != nil {
		panic(err)
	}
}
