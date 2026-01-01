package main

import (
	"encoding/json"
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
// /users/
func userHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/users/")
	fmt.Fprintf(w, "User ID: %s", id)
}

// reading request data
func requestData(w http.ResponseWriter, r *http.Request) {
	// /search?q=golang&page=2
	query := r.URL.Query()
	search := query.Get("q")
	page := query.Get("page")

	// headers
	userAgent := r.Header.Get("User-Agent")
	contentType := r.Header.Get(("Content-Type"))

	// JSON body
	var data map[string]any
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
}

func main() {
	basicServer()
	routingTesting()
}
