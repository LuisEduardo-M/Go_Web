package main

import (
	"fmt"
	"github.com/LuisEduardo-M/Go_Web/internal/models"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	games, err := app.games.GetAll()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, game := range games {
		fmt.Fprintf(w, "%+v\n", game)
	}

}

func (app *application) gameView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
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

	fmt.Fprintf(w, "%+v", game)
}

func (app *application) gameAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "Valorant"
	description := "Valorant is a free-to-play, team-based, competitive game."
	categories := "FPS, Action, Multiplayer"

	id, err := app.games.Insert(title, description, categories)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/game?id=%d", id), http.StatusSeeOther)
}
