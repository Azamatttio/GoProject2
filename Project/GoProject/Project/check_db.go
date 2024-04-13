package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// For windows use postgresql:// instead of postgres:// in the connection string first part
	db, err := sql.Open("postgres", "postgresql://postgres:Almira_24@localhost:5432/mydatabase?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to DB!")
}
