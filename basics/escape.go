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

func f() *int {
	x := 1
	return &x
}

func main() {
	println(add(1, 2))
	println(NewUser("John"))
	println(*f())
}
