package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"portfolio/go/dbFunc"
)

func AdminFormationHandler(w http.ResponseWriter, r *http.Request) {
	userCookie, cookieErr := r.Cookie("adminSession")
	if cookieErr != nil {
		http.Error(w, "Unauthorized access", http.StatusUnauthorized)
		return
	}

	tmpl, err := template.ParseFiles("./src/html/AdminFormation.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err == nil {
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
