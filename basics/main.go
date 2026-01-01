package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Basic
func basicServer() {
	fmt.Println("Basic func")
	// handler --> anonymous handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hellow, World"))
	})

	// handler_2
	http.HandleFunc("/hello", helloHadler)

	// start server
	// log.Fatal(http.ListenAndServe(":8000", nil))
	fmt.Println("Sever start")
}

// Handler patern
func helloHadler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World From Handler Function")
}

// routing --> for production usage
func routingTesting() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello2", helloHadler)
	mux.HandleFunc("/", homePageHandler)
	mux.HandleFunc("/users", userHandler)

	log.Fatal(http.ListenAndServe(":9000", mux))
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Homepage"))
}

// extracting from Path
func userHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/users/")
	fmt.Fprintf(w, "User ID: %s", id)
}

func main() {
	basicServer()
	routingTesting()
}
