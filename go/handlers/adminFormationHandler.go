package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"portfolio/go/dbFunc"
)

func AdminFormationHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./src/html/AdminFormation.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	nameFormation := r.FormValue("NomDeFormation")
	formation := r.FormValue("formation")
	etablissement := r.FormValue("etablissement")
	DateDeDebut := r.FormValue("dateDebut")
	DateDeFin := r.FormValue("dateFin")

	if nameFormation != "" && formation != "" && etablissement != "" && DateDeDebut != "" && DateDeFin != "" {
		db.AddFormation(dbFunc.Formation{Nom: nameFormation, Formation: formation, Etablissement: etablissement, DateDebut: DateDeDebut, DateFin: DateDeFin})
	}
	UserLogged.DataFormationList, err = db.GetFormation()
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
