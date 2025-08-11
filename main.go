package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// MySQL connection string format: username:password@tcp(host:port)/dbname
	dsn := "root:Akashdeep@1508@tcp(127.0.0.1:3306)/startersql"

	// Open connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	defer db.Close()

	// Test connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}
	fmt.Println("Connected to MySQL successfully!")

	// Example query
	rows, err := db.Query("SELECT id,name FROM users")
	if err != nil {
		log.Fatal("Query error:", err)
	}
	defer rows.Close()

	// Loop through results
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal("Row scan error:", err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}
}
