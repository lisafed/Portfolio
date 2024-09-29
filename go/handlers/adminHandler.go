package handlers

import (
	"html/template"
	"net/http"
)

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("adminSession")
	if err != nil {
		http.Error(w, "Unauthorized access", http.StatusUnauthorized)
		return
	}

	tmpl, err := template.ParseFiles("./src/html/Admin.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	authenticated, ok := UserSession[cookie.Value]

	if !ok || !authenticated {
		http.Error(w, "Forbidden access", http.StatusForbidden)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}
