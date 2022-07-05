package main

import (
	"path/filepath"
	"text/template"

	"github.com/LuisEduardo-M/Go_Web/internal/models"
)

// templateData type holds structure for any dynamic data that needs to be passed to a HTML template.
type templateData struct {
	Game  *models.Game
	Games []*models.Game
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		files := []string{
			"./ui/html/base.tmpl.html",
			"./ui/html/partials/nav.tmpl.html",
			page,
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
