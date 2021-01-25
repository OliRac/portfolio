package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"html/template"
	"os"
)


//Files & directories to serve, port to use
//Assuming globals use PascalCase in Go to differentiate from locals
var (
	StaticDir = filepath.FromSlash("../static/")
	TemplateDir = filepath.FromSlash("../templates/")
	DataDir = filepath.FromSlash("../data/")
	JSONFiles = []string{"education.json", "experience.json", "knowledge.json", "skills.json"}
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

func checkError(e error) {
	if e != nil{
		fmt.Println(e)
		os.Exit(1)
	}
}

func main() {
	fmt.Println("Starting up server...")
	
	fmt.Println("Importing data...")

	data := importDataTwo(DataDir, JSONFiles)

	fmt.Println(data)

	//data, _ := importData(DataDir)

	//fmt.Println(data)

	//edus, err := importEducation(DataDir + "education.json")
	//checkError(err)

	//fmt.Println(edus)

	/*edus, err := readJSONtoMap(DataDir + "education.json")
	checkError(err)

	exps, err := readJSONtoMap(DataDir + "experience.json")
	checkError(err)

	knows, err := readJSONtoMap(DataDir + "knowledge.json")
	checkError(err)

	skills, err := readJSONtoMap(DataDir + "skills.json")
	checkError(err)

	fmt.Println(edus, exps, knows, skills)*/

	/*server := http.FileServer(http.Dir(StaticDir))
	http.Handle("/static/", http.StripPrefix("/static/", server))
	http.HandleFunc("/", serveSample)

	fmt.Println("Listening on", Port)
	err = http.ListenAndServe(":" + Port, nil)

	if err != nil{
		panic(err)
	}*/
}