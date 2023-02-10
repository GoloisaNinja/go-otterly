package main

import (
	"github.com/GoloisaNinja/go-otterly/pkg/config"
	"github.com/GoloisaNinja/go-otterly/pkg/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	r := mux.NewRouter()
	// routes
	r.HandleFunc("/", handlers.Repo.Home)
	r.HandleFunc("/about", handlers.Repo.About)
	r.HandleFunc("/contact", handlers.Repo.Contact)
	r.HandleFunc("/games", handlers.Repo.Games)
	r.HandleFunc("/games/{id}", handlers.Repo.Game)
	// api routes
	r.HandleFunc("/api/gameNode", handlers.Repo.GetNode).Methods("POST")
	// load static assets
	staticFileDir := http.Dir("static")
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDir))
	r.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")
	// catch all
	r.PathPrefix("/").HandlerFunc(handlers.Repo.Home)
	return r
}
