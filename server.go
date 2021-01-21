package main

import (
	"fmt"
	"net/http"
	"path/filepath"
)


//Files & directories to serve, port to use
//Assuming globals use PascalCase in Go to differentiate from locals
var (
	StaticDir = filepath.FromSlash("./static")
	Port = "8080"
)


func main() {
	fmt.Println("Starting up server...")

	server := http.FileServer(http.Dir(StaticDir))
	http.Handle("/", server)	//http.StripPrefix("/static/", server))

	fmt.Println("Listening on", Port)
	http.ListenAndServe(":" + Port, nil)
}