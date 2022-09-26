package main

import (
	"fmt"
	"github.com/nambroa/go-mock-project/pkg/config"
	"github.com/nambroa/go-mock-project/pkg/handlers"
	render "github.com/nambroa/go-mock-project/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = templateCache
	app.UseCache = false
	//
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Starting application on port", portNumber)
	// Start a webserver and listen to a specific port.
	_ = http.ListenAndServe(portNumber, nil)
}
