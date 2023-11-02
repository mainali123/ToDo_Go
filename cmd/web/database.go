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

func (db *databaseConn) createUser(userName string, email string, password string, fName string, lName string, avatarURL string, app *application) int {
	sqlQuery := "INSERT INTO Users(Username, Email, Password, FirstName, LastName, AvatarURL) VALUES (?, ?, ?, ?, ?, ?)"

	exec, err := db.DB.Exec(sqlQuery, userName, email, password, lName, fName, avatarURL)
	if err != nil {
		app.errorLog.Fatal(err.Error())
	}
	rowsAffected, err := exec.RowsAffected()
	if err != nil {
		app.errorLog.Fatal(err.Error())
	}
	fmt.Println(rowsAffected)
	return int(rowsAffected)
}

func (db *databaseConn) updateUser() {

}
