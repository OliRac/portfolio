package main

import (
	"io/ioutil"
	"encoding/json"
	"path/filepath"
)

//Smells pretty bad in here
//"Hard coded" data import from JSON files in the data directory
//Will probably replace this with a database. IDK which yet though.
//Error handling happens internally, returning an error could also be a possibility
//Returns an object that is structed such that its entries represents the index page
func importData(dir string) *SiteData {
	var data SiteData

	eduData, err := ioutil.ReadFile(filepath.Join(dir, "education.json"))
	checkError(err, true)

	err = json.Unmarshal(eduData, &data.Edus.Entries)
	checkError(err, true)


	expData, err := ioutil.ReadFile(filepath.Join(dir, "experience.json"))
	checkError(err, true)

	err = json.Unmarshal(expData, &data.Exps.Entries)
	checkError(err, true)


	knowData, err := ioutil.ReadFile(filepath.Join(dir, "knowledge.json"))
	checkError(err, true)

	err = json.Unmarshal(knowData, &data.Knows.Entries)
	checkError(err, true)


	skillData, err := ioutil.ReadFile(filepath.Join(dir, "skills.json"))
	checkError(err, true)

	err = json.Unmarshal(skillData, &data.Skills.Entries)
	checkError(err, true)


	aboutData, err := ioutil.ReadFile(filepath.Join(dir, "about.json"))
	checkError(err, true)

	err = json.Unmarshal(aboutData, &data.Info)
	checkError(err, true) 


	heroData, err := ioutil.ReadFile(filepath.Join(dir, "hero.json"))
	checkError(err, true)

	err = json.Unmarshal(heroData, &data.Person)
	checkError(err, true) 


	//Setting headers
	headerData, err := ioutil.ReadFile(filepath.Join(dir, "headers.json"))
	checkError(err, true)

	err = json.Unmarshal(headerData, &data.NavbarItems)
	checkError(err, true) 

	//for the moment, the navbar is comprised of the page's headers. Might change later.
	data.Edus.Header = data.NavbarItems.Education
	data.Exps.Header = data.NavbarItems.Experience
	data.Knows.Header = data.NavbarItems.Knowledge
	data.Skills.Header = data.NavbarItems.Skills
	data.Info.Header = data.NavbarItems.About

	

	return &data
}