package handlers

import (
	"html/template"
	"net/http"
)

func ProjectHandlers(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./src/html/project.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userCookie, cookieErr := r.Cookie("adminSession")

	if cookieErr == nil {
		if !UserSession[userCookie.Value] {
			// Delete the cookie for the user if it doesn't exist in the map
			http.SetCookie(w, &http.Cookie{
				Name:   "adminSession",
				Value:  userCookie.Value,
				Path:   "/",
				MaxAge: -1, // Set the time to -1
			})
		}
	}

	UserLogged.DataProjectList, err = db.GetProjets()

	err = tmpl.Execute(w, UserLogged)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
