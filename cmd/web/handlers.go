package main

import (
	"html/template"
	"net/http"
)

// index is the handler for the home page of the website
func (app *application) index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

// login is the handler for the user loggin in for the first time. It takes care of inserting the user's value in the database
func (app *application) login(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./ui/html/login.html")
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

// signup is the handler for the user signing up for the first time. It takes care of validating the user's value and redirect to the homepage page of the website
func (app *application) signup(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./ui/html/signup.html")
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

// homepage is the handler for the user's homepage. It takes care of displaying the user's information and the user's tasks
func (app *application) homepage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./ui/html/homepage.html")
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
		return
	}
}
