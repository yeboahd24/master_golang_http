package main

import "database/sql"

func CreateUser(db *sql.DB, name, email string) error {
	_, err := db.Exec(
		"INSERT INTO users(name, email) VALUES($1, $2)",
		name,
		email,
	)
	return err
}
