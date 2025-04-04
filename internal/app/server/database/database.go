// internal/database/database.go
package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	// Setup PostgreSQL connection string
	connStr := "user=postgres password=yourpassword dbname=ecommerce sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	// Check the connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("Error pinging the database: ", err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")
}
