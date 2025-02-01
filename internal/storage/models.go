package storage

import (
	"database/sql"
)

type Movie struct {
	Name        string         `db:"name"`
	Trailer     sql.NullString `db:"trailer"`
	DiscordUser sql.NullString `db:"user_name"`
}
type Game struct {
	Name string
}

type MovieList struct {
	List []Movie
}
type GameList struct {
	List []Game
}
