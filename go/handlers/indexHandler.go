package handlers

import (
	"html/template"
	"net/http"
	"portfolio/go/dbFunc"
)

var UserLogged = Data{IsUserLogged: false}
var db = dbFunc.LoadDb()

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./src/html/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = r.Cookie("adminSession")
	if err == nil {
		UserLogged.IsUserLogged = true
	} else {
		UserLogged.IsUserLogged = false
	}

	err = tmpl.Execute(w, UserLogged)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
