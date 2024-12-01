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
func New() (*Storage, error) {
	db, err := sql.Open("sqlite3", os.Getenv("DB_PATH"))
	if err != nil {
		return nil, fmt.Errorf("cannot open database %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database %w", err)
	}
	return &Storage{db: db}, nil
}

// Add add movie to database
func (s *Storage) Add(ctx context.Context, movie string) error {
	query := `INSERT INTO watchlist (name) VALUES (?)`
	if _, err := s.db.ExecContext(ctx, query, movie); err != nil {
		return fmt.Errorf("failed to add movie %w", err)
	}
	return nil
}

// GetAll get all movies
func (s *Storage) GetAll(ctx context.Context) (*storage.Movie, error) {
	var moviesNames []string
	query := `SELECT name FROM watchlist WHERE watched=0`
	if err := s.db.QueryRowContext(ctx, query).Scan(&moviesNames); err != nil {
		return nil, fmt.Errorf("failed to get all movies %w", err)
	}
	return &storage.Movie{Names: moviesNames}, nil
}

// MarkAsWatched mark movie as watched
func (s *Storage) MarkAsWatched(ctx context.Context, movie string) error {
	query := `update watchlist set watched=1 where name=?`
	if _, err := s.db.ExecContext(ctx, query, movie); err != nil {
		return fmt.Errorf("failed to remove movie %w", err)
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
