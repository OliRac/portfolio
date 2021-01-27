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
func importData(dir string) *SiteData {
	var data SiteData

	eduData, err := ioutil.ReadFile(filepath.Join(dir, "education.json"))
	checkError(err, true)

	err = json.Unmarshal(eduData, &data.Edus)
	checkError(err, true)


	expData, err := ioutil.ReadFile(filepath.Join(dir, "experience.json"))
	checkError(err, true)

	err = json.Unmarshal(expData, &data.Exps)
	checkError(err, true)


	knowData, err := ioutil.ReadFile(filepath.Join(dir, "knowledge.json"))
	checkError(err, true)

	err = json.Unmarshal(knowData, &data.Knows)
	checkError(err, true)


	skillData, err := ioutil.ReadFile(filepath.Join(dir, "skills.json"))
	checkError(err, true)

	err = json.Unmarshal(skillData, &data.Skills)
	checkError(err, true)


	aboutData, err := ioutil.ReadFile(filepath.Join(dir, "about.json"))
	checkError(err, true)

	err = json.Unmarshal(aboutData, &data.Info)
	checkError(err, true) 


	heroData, err := ioutil.ReadFile(filepath.Join(dir, "hero.json"))
	checkError(err, true)

	err = json.Unmarshal(heroData, &data.Person)
	checkError(err, true) 

	return &data
}