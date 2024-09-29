package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"portfolio/go/dbFunc"
	"strconv"
)

func AdminEditExperienceHandler(w http.ResponseWriter, r *http.Request) {
	userCookie, cookieErr := r.Cookie("adminSession")
	if cookieErr != nil {
		http.Error(w, "Unauthorized access", http.StatusUnauthorized)
		return
	}

	tmpl, err := template.ParseFiles("./src/html/AdminEditExperience.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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

	if !ok || !authenticated {
		http.Error(w, "Forbidden access", http.StatusForbidden)
		return
	}

	var intPathVal int
	if r.PathValue("id") != "" && r.PathValue("id") != "js" {
		intPathVal, err = strconv.Atoi(r.PathValue("id"))
		if err != nil {
			fmt.Println(err)
		}
	}

	UserLogged.UnitDataExperience, err = db.ReadExperience(intPathVal)
	if err != nil {
		return
	}
	SupprButton := r.FormValue("DelButton")
	button := r.FormValue("editbutton")
	newNom := r.FormValue("nom")
	newEntreprise := r.FormValue("entreprise")
	newPoste := r.FormValue("poste")
	newDateDebut := r.FormValue("dateDebut")
	newDateFin := r.FormValue("dateFin")

	fmt.Println("rpathvalue :", r.PathValue("id"))
	fmt.Println("Suppr :", SupprButton, "\n", "Button :", button)
	if button == "Modifier" && r.PathValue("id") != "" {
		db.EditExperience(dbFunc.Experience{Nom: newNom, Entreprise: newEntreprise, Poste: newPoste, DateDebut: newDateDebut, DateFin: newDateFin, ID: intPathVal})
		http.Redirect(w, r, "http://localhost:8080/admin/experience", http.StatusSeeOther)
	}

	if SupprButton == "Supprimer" && r.PathValue("id") != "" {
		db.DeleteExperience(intPathVal)
		http.Redirect(w, r, "http://localhost:8080/admin/experience", http.StatusSeeOther)
	}
	err = tmpl.Execute(w, UserLogged)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

func AdminEditProjectHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./src/html/AdminEditProject.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(r.PathValue("id"))
	var intPathVal int
	if r.PathValue("id") != "" && r.PathValue("id") != "js" {
		intPathVal, err = strconv.Atoi(r.PathValue("id"))
		if err != nil {
			fmt.Println(err)
		}
	}

	UserLogged.UnitDataProject, err = db.ReadProjet(intPathVal)
	if err != nil {
		return
	}
	SupprButton := r.FormValue("DelButton")
	button := r.FormValue("editbutton")
	NewNomProjet := r.FormValue("nomDuProjet")
	NewDossier := r.FormValue("dossier")
	NewLangage := r.PathValue("langage")

	if button == "Modifier" && r.PathValue("id") != "" {
		db.EditProjet(dbFunc.Projet{NomProjet: NewNomProjet, LienRepo: NewDossier, Langage: NewLangage, ID: intPathVal})
		http.Redirect(w, r, "http://localhost:8080/admin/project", http.StatusSeeOther)
	}

	if SupprButton == "Supprimer" && r.PathValue("id") != "" {
		db.DeleteProject(intPathVal)
		http.Redirect(w, r, "http://localhost:8080/admin/project", http.StatusSeeOther)
	}
	err = tmpl.Execute(w, UserLogged)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

func AdminEditFormationHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./src/html/AdminEditFormation.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var intPathVal int
	if r.PathValue("id") != "" && r.PathValue("id") != "js" {
		intPathVal, err = strconv.Atoi(r.PathValue("id"))
		if err != nil {
			fmt.Println(err)
		}
	}

	UserLogged.DataFormation, err = db.ReadFormation(intPathVal)
	if err != nil {
		return
	}
	SupprButton := r.FormValue("DelButton")
	button := r.FormValue("editbutton")
	NewNomDeFormation := r.FormValue("nomDeFormation")
	NewFormation := r.FormValue("formation")
	NewEtablissement := r.FormValue("etablissement")
	newDateDeDebut := r.FormValue("dateDeDebut")
	newDateDeFin := r.FormValue("dateDeFin")

	if button == "Modifier" && r.PathValue("id") != "" {
		db.EditFormation(dbFunc.Formation{Nom: NewNomDeFormation, Formation: NewFormation, Etablissement: NewEtablissement, DateDebut: newDateDeDebut, DateFin: newDateDeFin, ID: intPathVal})
		http.Redirect(w, r, "http://localhost:8080/admin/formation", http.StatusSeeOther)
	}

	if SupprButton == "Supprimer" && r.PathValue("id") != "" {
		db.DeleteFormation(intPathVal)
		http.Redirect(w, r, "http://localhost:8080/admin/formation", http.StatusSeeOther)
	}
	err = tmpl.Execute(w, UserLogged)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}
