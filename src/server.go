package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"html/template"
	"io/ioutil"
	"encoding/json"
	"os"
)


//Files & directories to serve, port to use
//Assuming globals use PascalCase in Go to differentiate from locals
var (
	StaticDir = filepath.FromSlash("../static/")
	TemplateDir = filepath.FromSlash("../templates/")
	DataDir = filepath.FromSlash("../data/")
	Port = "8080"
)


//A dummy sample. Will simply serve the same page as static/index.html but split off in template sections
//Doing this to learn go html/templates
//Going step by step! slowly but surely.
func serveSample(res http.ResponseWriter, req *http.Request) {
	tmpl, _ := template.ParseGlob(TemplateDir + "*.html")

	err := tmpl.ExecuteTemplate(res, "resume", nil)

	if err != nil{
		fmt.Println("Error executing template :(")
		fmt.Println(err)
	}
}

//Reads the JSON file at the given path and returns a pre-set array of maps string:string
//Errors are not handled here, just raised and returned
func readJSONtoMap(path string) ([] map[string]string, error){
	fmt.Println("Attempting to read", path)

	rawData, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var result [] map[string] string

	err = json.Unmarshal(rawData, &result)

	if err != nil {
		return nil, err
	}

	fmt.Println("Import successful")

	return result, nil
}

func checkError(e error) {
	if e != nil{
		fmt.Println(e)
		os.Exit(1)
	}
}

func main() {
	fmt.Println("Starting up server...")
	
	fmt.Println("Importing data...")

	edus, err := readJSONtoMap(DataDir + "education.json")
	checkError(err)

	exps, err := readJSONtoMap(DataDir + "experience.json")
	checkError(err)

	knows, err := readJSONtoMap(DataDir + "knowledge.json")
	checkError(err)

	skills, err := readJSONtoMap(DataDir + "skills.json")
	checkError(err)

	fmt.Println(edus, exps, knows, skills)

	server := http.FileServer(http.Dir(StaticDir))
	http.Handle("/static/", http.StripPrefix("/static/", server))
	http.HandleFunc("/", serveSample)

	fmt.Println("Listening on", Port)
	err = http.ListenAndServe(":" + Port, nil)

	if err != nil{
		panic(err)
	}
}