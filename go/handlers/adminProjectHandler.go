package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"portfolio/go/dbFunc"
)

func AdminProjectHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./src/html/AdminProject.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nameProject := r.FormValue("nomDuProjet")
	Dossier := r.FormValue("dossier")
	Langage := r.FormValue("langage")

	if nameProject != "" && Dossier != "" && Langage != "" {
		db.AddProjet(dbFunc.Projet{NomProjet: nameProject, LienRepo: Dossier, Langage: Langage})
	}
	UserLogged.DataProjectList, err = db.GetProjets()
	if err != nil {
		fmt.Println(err)
	}
	err = tmpl.Execute(w, UserLogged)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}
