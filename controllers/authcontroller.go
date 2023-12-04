package controllers

import (
	"html/template"
	"net/http"
)

type UserInput struct {
	Username string
	Password string
}

func Index(w http.ResponseWriter, r *http.Request) {

	temp, _ := template.ParseFiles("views/index.html")
	temp.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		temp, _ := template.ParseFiles("views/login.html")
		temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		UserInput := &UserInput{
			Username: r.Form.Get("username"),
			Password: r.Form.Get("password"),
		}
	}

}
