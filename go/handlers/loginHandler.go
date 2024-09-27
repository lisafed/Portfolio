package handlers

import (
	"fmt"
	"github.com/gofrs/uuid"
	"html/template"
	"log"
	"net/http"
	"portfolio/go/dbFunc"
)

var UserSession = make(map[string]bool)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./src/html/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userCookie, sessErr := r.Cookie("adminSession")
	fmt.Println(err)

	uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	LoadedDB := dbFunc.LoadDb()

	userData, err := dbFunc.GetUtilisateur(LoadedDB, username)
	if err != nil {
		log.Fatalf("Erreur :", err)
	}

	if username != "" && password != "" && userData != nil {
		if username == userData.NomUtilisateur && password == userData.MotDePasse && sessErr != nil {
			UserSession[uuid.String()] = true
			http.SetCookie(w, &http.Cookie{
				Name:   "adminSession",
				Value:  uuid.String(),
				Path:   "/",
				MaxAge: 1800, // Set the expiration time to 1800 seconds (30 minutes)
			})
			http.Redirect(w, r, "http://localhost:8080", http.StatusSeeOther)
		}
	}

	if sessErr == nil {
		if !UserSession[userCookie.Value] {
			http.SetCookie(w, &http.Cookie{
				Name:   "adminSession",
				Value:  userCookie.Value,
				Path:   "/",
				MaxAge: -1, // Set the expiration time to 1800 seconds (30 minutes)
			})
		}
		if UserSession[userCookie.Value] {
			fmt.Println(UserSession[userCookie.Value] == true)
			http.Redirect(w, r, "http://localhost:8080", http.StatusSeeOther)
		}

	}

	err = tmpl.Execute(w, UserSession)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
