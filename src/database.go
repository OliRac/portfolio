package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "portfolio_viewer"
    password = "ChargerPledgeBamboo0192"
    dbname   = "portfolio_dev"
)

type Database struct {
	conn *sql.DB
}

func (db *Database) Connect() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	c, err := sql.Open("postgres", connStr)
	checkError(err, true)

	db.conn = c
}

func (db *Database) Disconnect() {
	checkError(db.conn.Close(), true)
}

func (db *Database) CheckConn() {
	checkError(db.conn.Ping(), true)
}

func (db *Database) GetView(view string) *sql.Rows {
	db.CheckConn()

	q := fmt.Sprintf("SELECT * FROM %s", view)

	rows, err := db.conn.Query(q)
	checkError(err, true)			//split up checkError in two: fatal and non fatal

	return rows
}