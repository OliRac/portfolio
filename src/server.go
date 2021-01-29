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


//Will simply serve the same page as static/index.html but split off in template sections
//Doing this to learn go html/templates
//Going step by step! slowly but surely.
func serveIndex(res http.ResponseWriter, req *http.Request) {

	//Should site data be fetched with every request? or be cached? I'm not entirely sure how to proceed.
	//Leaving it this way is great for debugging, don't need to re-launch the server every time
	//Might cache it on a later release 
	data := importData(DataDir)

	tmpl := template.Must(template.ParseGlob(TemplateDir + "*.html"))

	err := tmpl.ExecuteTemplate(res, "resume", data)

	//firewall causes an error here sometimes, but it does not cause a crash, so its more like a warning lol...
	CheckErrorNonFatal(err)
}

//Stops the current process if there is an error
func CheckErrorFatal(e error){
	if e != nil {
		os.Exit(1)
	}
}

//Will not stop the process on error, but will show the error on console
func CheckErrorNonFatal(e error){
	if e != nil {
		fmt.Println(e)
	}
}


func main() {
	/*fmt.Println("Starting up server...")

	server := http.FileServer(http.Dir(StaticDir))
	http.Handle("/static/", http.StripPrefix("/static/", server))
	http.HandleFunc("/", serveIndex)

	fmt.Println("Listening on", Port)
	err := http.ListenAndServe(":" + Port, nil)

	CheckErrorFatal(err)

	/*query example*/
	var db Database

	db.Connect()

	rows := db.GetView("v_edu_en")

	var edu Education
	var collect []Education

	for rows.Next() {
		rows.Scan(&edu.Degree, &edu.Institution, &edu.Location, &edu.Date)
		collect = append(collect, edu)
	}

	fmt.Println(collect)

	db.Disconnect()
}