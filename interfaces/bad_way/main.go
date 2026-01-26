package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// What sucks here (concrete)

// 1. Config scattered and coupled
// 2. Business logic directly uses sql.DB
// 3. Need to touch every file (even the config) to change the database
// 4. Need a real Postgres instance (for testing)

func main() {
	// config scattered and coupled
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// business logic directly uses sql.DB
	err = CreateUser(db, "Dominic", "dominic@example.com")
	if err != nil {
		log.Fatal(err)
	}
}
