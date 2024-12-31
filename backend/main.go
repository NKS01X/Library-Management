package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	//"Library-Management/backend/routeslist"

	"github.com/gorilla/mux"
	"github.com/nks01x/Library-Management/routeslist"
)

func main() {
	// creating server
	r := mux.NewRouter()
	//used to serve static files such as ss files and all
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./backend/templates"))))
	//handler functions
	r.HandleFunc("/", serveHome).Methods("GET")
	//signup page
	r.HandleFunc("/signup", routeslist.Signup)
	//signin page
	r.HandleFunc("/signin", routeslist.Signin)

	//server
	fmt.Println("Server started at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
func NilCheckError(err error) {
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
	NilCheckError(err)
	//Executing template
	err = templates.Execute(w, nil)
	NilCheckError(err)

}
