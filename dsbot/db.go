package dsbot

import (
	"log/slog"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func connect() *sqlx.DB {
	db, err := sqlx.Connect("sqlite", "watchlist.db")
	if err != nil {
		slog.Error("failed to connect database")
		os.Exit(1)
	}

	return db
}

func getAll(db *sqlx.DB) (watchlist []string) {
	err := db.Select(&watchlist, "SELECT * FROM watchlist")
	if err != nil {
		panic(err)
	}
	return watchlist
}

func add(db *sqlx.DB, movie string) {
	_, err := db.Exec("INSERT INTO watchlist (movie) VALUES (?)", movie)
	if err != nil {
		panic(err)
	}
}

func remove(db *sqlx.DB, movie string) {
	_, err := db.Exec("DELETE FROM watchlist WHERE movie = ?", movie)
	if err != nil {
		panic(err)
	}
}
