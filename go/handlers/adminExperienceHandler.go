package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"portfolio/go/dbFunc"
)

func AdminExperienceHandler(w http.ResponseWriter, r *http.Request) {
	userCookie, cookieErr := r.Cookie("adminSession")
	tmpl, err := template.ParseFiles("./src/html/AdminExperience.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if cookieErr != nil {
		http.Error(w, "Unauthorized access", http.StatusUnauthorized)
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
