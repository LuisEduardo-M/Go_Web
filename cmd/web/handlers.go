package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/LuisEduardo-M/Go_Web/internal/models"
	"github.com/julienschmidt/httprouter"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	games, err := app.games.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Games = games

	app.render(w, http.StatusOK, "home.tmpl.html", data)

}

func (app *application) gameView(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	game, err := app.games.Get(id)
	if err != nil {
		if err == models.ErrNoRecord {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Game = game

	app.render(w, http.StatusOK, "view.tmpl.html", data)

}

func (app *application) gameAdd(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	app.render(w, http.StatusOK, "add.tmpl.html", data)
}

type gameAddForm struct {
	Title       string
	Description string
	Categories  string
	FieldErrors map[string]string
}

func (app *application) gameAddPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := gameAddForm{
		Title:       r.PostForm.Get("title"),
		Description: r.PostForm.Get("description"),
		Categories:  r.PostForm.Get("categories"),
		FieldErrors: map[string]string{},
	}

	if strings.TrimSpace(form.Title) == "" {
		form.FieldErrors["title"] = "Title is required"
	} else if len(form.Title) > 120 {
		form.FieldErrors["title"] = "Title must be less than 120 characters"
	}
	if strings.TrimSpace(form.Description) == "" {
		form.FieldErrors["description"] = "Description is required"
	}
	if strings.TrimSpace(form.Categories) == "" {
		form.FieldErrors["categories"] = "Categories is required"
	}

	if len(form.FieldErrors) > 0 {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, http.StatusUnprocessableEntity, "add.tmpl.html", data)
		return
	}

	id, err := app.games.Insert(form.Title, form.Description, form.Categories)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/game/view/%d", id), http.StatusSeeOther)
}
