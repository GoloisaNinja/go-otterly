package main

import (
	"github.com/GoloisaNinja/go-otterly/pkg/config"
	"github.com/GoloisaNinja/go-otterly/pkg/handlers"
	"github.com/GoloisaNinja/go-otterly/pkg/helpers"
	"github.com/GoloisaNinja/go-otterly/pkg/render"
	"log"
	"net/http"
)

var app config.AppConfig
var gc config.GameConfig

func main() {

	port := ":8000"
	app.IsProduction = false
	helpers.LoadGames(&gc)
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("could not create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepository(&app, &gc)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
