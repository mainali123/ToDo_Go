package main

import (
	"database/sql"
	"errors"
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

// login is the handler for the user logging in for the first time. It takes care of inserting the user's value in the database
func (app *application) login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		inputEmail := r.FormValue("email")
		inputPassword := r.FormValue("password")

		// Check if the email exists
		doesUserExists := "SELECT Username FROM Users WHERE Email = ?"

		row := app.database.DB.QueryRow(doesUserExists, inputEmail)

		var userExists string
		err := row.Scan(&userExists)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				// User doesn't exists
				http.Redirect(w, r, "/signup", 301)
			} else {
				// Handle error
				http.Error(w, "Error while logging in", 500)
			}
		} else {
			// User exists
			psw, err := app.database.loginUser(inputEmail)
			if err != nil {
				app.errorLog.Fatal(err.Error())
			}
			if psw == inputPassword {
				http.Redirect(w, r, "/homepage", 301)
			} else {
				http.Error(w, "Wrong password", 500)
			}
		}
	} else {
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

		ifUserExists := "SELECT Username FROM Users WHERE Email = ?"
		//row := app.database.db.QueryRow(ifUserExists, email)
		row := app.database.DB.QueryRow(ifUserExists, email)

		var userExists string
		err := row.Scan(&userExists)

		// if user exists redirect to the login page
		if err != nil {

			if errors.Is(err, sql.ErrNoRows) {
				// User doesn't exists
				if password != passwordVerification {
					// display error in the browser
					http.Error(w, "Passwords don't match", 500)
					return
				}

				app.database.createUser(username, email, password, firstname, lastname, "./ui/static/avatar/default_avatar.png", app)
				http.Redirect(w, r, "/login", 301)

			} else {
				// User exists
				http.Redirect(w, r, "/login", 301)
			}
		} else {
			// display error in the browser
			http.Error(w, "Error while creating user", 500)
			return
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
