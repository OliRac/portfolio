package main

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
	End string
}

type Knowledge struct {
	Name string
}

type Skill struct {
	Name string
}