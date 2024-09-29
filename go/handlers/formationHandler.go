package handlers

import (
	"html/template"
	"net/http"
)

// FormationHandler is a function used to handle the formation page that shows to the user
func FormationHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./src/html/formation.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Is used to get the cookie from the client and check if he is logged
	userCookie, err := r.Cookie("adminSession")

	// If the user has a cookie deletes it if it doesn't appear in the map
	if err == nil {
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

	err = tmpl.Execute(w, UserLogged)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
