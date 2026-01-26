package main

import (
	"context"
	"database/sql"
	"good-way/config"
	"good-way/db"
	"good-way/user"
	"log"
)

func main() {
	cfg := config.DatabaseConfig{
		Driver:   "postgres",
		Host:     "localhost",
		Port:     5432,
		User:     "app",
		Password: "secret",
		Name:     "app_db",
	}

	pg := db.NewPostgres(cfg)

	if err := pg.Connect(); err != nil {
		log.Fatal(err)
	}
	defer pg.Close()

	repo := user.NewPostgresRepository(pg.DB().(*sql.DB))
	service := user.NewService(repo)

	if err := service.Create(context.Background(), "Dominic", "dom@example.com"); err != nil {
		log.Fatal(err)
	}
}
