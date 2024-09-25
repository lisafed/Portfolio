package main

import (
	"fmt"
	"net/http"
	"portfolio/go/handlers"
)

type webserver struct {
	core *http.ServeMux
	port int
}

func main() {
	server := webserver{
		core: http.NewServeMux(),
		port: 8080,
	}

	server.Router()
	server.Launch()
}

func (s *webserver) Router() {
	s.core.HandleFunc("/index", handlers.IndexHandler)

	s.core.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./src/css"))))
	s.core.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./src/img"))))
	s.core.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./src/js"))))
	fmt.Printf("http://localhost:%d \n", s.port)
}

func (s *webserver) Launch() {
	// Lancement du serveur HTTP sur le port spécifié
	http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.core)
}
