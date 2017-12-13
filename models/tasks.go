package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Task is a structure containing Task Data
type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// TaskCollection is a collection of tasks
type TaskCollection struct {
	Tasks []Task `json:"items"`
}

func GetTasks(db *sql.DB) TaskCollection {
	sql := "SELECT * FROM tasks"
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	result := TaskCollection{}
	for rows.Next() {
		task := Task{}
		err2 := rows.Scan(&task.ID, &task.Name)

		if err2 != nil {
			panic(err2)
		}

		result.Tasks = append(result.Tasks, task)
	}
	return result
}

func PutTask(db *sql.DB, name string) (int64, error) {
	sql := "INSERT INTO tasks(name) VALUES(?)"

	statement, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	// call immediately before the function executing returns
	defer statement.Close()

	// replace ? with the variable in our query
	result, err2 := statement.Exec(name)

	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

func DeleteTask(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM tasks WHERE id = ?"

	statement, error := db.Prepare(sql)

	isError(error)

	result, err := statement.Exec(id)

	isError(err)

	return result.RowsAffected()
}

func isError(inputError error) {
	if inputError != nil {
		panic(inputError)
	}
}
