package models

import (
	"database/sql"
	"errors"
)

// Game struct to hold games data
type Game struct {
	ID          int
	Title       string
	Description string
	Categories  string
}

// GameModel type wraps a sql.DB connection pool
type GameModel struct {
	DB *sql.DB
}

/*
DB.Query() - Used for SELECT queries which return multiple rows.
DB.QueryRow() - Used for SELECT queries which return a single row.
DB.Exec() - Used for statements which don't return rows (like INSERT and DELETE).
*/

// Insert will insert a new game into the database
func (g *GameModel) Insert(title, description, categories string) (int, error) {
	stmt := `INSERT INTO games (title, description, categories) VALUES (?, ?, ?)`

	result, err := g.DB.Exec(stmt, title, description, categories)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Get will return a game from the database
func (g *GameModel) Get(id int) (*Game, error) {
	stmt := `SELECT id, title, description, categories FROM games WHERE id = ?`

	row := g.DB.QueryRow(stmt, id)

	game := &Game{}

	err := row.Scan(&game.ID, &game.Title, &game.Description, &game.Categories)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return game, nil
}

// GetAll will return all games from the database
func (g *GameModel) GetAll() ([]*Game, error) {
	stmt := `SELECT id, title, description, categories FROM games`

	rows, err := g.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	var games []*Game

	for rows.Next() {
		game := &Game{}

		err := rows.Scan(&game.ID, &game.Title, &game.Description, &game.Categories)
		if err != nil {
			return nil, err
		}

		games = append(games, game)
	}

	return games, nil
}
