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
	Port = "8080"
)


//Will simply serve the same page as static/index.html but split off in template sections
//Doing this to learn go html/templates
//Going step by step! slowly but surely.
func ServeResume(res http.ResponseWriter, req *http.Request){
	lang := "en"	//fetching the lang in request later

	data := BuildResume(lang)

	tmpl := template.Must(template.ParseGlob(TemplateDir + "resume/" + "*.html"))

	err := tmpl.ExecuteTemplate(res, "resume", data)

	CheckErrorNonFatal(err)
}

//Stops the current process if there is an error
func CheckErrorFatal(e error){
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

//Will not stop the process on error, but will still show the error on console
func CheckErrorNonFatal(e error){
	if e != nil {
		fmt.Println(e)
	}
}


func main() {
	fmt.Println("Starting up server...")

	server := http.FileServer(http.Dir(StaticDir))
	http.Handle("/static/", http.StripPrefix("/static/", server))
	http.HandleFunc("/", ServeResume)

	fmt.Println("Listening on", Port)
	err := http.ListenAndServe(":" + Port, nil)

	CheckErrorFatal(err)
}