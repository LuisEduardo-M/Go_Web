package main

import (
	"fmt"
	"net/http"
	"strconv"

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
	w.Write([]byte("Display a form to create a new game"))
}

func (app *application) gameAddPost(w http.ResponseWriter, r *http.Request) {
	title := "Valorant"
	description := "Valorant is a free-to-play, team-based, competitive game."
	categories := "FPS, Action, Multiplayer"

	id, err := app.games.Insert(title, description, categories)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/game/view/%d", id), http.StatusSeeOther)
}
