package models

import (
	"log"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// DB is a global variable for accessing the DB
var DB *pg.DB

// InitDB intialises the postgres DB and stores it in a global variable
func InitDB() {
	DB = pg.Connect(&pg.Options{
		Database: "crypto_compare",
		User:     "pocockn",
		Password: "only8deb",
	})

	err := createSchema()
	if err != nil {
		panic(err)
	}
}

func createSchema() error {
	err := DB.CreateTable(&Wallet{}, &orm.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		return err
	}
	log.Println("Database schema created")
	return nil
}
