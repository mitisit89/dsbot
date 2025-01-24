package pg

import (
	"context"
	"dsbot/dsbot/storage"
	"fmt"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	db *pgxpool.Pool
}

func New() *Storage {
	url := "postgres://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_HOST") + "/" + os.Getenv("DB_NAME")
	db, err := pgxpool.New(context.Background(), url)
	if err != nil {
		slog.Error("storage.New: cannot open database %w", slog.String("error", err.Error()))
		os.Exit(1)
	}
	// defer db.Close()
	return &Storage{db: db}
}
func (s *Storage) AddUser(ctx context.Context, dsUser string) error {

	q := `
    INSERT INTO discord_user (name)
    SELECT $1
    WHERE NOT EXISTS (
        SELECT 1 FROM discord_user WHERE name = $2
    );
    `
	if _, err := s.db.Exec(ctx, q, dsUser, dsUser); err != nil {
		return fmt.Errorf("storage.Add: failed to add user%w", err)
	}
	return nil
}
func (s *Storage) Add(ctx context.Context, dsUser string, args []string) error {
	movie := args[0]
	trailer := ""
	if len(args) > 1 {
		trailer = args[1]
	}
	if err := s.AddUser(ctx, dsUser); err != nil {
		return err
	}
	q := `INSERT INTO movies (name, trailer, discord_user_id)
    VALUES (
        @movie,
        @trailer,
        (SELECT id FROM discord_user WHERE name = @dsUser)
    );
    `

	if _, err := s.db.Exec(ctx, q, pgx.NamedArgs{"movie": movie, "trailer": trailer, "dsUser": dsUser}); err != nil {
		return fmt.Errorf("storage.Add: failed to add movie %w", err)
	}
	defer s.db.Close()
	return nil
}

// GetAll get all movies
func (s *Storage) GetAll(ctx context.Context) (*storage.Movie, error) {
	// TODO: update struct
	// 	query := `SELECT
	//     discord_user.name AS user_name,
	//     movies.name AS movie_name
	// FROM
	//     movies
	// left join
	//     discord_user ON movies.discord_user_id = discord_user.id
	// WHERE
	//     movies.watched = FALSE;`

	query := `SELECT name FROM movies WHERE watched=false`
	rows, err := s.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("storage.GetAll: failed to get all movies %w", err)
	}
	defer rows.Close()
	var moviesNames []string
	for rows.Next() {
		var movie string
		// var user string
		if err := rows.Scan(&movie); err != nil {
			return nil, fmt.Errorf("storage.GetAll: failed to scan movie %w", err)
		}
		moviesNames = append(moviesNames, movie)
	}

	return &storage.Movie{Names: moviesNames}, nil
}

// // MarkAsWatched mark movie as watched
func (s *Storage) MarkAsWatched(ctx context.Context, movie string) error {
	query := `update movies set watched=true where name=@movie`
	if _, err := s.db.Exec(ctx, query, pgx.NamedArgs{"movie": movie}); err != nil {
		return fmt.Errorf("storage.MarkAsWatched: failed to remove movie %w", err)
	}
	return nil
}
