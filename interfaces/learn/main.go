package main

import "fmt"

// The "contract" — anything that can Speak() is an Animal
type Animal interface {
	Speak() string
}

// Dog has a Speak() method, so it's an Animal automatically
type Dog struct{}

func (d Dog) Speak() string { return "Woof!" }

// Cat has a Speak() method, so it's also an Animal
type Cat struct{}

func (c Cat) Speak() string { return "Meow!" }

// This function doesn't care if it's a Dog, Cat, or anything else.
// It just needs something that can Speak().
func MakeNoise(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	MakeNoise(Dog{}) // Woof!
	MakeNoise(Cat{}) // Meow!
}

// Testing
// Instead of depending on a real DB directly...
type UserStore interface {
	GetUser(id int) (User, error)
}

// Your function depends on the interface
func Greet(store UserStore, id int) string {
	user, _ := store.GetUser(id)
	return "Hello, " + user.Name
}
