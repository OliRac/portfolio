package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func homeHandler(resp http.ResponseWriter, req *http.Request){

	//since the pattern is just a slash, just checking if theres anything else in there
	//Entering junk along with slash doesnt give 404 errors
	if req.URL.Path != "/" {
		//errorHandler(resp, req, http.StatusNotFound)
		
		resp.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(resp, "404 not found")		//custom 404 could go here

        return
    }

	fmt.Fprintf(resp, home)
}

func glslHandler(resp http.ResponseWriter, req *http.Request){
	fmt.Fprintf(resp, "to do")
}

//Helper for errors
func checkErr(e error){
	if e!= nil {
		panic(e)
	}
}

//Loads a text file in memory
//Panic on error
func loadFile(path string) string {
	data, err := ioutil.ReadFile(path)

	checkErr(err)

	return string(data)
}

//Pages to load in memory
var home = loadFile("./static/sample.html")
//var glsl = loadFile("...")

func main() {
	fmt.Println("Starting up server...")

	//Load up the pages to server memory
	//page := loadFile("./static/sample.html")
	//fmt.Println(page)

	//staticDir := "./static/sample.html"
	//fs := http.FileServer(http.Dir(staticDir))

	//Using the default Mux (this is a simple project anyway, it will suffice)
	/*http.HandleFunc("/", func(res http.ResponseWriter, req * http.Request){
		http.ServeFile(res, req, staticDir)
	})*/

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/glsl", glslHandler)

	//Create & launch server
	http.ListenAndServe(":8080", nil)
}