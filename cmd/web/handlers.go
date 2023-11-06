package main

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

// index is the handler for the home page of the website
func (app *application) index(c *gin.Context) {
	t, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		app.serverError(c.Writer, err)
		return
	}
	err = t.Execute(c.Writer, nil)
	if err != nil {
		app.serverError(c.Writer, err)
		return
	}
}

// login is the handler for the user logging in for the first time. It takes care of inserting the user's value in the database
func (app *application) login(c *gin.Context) {
	if c.Request.Method == http.MethodPost {
		inputEmail := c.PostForm("email")
		inputPassword := c.PostForm("password")

		// Check if the email exists
		doesUserExists := "SELECT Username FROM Users WHERE Email = ?"

		row := app.database.DB.QueryRow(doesUserExists, inputEmail)

		var userExists string
		err := row.Scan(&userExists)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				// User doesn't exists
				c.Redirect(http.StatusMovedPermanently, "/signup")
			} else {
				// Handle error
				c.String(http.StatusInternalServerError, "Error while logging in")
			}
		} else {
			// User exists
			psw, err := app.database.loginUser(inputEmail)
			if err != nil {
				app.errorLog.Fatal(err.Error())
			}
			if psw == inputPassword {
				c.Redirect(http.StatusMovedPermanently, "/homepage")
			} else {
				c.String(http.StatusInternalServerError, "Wrong password")
			}
		}
	} else {
		c.HTML(http.StatusOK, "login.html", nil)
	}
}

// signup is the handler for the user signing up for the first time. It takes care of validating the user's value and redirect to the homepage page of the website
func (app *application) signup(c *gin.Context) {
	if c.Request.Method == http.MethodPost {
		firstname := c.PostForm("firstname")
		lastname := c.PostForm("lastname")
		email := c.PostForm("email")
		password := c.PostForm("password")
		passwordVerification := c.PostForm("password_verification")
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
					c.String(http.StatusInternalServerError, "Password don't match")
					return
				}

				app.database.createUser(username, email, password, firstname, lastname, "./ui/static/avatar/default_avatar.png", app)
				c.Redirect(http.StatusMovedPermanently, "/login")

			} else {
				// User exists
				c.Redirect(http.StatusMovedPermanently, "/login")
			}
		} else {
			// display error in the browser
			c.String(http.StatusInternalServerError, "Error while creating user")
			return
		}

	} else {
		c.HTML(http.StatusOK, "signup.html", nil)

	}
}

// homepage is the handler for the user's homepage. It takes care of displaying the user's information and the user's tasks
func (app *application) homepage(c *gin.Context) {
	t, err := template.ParseFiles("./ui/html/homepage.html")
	if err != nil {
		app.serverError(c.Writer, err)
		return
	}
	err = t.Execute(c.Writer, nil)
	if err != nil {
		app.serverError(c.Writer, err)
		return
	}
}
