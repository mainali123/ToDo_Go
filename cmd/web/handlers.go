package main

import (
	"html/template"
	"net/http"
	"time"
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
	if r.Method == http.MethodPost {
		firstname := r.FormValue("firstname")
		lastname := r.FormValue("lastname")
		email := r.FormValue("email")
		password := r.FormValue("password")
		passwordVerification := r.FormValue("password_verification")
		username := firstname + lastname

		if password != passwordVerification {
			// display error in the browser
			http.Error(w, "Passwords don't match", 500)
			return
		}

		val := app.database.createUser(username, email, password, firstname, lastname, "./ui/static/avatar/default_avatar.png", app)
		if val > 0 {
			// redirect to the login page
			http.Redirect(w, r, "/login", 301)
		} else {
			ifUserExists := "SELECT * FROM users WHERE email = $1"
			//row := app.database.db.QueryRow(ifUserExists, email)
			row := app.database.DB.QueryRow(ifUserExists, email)

			// if user exists redirect to the login page
			if row != nil {
				// show info in the browser
				http.Error(w, "User already exists redirecting to login page", 500)
				// add delay of 5 seconds
				time.Sleep(5 * time.Second)
				http.Redirect(w, r, "/login", 301)
			} else {
				// display error in the browser
				http.Error(w, "Error while creating user", 500)
				return
			}

		}
	} else {
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
