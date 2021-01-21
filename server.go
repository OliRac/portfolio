package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"html/template"
	"io/ioutil"
	"encoding/json"
)


//Files & directories to serve, port to use
//Assuming globals use PascalCase in Go to differentiate from locals
var (
	StaticDir = filepath.FromSlash("./static/")
	TemplateDir = filepath.FromSlash("./templates/")
	DataDir = filepath.FromSlash("./data/")
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


func main() {
	fmt.Println("Starting up server...")

	fmt.Println("Reading JSON data...")

	dataFile, err := ioutil.ReadFile(DataDir + "education.json")

	if err != nil {
		fmt.Println(err)
	}

	var test []Education

	err = json.Unmarshal(dataFile, &test)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(test)

	server := http.FileServer(http.Dir(StaticDir))
	http.Handle("/static/", http.StripPrefix("/static/", server))
	http.HandleFunc("/", serveSample)

	fmt.Println("Listening on", Port)
	err = http.ListenAndServe(":" + Port, nil)

	if err != nil{
		panic(err)
	}
}