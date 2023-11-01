package main

import (
	"html/template"
	"log"
	"net/http"
)

// errorHandler is a helper function to handle errors
func errorHandler(err error) {
	log.Fatal(err.Error())
}

// index is the handler for the home page of the website
func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./ui/html/index.html")
	errorHandler(err)
	err = t.Execute(w, nil)
	errorHandler(err)
}

// login is the handler for the user loggin in for the first time. It takes care of inserting the user's value in the database
func login(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./ui/html/login.html")
	errorHandler(err)
	err = t.Execute(w, nil)
	errorHandler(err)
}

// signup is the handler for the user signing up for the first time. It takes care of validating the user's value and redirect to the homepage page of the website
func signup(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./ui/html/signup.html")
	errorHandler(err)
	err = t.Execute(w, nil)
	errorHandler(err)
}

// homepage is the handler for the user's homepage. It takes care of displaying the user's information and the user's tasks
func homepage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./ui/html/homepage.html")
	errorHandler(err)
	err = t.Execute(w, nil)
	errorHandler(err)
}
