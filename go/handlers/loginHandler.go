package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"portfolio/go/dbFunc"

	"github.com/gofrs/uuid"
)

var UserSession = make(map[string]bool)

// LoginHandler is a function used to handle the login page
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./src/html/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Is used to get the cookie from the client and check if he is logged
	userCookie, sessErr := r.Cookie("adminSession")
	if err != nil {
		fmt.Println(err)
	}

	// Generate a uuid
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}

	// Get the username and the password the user insert in the form and load the database
	username := r.FormValue("username")
	password := r.FormValue("password")
	LoadedDB := dbFunc.LoadDb()

	// Check if the user exist
	userData, err := dbFunc.GetUtilisateur(LoadedDB, username)
	if err != nil {
		log.Fatalf("Erreur :", err)
	}

	// Check if the username & password exist to set the cookie to the user and check if the username & password
	// are the ones in the database
	if username != "" && password != "" && userData != nil {
		if username == userData.NomUtilisateur && password == userData.MotDePasse && sessErr != nil {
			UserSession[uuid.String()] = true
			http.SetCookie(w, &http.Cookie{
				Name:   "adminSession",
				Value:  uuid.String(),
				Path:   "/",
				MaxAge: 1800, // Set the expiration time to 1800 seconds (30 minutes)
			})
			// Redirect once the user is logged as an admin
			http.Redirect(w, r, "http://localhost:8080", http.StatusSeeOther)
		}
	}

	if sessErr == nil {
		// Delete the cookie if the user doesn't exist in the database
		if !UserSession[userCookie.Value] {
			// Delete the cookie for the user if it doesn't exist in the map
			http.SetCookie(w, &http.Cookie{
				Name:   "adminSession",
				Value:  userCookie.Value,
				Path:   "/",
				MaxAge: -1, // Set the time to -1
			})
		}
		// If the user is logged redirect him to the home page
		if UserSession[userCookie.Value] {
			http.Redirect(w, r, "http://localhost:8080", http.StatusSeeOther)
		}
	}

	err = tmpl.Execute(w, UserSession)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
