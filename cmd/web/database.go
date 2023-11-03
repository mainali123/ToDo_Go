package main

import (
	"database/sql"
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
	sqlQueryInsert := "INSERT INTO Users(Username, Email, Password, FirstName, LastName, AvatarURL) VALUES (?, ?, ?, ?, ?, ?)"
	//exec, err := db.DB.Exec(sqlQueryInsert, userName, email, password, fName, lName, avatarURL)
	_, err := db.DB.Exec(sqlQueryInsert, userName, email, password, fName, lName, avatarURL)

	if err != nil {
		app.errorLog.Fatal(err.Error())
	}
}

func (db *databaseConn) updateUser() {

}

func (db *databaseConn) loginUser(email string) (string, error) {
	sqlQuery := "SELECT Password FROM Users WHERE Email=?"

	exec := db.DB.QueryRow(sqlQuery, email)

	var password_ string

	if err := exec.Scan(&password_); err != nil {
		return "", err
	} else {
		return password_, nil
	}
}
