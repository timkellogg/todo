package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// Store : pointer to application database
var Store *sql.DB

// InitializeStore : loads database connection (Store)
func InitializeStore() {
	var err error
	Store, err = sql.Open("postgres", "user=tkellogg host=localhost dbname=todo_go sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	buildTables()
	log.Println("initialized Store")
}

// buildTables : creates database tables if they don't exist
func buildTables() {
	cmd := "CREATE TABLE IF NOT EXISTS Todos (id serial PRIMARY KEY, name varchar(64) NOT NULL, completed bool, due timestamp);"
	_, err := Store.Exec(cmd)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("built tables")
}
