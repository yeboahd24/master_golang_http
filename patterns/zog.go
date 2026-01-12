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

// Schema

var userSchema = z.Struct(z.Shape{
	"name":  z.String().Min(3, z.Message("Override default message")).Max(10),
	"email": z.String().Email(),
	"age":   z.Int().GTE(18),
	"bio":   z.String().Optional(),
})

// Schema Validation
func main(){
	user := User{
		Name: "Dominic",
		Email: "wrongmail.com",
		Age: 12,
		Bio: "",
	}

	errs := userSchema.Validate(&user)
	if errs !=nil{
		for issue : range errs{
			fmt.Printf("%s: %s\n", z.ZogIssueList(z.Issues.GroupByFlattenedPath()[]))
		}
	}

}
