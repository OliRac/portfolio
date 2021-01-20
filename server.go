package main

import (
	"fmt"
	"net/http"
	"path/filepath"
)

//Also handles 404 if anything is after /
//Servefile does some sanitizing already
//home is unforunately a global. Will try to come up with a way to not use globals.
func homeHandler(res http.ResponseWriter, req *http.Request){
	if req.URL.Path != "/" {
		res.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(res, "404 not found")		//custom 404 could go here
		return
	}

	http.ServeFile(res, req, home)
}

func glslHandler(res http.ResponseWriter, req *http.Request){
	fmt.Fprintf(res, "to do")
}


//Files to serve
var home = filepath.FromSlash("./static/sample.html")
//var glsl = filepath.FromSlash("./static/glsl.html")

func main() {
	fmt.Println("Starting up server...")

	//init and launch
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/glsl", glslHandler)

	http.ListenAndServe(":8080", nil)
}