package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"portfolio/go/dbFunc"
)

// AdminFormationHandler is a function used to handle the admin formation page
func AdminFormationHandler(w http.ResponseWriter, r *http.Request) {
	// Is used to get the cookie from the client and check if he is logged
	userCookie, cookieErr := r.Cookie("adminSession")
	// If the user doesn't have a cookie return this error
	if cookieErr != nil {
		http.Error(w, "Unauthorized access", http.StatusUnauthorized)
		return
	}

	tmpl, err := template.ParseFiles("./src/html/AdminFormation.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If the user has a cookie deletes it if it doesn't appear in the map
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

	// Print an error if the user is not logged
	if !ok || !authenticated {
		http.Error(w, "Forbidden access", http.StatusForbidden)
		return
	}

	// Get the value from the form to create a new experience
	nameFormation := r.FormValue("NomDeFormation")
	formation := r.FormValue("formation")
	etablissement := r.FormValue("etablissement")
	DateDeDebut := r.FormValue("dateDebut")
	DateDeFin := r.FormValue("dateFin")

	// Check if the value are not empty otherwise don't add the formation into the database
	if nameFormation != "" && formation != "" && etablissement != "" && DateDeDebut != "" && DateDeFin != "" {
		db.AddFormation(dbFunc.Formation{Nom: nameFormation, Formation: formation, Etablissement: etablissement, DateDebut: DateDeDebut, DateFin: DateDeFin})
	}

	// Get the data from the db to show it on the page
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
