package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver anonymously
)

// DB is the global database connection pool
var DB *sql.DB

// DBconnect initializes the connection to the MySQL database
func DBconnect() {
	var err error

	// Define the Data Source Name (DSN)
	// Format: username:password@tcp(host:port)/databaseName?params
	// Note: Replace "localhost" with your actual MySQL container IP if needed
	dsn := "root:root@tcp(localhost:3306)/forum?charset=utf8mb4&parseTime=True&loc=Local"

	// Open a database connection using the MySQL driver
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		// If opening the connection fails, terminate the application
		panic(fmt.Sprintf("Error opening database connection: %v", err))
	}

	// Ping the database to test if the connection is alive
	err = DB.Ping()
	if err != nil {
		// If the ping fails, terminate the application
		panic(fmt.Sprintf("Unable to ping the database: %v", err))
	}

	// Log successful connection
	fmt.Println("Successfully connected to the database")
}
