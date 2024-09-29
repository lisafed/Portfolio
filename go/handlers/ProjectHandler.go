package handlers

import (
	"html/template"
	"net/http"
)

// ProjectHandlers is a function used to handle the project page
func ProjectHandlers(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./src/html/project.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Is used to get the cookie from the client and check if he is logged
	userCookie, cookieErr := r.Cookie("adminSession")

	// If the user has a cookie but isn't in the map delete the cookie
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

	// Get every project
	UserLogged.DataProjectList, err = db.GetProjets()

	err = tmpl.Execute(w, UserLogged)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
