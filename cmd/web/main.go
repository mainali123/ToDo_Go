package main

import (
	"fmt"
	"net/http"
)

// To run the web-server enter `go run ./cmd/web` from the root directory of the project file

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/homepage", homepage)

	fmt.Println("INFO:\tStarting server on localhost:8080.")
	err := http.ListenAndServe("localhost:8080", mux)
	errorHandler(err)
}
