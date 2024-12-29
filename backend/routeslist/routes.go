package routeslist

import (
	"fmt"
	"net/http"
	"text/template"
)

func NilCheckError(err error) {
	if err != nil {
		fmt.Println("There's some issue", err)
		return
	}
}

//registration

// signup
func Signup(w http.ResponseWriter, r *http.Request) {
	SignupTemp, err := template.ParseFiles("templates/signup.html")
	NilCheckError(err)
	err = SignupTemp.Execute(w, nil)
	NilCheckError(err)
}

// Signin
func Signin(w http.ResponseWriter, r *http.Request) {
	SigninTemp, err := template.ParseFiles("templates/signin.html")
	NilCheckError(err)
	err = SigninTemp.Execute(w, nil)
	NilCheckError(err)
}
