package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type databaseConn struct {
	DB *sql.DB
}

func (db *databaseConn) createTask(taskName string, date string, time string) {
}

func (db *databaseConn) readTask() {

}

func (db *databaseConn) updateTask() {

}

func (db *databaseConn) deleteTask() {

}

func (db *databaseConn) createUser(userName string, email string, password string, fName string, lName string, avatarURL string, app *application) {
	sqlQuery := "INSERT INTO users(Username, Email, Password, FirstName, LastName, AvatarURL) VALUES (?, ?, ?, ?, ?, ?)"

	exec, err := db.DB.Exec(sqlQuery, userName, email, password, lName, fName, avatarURL)
	app.errorHandler(err)
	rowsAffected, err := exec.RowsAffected()
	app.errorHandler(err)
	fmt.Println(rowsAffected)
}

func (db *databaseConn) updateUser() {

}
