package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {
	// creating server
	r := mux.NewRouter()

	//handler function
	r.HandleFunc("/", serveHome)
	//server
	log.Fatal(http.ListenAndServe(":8000", r))
	fmt.Println("Server started at http://localhost:8000")
}
func nilCheckError(err error) {
	//checking the error
	if err != nil {
		fmt.Println("Error occured in :\n", err)
		return
	}
}

// home page
func serveHome(w http.ResponseWriter, r *http.Request) {
	//include html
	//Library-Management/backend/
	templates, err := template.ParseFiles("templates/home.html")
	//templates = name of the parsed file
	nilCheckError(err)
	//Executing template
	err = templates.Execute(w, nil)
	nilCheckError(err)

}
