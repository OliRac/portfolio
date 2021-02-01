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
	var locale string		//for now, locale is a part of the view. this will probably change it its bugging for french anyway...

	for rows.Next() {
		rows.Scan(&locale, &exp.Title, &exp.Description, &exp.Company, &exp.Start, &exp.Stop)
		collect = append(collect, exp)
	}

	return &collect
} 

//Retrieves about section information with the given language
func (db *Database) GetAbout(lang string) *About {
	rows := db.GetView("v_about_" + lang)

	var ab About

	for rows.Next() {
		rows.Scan(&ab.Quote, &ab.Description)
	}

	return &ab
} 

//Retrieves hero section information with the given language
func (db *Database) GetHero(lang string) *Hero {
	rows := db.GetView("v_hero_" + lang)
	
	var hero Hero

	for rows.Next() {
		rows.Scan(&hero.Name, &hero.Title)
	}

	return &hero
} 

//Retrieves the skill view. These are language independant.
func (db *Database) GetSkill() *[]Skill {
	rows := db.GetView("v_skill")

	var skill Skill
	var collect []Skill

	for rows.Next() {
		rows.Scan(&skill.Name)
		collect = append(collect, skill)
	}

	return &collect
}

//Retrieves all section headers of the given language
func (db *Database) GetSectionHeaders(lang string) *Headers{
	rows := db.GetView("v_section_" + lang)

	var str string
	var collect []string
	var head Headers

	for rows.Next() {
		rows.Scan(&str)
		collect = append(collect, str)
	}

	//Don't do as I do LOL ¯\_(ツ)_/¯
	head.About = collect[0]
	head.Skills = collect[1]
	head.Knowledge = collect[2]
	head.Education = collect[3]
	head.Experience = collect[4]
	head.Links = collect[5]

	return &head
}