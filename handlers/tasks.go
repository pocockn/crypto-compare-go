package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/pocockn/crypto-compare-go/models"
)

// allows us to return an arbitary JSON response
// A map, with strings as keys and anything as values
// interface represents anything, similar to generics?
type jsonEntity map[string]interface{}

// GetTasks endpoint
func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		// Fetch the tasks using our new model
		return context.JSON(http.StatusOK, models.GetTasks(db))
	}
}

// PutTasks endpoint
func PutTasks(db *sql.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		// Instantiate a new Task
		var task models.Task
		// Map the incoming JSON body to the new task
		context.Bind(&task)
		id, err := models.PutTask(db, task.Name)
		if err == nil {
			return context.JSON(http.StatusCreated, jsonEntity{
				"created": id,
			})
		} else {
			return err
		}
	}
}

// DeleteTasks endpoint
func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(context echo.Context) error {
		// Casts our ID to an int, just to be sure.
		id, _ := strconv.Atoi(context.Param("id"))
		_, err := models.DeleteTask(db, id)
		if err == nil {
			return context.JSON(http.StatusOK, jsonEntity{
				"deleted": id,
			})
		} else {
			return err
		}
	}
}
