package models

import (
	"github.com/go-pg/pg"
)

var db *pg.DB

// InitDB intialises the postgres DB and stores it in a global variable
func InitDB() {
	db = pg.Connect(&pg.Options{
		Database: "crypto_compare",
		User:     "pocockn",
		Password: "only8deb",
	})
}
