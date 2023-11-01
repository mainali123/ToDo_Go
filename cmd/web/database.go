package main

import "database/sql"

type databaseConn struct {
	DB *sql.DB
}

// func (db *databaseConn)