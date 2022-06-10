package main

import (
	"fmt"
	"github.com/LuisEduardo-M/Go_Web/pkg/config"
	"github.com/LuisEduardo-M/Go_Web/pkg/handlers"
	"github.com/LuisEduardo-M/Go_Web/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on http://localhost%s/\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
