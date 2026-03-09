package main

import (
	"database/sql"
	"log"

	"github.com/yeboahd24/hexagonal/adapters/http"
	"github.com/yeboahd24/hexagonal/adapters/postgres"
	"github.com/yeboahd24/hexagonal/hexagonal/internal/adapters/httpadapter"
	_ "github.com/yeboahd24/hexagonal/internal/adapters/httpadapter"
	"github.com/yeboahd24/hexagonal/internal/application"
)

func setupDatabase() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=hexagonal sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func main() {
	db := setupDatabase()

	repo := postgres.NewUserRepository(db)

	service := application.NewUserService(repo)
	handler := httpadapter.NewUserHandler(service)

	http.HandleFunc("/users", handler.CreateUser)

	http.ListenAndServe(":8080", nil)
}
