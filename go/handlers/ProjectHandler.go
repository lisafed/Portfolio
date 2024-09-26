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

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
