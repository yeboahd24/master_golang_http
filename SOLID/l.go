package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

// L — Liskov Substitution Principle
// Any value that satisfies an interface should be substitutable for any other.
// In Go, this is mostly about not lying with your interface — if your implementation panics, does nothing, or returns wrong types, you've violated LSP.

type Reader interface {
	Read(p []byte) (n int, err error)
}

// Both satisfy io.Reader and can be used interchangeably
var (
	r Reader = strings.NewReader("hello")
	r Reader = bytes.NewBuffer([]byte("hello"))
	r Reader = os.Open("file.txt") // *os.File also satisfies it
)

// Violation: a "reader" that does nothing real
type FakeReader struct{}

func (f FakeReader) Read(p []byte) (int, error) {
	return 0, nil // silently returns nothing — breaks callers' expectations
}

func main() {
	var r Reader = FakeReader{}
	b := make([]byte, 10)
	n, err := r.Read(b)
	fmt.Println(n, err)
}
