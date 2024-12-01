package dsbot

import (
	"log/slog"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func connect() *sqlx.DB {
	db, err := sqlx.Connect(os.Getenv("DB_TYPE"), os.Getenv("DB_NAME"))
	if err != nil {
		slog.Error("failed to connect database")
		os.Exit(1)
	}

	return db
}

func getAll(db *sqlx.DB) (watchlist []string) {
	err := db.Select(&watchlist, "SELECT name FROM watchlist WHERE watched=0")
	if err != nil {
		panic(err)
	}
	return watchlist
}

func add(db *sqlx.DB, movie string) {
	_, err := db.Exec("INSERT INTO watchlist (name) VALUES (?)", movie)
	if err != nil {
		panic(err)
	}
}

func remove(db *sqlx.DB, movie string) {
	_, err := db.Exec("update watchlist set watched=1 where name=?", movie)
	if err != nil {
		panic(err)
	}
}
