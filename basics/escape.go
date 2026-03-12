package main

// No escape analysis

func add(a, b int) int {
	return a + b
}

// Escape analysis

type User struct {
	Name string
}

func NewUser(name string) *User {
	return &User{Name: name}
}

func main() {
	println(add(1, 2))
	println(NewUser("John"))
}
