package main

//Struct that wraps all others, contains lists/arrays of every type, to use with html templates
type SiteData struct{
	Edus []Education
	Exps []Experience
	Knows []Knowledge
	Skills []Skill
}

type Education struct{
	Degree string
	Institution string
	Location string
	Date string
}

type Experience struct {
	Title string
	Description string
	Company string
	Start string
	Stop string
}

type Knowledge struct {
	Name string
}

type Skill struct {
	Name string
}