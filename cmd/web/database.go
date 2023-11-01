package main

import (
	"database/sql"
)

type databaseConn struct {
	DB *sql.DB
}

func (db *databaseConn) create() {
	// hey this is comment
}

func (db *databaseConn) read() {

}

func (db *databaseConn) update() {

}

func (db *databaseConn) delete() {

}
