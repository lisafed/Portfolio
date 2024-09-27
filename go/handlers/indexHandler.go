package handlers

import (
	"html/template"
	"net/http"
)

type UserData struct {
	IsUserLogged bool
}

var UserLogged = UserData{IsUserLogged: false}

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
