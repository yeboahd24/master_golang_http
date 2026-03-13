package main

import "github.com/yeboahd24/interfaces/mysqldb"

type dbContract interface {
	Close()
	InsertUser(userName string) error
	SelectSingleUser(userName string) (string, error)
}

type Application struct {
	db dbContract
}

func NewApplication(db dbContract) *Application {
	return &Application{db: db}
}

func main() {
	db, err := mysqldb.New("root", "root", "localhost", "3306", "interface-v2")
	if err != nil {
		panic(err)
	}

	app := NewApplication(db) // conforms to dbContract

	err = app.db.InsertUser("user")
	if err != nil {
		panic(err)
	}

	user, err := app.db.SelectSingleUser("user")
	if err != nil {
		panic(err)
	}
	println(user)

	app.db.Close()
}
