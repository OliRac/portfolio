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
	CheckErrorFatal(err)

	err = json.Unmarshal(eduData, &data.Edus.Entries)
	CheckErrorFatal(err)


	expData, err := ioutil.ReadFile(filepath.Join(dir, "experience.json"))
	CheckErrorFatal(err)

	err = json.Unmarshal(expData, &data.Exps.Entries)
	CheckErrorFatal(err)


	knowData, err := ioutil.ReadFile(filepath.Join(dir, "knowledge.json"))
	CheckErrorFatal(err)

	err = json.Unmarshal(knowData, &data.Knows.Entries)
	CheckErrorFatal(err)


	skillData, err := ioutil.ReadFile(filepath.Join(dir, "skills.json"))
	CheckErrorFatal(err)

	err = json.Unmarshal(skillData, &data.Skills.Entries)
	CheckErrorFatal(err)


	aboutData, err := ioutil.ReadFile(filepath.Join(dir, "about.json"))
	CheckErrorFatal(err)

	err = json.Unmarshal(aboutData, &data.Info)
	CheckErrorFatal(err) 


	heroData, err := ioutil.ReadFile(filepath.Join(dir, "hero.json"))
	CheckErrorFatal(err)

	err = json.Unmarshal(heroData, &data.Person)
	CheckErrorFatal(err) 


	//Setting headers
	headerData, err := ioutil.ReadFile(filepath.Join(dir, "headers.json"))
	CheckErrorFatal(err)

	err = json.Unmarshal(headerData, &data.NavbarItems)
	CheckErrorFatal(err) 

	//for the moment, the navbar is comprised of the page's headers. Might change later.
	data.Edus.Header = data.NavbarItems.Education
	data.Exps.Header = data.NavbarItems.Experience
	data.Knows.Header = data.NavbarItems.Knowledge
	data.Skills.Header = data.NavbarItems.Skills
	data.Info.Header = data.NavbarItems.About

	

	return &data
}