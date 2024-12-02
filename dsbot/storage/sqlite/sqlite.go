package storage

import (
	"context"
	"database/sql"
	"dsbot/dsbot/storage"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

// New create new storage
func New() *Storage {
	db, err := sql.Open("sqlite3", os.Getenv("DB_PATH"))
	if err != nil {
		panic(fmt.Errorf("storage.New: cannot open database %w", err))
	}
	if err := db.Ping(); err != nil {
		panic(fmt.Errorf("storage.New: failed to connect database %w", err))
	}
	return &Storage{db: db}
}

// Add add movie to database
func (s *Storage) Add(ctx context.Context, movie string) error {
	query := `INSERT INTO watchlist (name) VALUES (?)`
	if _, err := s.db.ExecContext(ctx, query, movie); err != nil {
		return fmt.Errorf("storage.Add: failed to add movie %w", err)
	}
	return nil
}

// GetAll get all movies
func (s *Storage) GetAll(ctx context.Context) (*storage.Movie, error) {
	query := `SELECT name FROM watchlist WHERE watched=0`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("storage.GetAll: failed to get all movies %w", err)
	}
	defer rows.Close()
	var moviesNames []string
	for rows.Next() {
		var movie string
		if err := rows.Scan(&movie); err != nil {
			return nil, fmt.Errorf("storage.GetAll: failed to scan movie %w", err)
		}
		moviesNames = append(moviesNames, movie)
	}

	return &storage.Movie{Names: moviesNames}, nil
}

// MarkAsWatched mark movie as watched
func (s *Storage) MarkAsWatched(ctx context.Context, movie string) error {
	query := `update watchlist set watched=1 where name=?`
	if _, err := s.db.ExecContext(ctx, query, movie); err != nil {
		return fmt.Errorf("storage.MarkAsWatched: failed to remove movie %w", err)
	}
	return nil
}

// func (s *Storage) Init(ctx context.Context) error {
// 	query := `create table if not exists watchlist (
//             id integer primary key autoincrement,
//             name string not null unique,
//             wattched integer default 0);`
//
// 	if _, err := s.db.ExecContext(ctx, query); err != nil {
// 		return fmt.Errorf("failed to create table %w", err)
// 	}
// 	return nil
//
// }
