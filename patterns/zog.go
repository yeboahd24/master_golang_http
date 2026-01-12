package main

import (
	"fmt"

	z "github.com/Oudwins/zog"
)

type User struct {
	Name  string
	Email string
	Age   int
	Bio   string
}

var userSchema = z.Struct(z.Shape{
	"name":  z.String().Min(3).Max(10),
	"email": z.String().Email(),
	"age":   z.Int().GTE(18),
	"bio":   z.String().Optional(),
})

func main() {
	user := User{
		Name:  "Dominic",
		Email: "dominic@gmail.com",
		Age:   20,
		Bio:   "",
	}

	errs := userSchema.Validate(&user)
	if errs != nil {
		for _, issue := range errs {
			fmt.Printf("%s: %s\n", issue.Path, issue.Message)
		}
	}
}
