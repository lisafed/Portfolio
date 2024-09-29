package handlers

import (
	"html/template"
	"net/http"
)

// AdminHandler is function used to handle the admin page
func AdminHandler(w http.ResponseWriter, r *http.Request) {
	// Is used to get the cookie from the client and check if he is logged
	userCookie, cookieErr := r.Cookie("adminSession")
	// If the user doesn't have a cookie return this error
	if cookieErr != nil {
		http.Error(w, "Unauthorized access", http.StatusUnauthorized)
		return
	}

	tmpl, err := template.ParseFiles("./src/html/Admin.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If the user has a cookie deletes it if it doesn't appear in the map
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

	authenticated, ok := UserSession[userCookie.Value]

	// Print an error if the user is not logged
	if !ok || !authenticated {
		http.Error(w, "Forbidden access", http.StatusForbidden)
		return
	}

	err = tmpl.Execute(w, UserLogged)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}
