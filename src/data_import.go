package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"path/filepath"
)

//Smells pretty bad in here
func importData(dir string) (SiteData, error){
	fmt.Println("Attempting to import data from", dir)

	var data SiteData

	eduData, err := ioutil.ReadFile(filepath.Join(dir, "education.json"))
	checkError(err)

	err = json.Unmarshal(eduData, &data.Edus)
	checkError(err)


	expData, err := ioutil.ReadFile(filepath.Join(dir, "experience.json"))
	checkError(err)

	err = json.Unmarshal(expData, &data.Exps)
	checkError(err)


	knowData, err := ioutil.ReadFile(filepath.Join(dir, "knowledge.json"))
	checkError(err)

	err = json.Unmarshal(knowData, &data.Knows)
	checkError(err)


	skillData, err := ioutil.ReadFile(filepath.Join(dir, "skills.json"))
	checkError(err)

	err = json.Unmarshal(skillData, &data.Skills)
	checkError(err)

	return data, err
}

func importDataTwo(dir string, files [] string) *SiteData {
	for _, f := range files {
		fmt.Println(f)
	}

	return &SiteData{}
}

func importEducation(path string) ([]Education, error) {
	fmt.Println("Attempting to read", path)

	rawData, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var result [] Education

	err = json.Unmarshal(rawData, &result)

	if err != nil {
		return nil, err
	}

	fmt.Println("Import successful")

	return result, nil
}


func importExperience(path string) ([]Experience, error) {
	fmt.Println("Attempting to read", path)

	rawData, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var result [] Experience

	err = json.Unmarshal(rawData, &result)

	if err != nil {
		return nil, err
	}

	fmt.Println("Import successful")

	return result, nil
}