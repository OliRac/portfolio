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

//Connects to the pre-defined server, see constants above
//Exits on eror
func (db *Database) Connect() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	c, err := sql.Open("postgres", connStr)
	CheckErrorFatal(err)

	db.conn = c
}

//Exits on error
func (db *Database) Disconnect() {
	CheckErrorFatal(db.conn.Close())
}

//Use this to make sure the connection is still up
//Exits on error
func (db *Database) CheckConn() {
	CheckErrorFatal(db.conn.Ping())
}

//Get an existing view (or table) from the database
func (db *Database) GetView(view string) *sql.Rows {
	db.CheckConn()

	q := fmt.Sprintf("SELECT * FROM %s", view)

	rows, err := db.conn.Query(q)
	CheckErrorFatal(err)

	return rows
}

//Retrieves all education of the given language
func (db *Database) GetEducation(lang string) *[] Education{
	rows := db.GetView("v_edu_" + lang)

	var edu Education
	var collect []Education

	for rows.Next() {
		rows.Scan(&edu.Degree, &edu.Institution, &edu.Location, &edu.Date)
		collect = append(collect, edu)
	}

	return &collect
} 

//Retrieves all knowledge of the given language
func (db *Database) GetKnowledge(lang string) *[] Knowledge{
	rows := db.GetView("v_know_" + lang)

	var know Knowledge
	var collect []Knowledge

	for rows.Next() {
		rows.Scan(&know.Name)
		collect = append(collect, know)
	}

	return &collect
} 

//Retrieves all experiences of the given language
func (db *Database) GetExperience(lang string) *[] Experience{
	rows := db.GetView("v_exp_" + lang)

	var exp Experience
	var collect []Experience

	for rows.Next() {
		rows.Scan(&exp.Title, &exp.Description, &exp.Company, &exp.Duration)
		collect = append(collect, exp)
	}

	return &collect
} 