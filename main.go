package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Define your routes here
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to MongoDB API")
	})

	// Add more routes here...

	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
