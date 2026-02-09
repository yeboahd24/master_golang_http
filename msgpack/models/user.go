package models

//go:generate go run github.com/tinylib/msgp@latest -file user.go

type User struct {
	ID       int    `msgpack:"id"`
	Username string `msgpack:"username"`
	Email    string `msgpack:"email"`
}
