package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"portfolio/go/dbFunc"
)

// AdminExperienceHandler is used to handle the admin experience page
func AdminExperienceHandler(w http.ResponseWriter, r *http.Request) {
	// Is used to get the cookie from the client and check if he is logged
	userCookie, cookieErr := r.Cookie("adminSession")
	tmpl, err := template.ParseFiles("./src/html/AdminExperience.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// If the user doesn't have a cookie return this error
	if cookieErr != nil {
		http.Error(w, "Unauthorized access", http.StatusUnauthorized)
		return
	}

	// If the user has a cookie deletes it if it doesn't appear in the map
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

	// Print an error if the user is not logged
	if !ok || !authenticated {
		http.Error(w, "Forbidden access", http.StatusForbidden)
		return
	}

	// Get the value from the form to create a new experience
	nameUser := r.FormValue("Nom")
	entreprise := r.FormValue("entreprise")
	Poste := r.FormValue("poste")
	DateDebut := r.FormValue("dateDebut")
	DatedeFin := r.FormValue("dateFin")

	// Check if the value are not empty otherwise don't add the experience into the database
	if nameUser != "" && entreprise != "" && Poste != "" && DateDebut != "" && DatedeFin != "" {
		db.AddExperience(dbFunc.Experience{Nom: nameUser, Entreprise: entreprise, Poste: Poste, DateDebut: DateDebut, DateFin: DatedeFin})

	}

	// Get the data from the db to show it on the page
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
