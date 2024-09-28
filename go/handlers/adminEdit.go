package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func AdminEditHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./src/html/Admin-Edit.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var intPathVal int
	if r.PathValue("id") != "" {
		intPathVal, err = strconv.Atoi(r.PathValue("id"))
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	UserLogged.UnitDataExperience, err = db.ReadExperience(intPathVal)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = tmpl.Execute(w, UserLogged)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}
