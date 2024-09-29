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
	userCookie, cookieErr := r.Cookie("adminSession")

	if cookieErr == nil {
		UserLogged.IsUserLogged = true
		if !UserSession[userCookie.Value] {
			// Delete the cookie for the user if it doesn't exist in the map
			http.SetCookie(w, &http.Cookie{
				Name:   "adminSession",
				Value:  userCookie.Value,
				Path:   "/",
				MaxAge: -1, // Set the time to -1
			})
		}
	} else {
		UserLogged.IsUserLogged = false
	}

	err = tmpl.Execute(w, UserLogged)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
