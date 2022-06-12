package handlers

import (
	"github.com/LuisEduardo-M/Go_Web/pkg/config"
	"github.com/LuisEduardo-M/Go_Web/pkg/models"
	"github.com/LuisEduardo-M/Go_Web/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the new handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.tmpl.html", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again..."

	render.RenderTemplate(w, "about.tmpl.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
