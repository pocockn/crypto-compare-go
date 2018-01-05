package persistance

import (
	"log"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// DB is a global variable for accessing the DB
var DB *pg.DB

// InitDB intialises the postgres DB and stores it in a global variable
func InitDB(dbName string) {
	DB = pg.Connect(&pg.Options{
		Database: dbName,
		User:     "pocockn",
		Password: "only8deb",
	})

}

func CreateSchema(model interface{}) error {
	err := DB.CreateTable(model, &orm.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		return err
	}
	log.Println("Database schema created")
	return nil
}
