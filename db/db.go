package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

//go-sqlite3 package isnt directly used,but go uses it under the hood as we interact with the built in sql package part of go's std library. We append _ to it which tells go we need that import, although we dont use it directly, but it exposes functionality that is used under the hood by the built in sql package
//sqlite3 is a driver

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("COuld not connect to the database") //Raising a panic doesnt crash the server but we get a log(message) highlighting the issue
	}

	DB.SetMaxOpenConns(10) //SetMaxOpenConns sets the maximum number of open connections to the database.
	DB.SetMaxIdleConns(5)  //SetMaxIdleConns sets the maximum number of connections in the idle connection pool.

	// Open connections are the connections that are actively being used to execute queries or transactions. Setting a limit helps prevent your application from opening too many connections, which could lead to resource exhaustion.

	// Idle connections are those that are not currently being used but are kept open in the connection pool for potential reuse. The connection pool maintains a certain number of idle connections to reduce the overhead of opening and closing connections for each database operation.

	createTables()

}

func createTables() {
	//Multi line string: CREATE TABLE statement
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time DATETIME NOT NULL,
		user_id INTEGER
	)
	`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic("Events rable creation failed:(")
	}
}
