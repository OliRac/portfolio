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
	CertPath = filepath.FromSlash("C:/certs/server.crt")	//cert and key are self signed for the moment
	KeyPath = filepath.FromSlash("C:/certs/server.key")
	Port = "8080"
)

//Will simply serve the same page as static/index.html but split off in template sections
//Doing this to learn go html/templates
//Going step by step! slowly but surely.
func ServeResume(res *http.ResponseWriter, req *http.Request, langCode string){
	data := BuildResume(langCode)

	tmpl := template.Must(template.ParseGlob(TemplateDir + "resume/" + "*.html"))

	err := tmpl.ExecuteTemplate(*res, "resume", data)

	CheckErrorNonFatal(err)
}


//Site will be served in english by default
func ServeResumeDefault(res http.ResponseWriter, req *http.Request){
	ServeResume(&res, req, "en")
}


//Adding a special rule to just redirect to default if /en/ is visited (for some reason)
func ServeResumeEN(res http.ResponseWriter, req *http.Request){
	http.Redirect(res, req, "/", http.StatusSeeOther)
}


// of course /fr/ will result in sreving the french version of the site
func ServeResumeFR(res http.ResponseWriter, req *http.Request){
	ServeResume(&res, req, "fr")
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
	//keeping the static version up for now, felt cute might delete later
	http.Handle("/static/", http.StripPrefix("/static/", server))

	http.HandleFunc("/", ServeResumeDefault)
	http.HandleFunc("/en/", ServeResumeEN)
	http.HandleFunc("/fr/", ServeResumeFR)

	fmt.Println("Listening on", Port)
	err := http.ListenAndServeTLS(":" + Port, CertPath, KeyPath, nil)

	CheckErrorFatal(err)
}