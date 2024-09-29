package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"portfolio/go/dbFunc"
)

func AdminExperienceHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./src/html/AdminExperience.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nameUser := r.FormValue("Nom")
	entreprise := r.FormValue("entreprise")
	Poste := r.FormValue("poste")
	DateDebut := r.FormValue("dateDebut")
	DatedeFin := r.FormValue("dateFin")

	if nameUser != "" && entreprise != "" && Poste != "" && DateDebut != "" && DatedeFin != "" {
		db.AddExperience(dbFunc.Experience{Nom: nameUser, Entreprise: entreprise, Poste: Poste, DateDebut: DateDebut, DateFin: DatedeFin})

	}
	UserLogged.DataExperienceList, err = db.GetExperience()
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
