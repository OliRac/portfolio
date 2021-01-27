package main

//Struct that wraps all others, contains lists/arrays of every type, to use with html templates
type SiteData struct{
	Edus EduSection
	Exps ExpSection
	Knows KnowSection
	Skills SkillSection
	Person Hero
	Info About
	NavbarItems Headers			//for the moment, the navbar is comprised of the page's headers. Might change later.
}

/*Each section will have entries and a header*/
type EduSection struct {
	Header string
	Entries []Education
}

type ExpSection struct {
	Header string
	Entries []Experience
}

type KnowSection struct {
	Header string
	Entries []Knowledge
}

type SkillSection struct {
	Header string
	Entries []Skill
}

/*These sections do not have lists of entries*/
type Hero struct {
	Name string
	Title string
}

type About struct {
	Header string
	Quote string
	Description string
}

type Headers struct {
	About string
	Skills string
	Knowledge string
	Education string
	Experience string
	Links string
}

/*Structs for entries of different types*/
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

/*wow so much difference here!*/
type Knowledge struct {
	Name string
}

type Skill struct {
	Name string
}