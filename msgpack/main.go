package main

import (
	"github.com/vmihailenco/msgpack/v5"
	"github.com/yeboahd24/msgpack/models"
)

type Item struct {
	Name  string
	Price int
}

// struct tags
type Item2 struct {
	Name  string `msgpack:"name"`
	Price int    `msgpack:"price"`
}

func main() {
	b, err := msgpack.Marshal(&Item{"Apple", 100})
	if err != nil {
		panic(err)
	}
	var item Item
	err = msgpack.Unmarshal(b, &item)
	if err != nil {
		panic(err)
	}
	println(item.Name)

	b2, err := msgpack.Marshal(&Item2{"Apple", 100})
	if err != nil {
		panic(err)
	}
	var item2 Item2
	err = msgpack.Unmarshal(b2, &item2)
	if err != nil {
		panic(err)
	}
	println(item2.Name)

	// Use generated methods
	user := models.User{
		ID:       1,
		Username: "yeboahd24",
		Email:    "yeboahd24@gmail.com",
	}
	data, err := user.MarshalMsg(nil)
	if err != nil {
		panic(err)
	}

	var decoded models.User
	_, err = decoded.UnmarshalMsg(data)
	if err != nil {
		panic(err)
	}
	println(decoded.Username)
}
