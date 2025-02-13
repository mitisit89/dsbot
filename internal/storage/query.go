package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (q *Queries) AddUser(ctx context.Context, dsUser string) error {

	addUser := `
    INSERT INTO discord_user (name)
    SELECT $1
    WHERE NOT EXISTS (
        SELECT 1 FROM discord_user du WHERE du.name = $2
    );
    `
	if _, err := q.db.Exec(ctx, addUser, dsUser, dsUser); err != nil {
		return fmt.Errorf("storage.Add: failed to add user%w", err)
	}
	return nil
}
func (q *Queries) Add(ctx context.Context, dsUser string, args []string) error {
	movie := args[0]
	trailer := ""
	if len(args) > 1 {
		trailer = args[1]
	}
	if err := q.AddUser(ctx, dsUser); err != nil {
		return err
	}
	add := `INSERT INTO movies (name, trailer, discord_user_id)
    VALUES (
        @movie,
        @trailer,
        (SELECT id FROM discord_user du WHERE du.name = @dsUser)
    );
    `

	if _, err := q.db.Exec(ctx, add, pgx.NamedArgs{"movie": movie, "trailer": trailer, "dsUser": dsUser}); err != nil {
		return fmt.Errorf("storage.Add: failed to add movie %w", err)
	}
	return nil
}

// GetAll get all movies
func (q *Queries) GetAll(ctx context.Context) ([]Movie, error) {
	getAll := `SELECT
	    discord_user.name AS user_name,
	    movies.name AS movie_name,
        movies.trailer
	FROM
	    movies
	left join
	    discord_user ON movies.discord_user_id = discord_user.id
	WHERE
    movies.watched = FALSE order by movies.id limit 25;` //TODO: figure out pagination

	rows, err := q.db.Query(ctx, getAll)
	if err != nil {
		return nil, fmt.Errorf("storage.GetAll: failed to get all movies %w", err)
	}
	defer rows.Close()
	movies, err := pgx.CollectRows(rows, pgx.RowToStructByName[Movie])
	if err != nil {
		return nil, fmt.Errorf("storage.GetAll: failed to collect rows %w", err)
	}

	return movies, nil
}
func (q *Queries) AddGame(ctx context.Context, dsUser string, game string) error {
	if err := q.AddUser(ctx, dsUser); err != nil {
		return err
	}
	addGame := `INSERT INTO games (name, discord_user_id)
    VALUES (
        @game,
        (SELECT id FROM discord_user WHERE name = @dsUser)
    );
    `
	if _, err := q.db.Exec(ctx, addGame, pgx.NamedArgs{"game": game, "dsUser": dsUser}); err != nil {
		return fmt.Errorf("storage.Add: failed to add movie %w", err)
	}

	return nil

}
func (q *Queries) GameList(ctx context.Context) (*GameList, error) {

	gameList := `SELECT name FROM games`
	rows, err := q.db.Query(ctx, gameList)
	if err != nil {
		return nil, fmt.Errorf("storage.GetAll: failed to get all games %w", err)
	}
	defer rows.Close()
	var gameNames []Game
	for rows.Next() {
		var game Game
		if err := rows.Scan(&game); err != nil {
			return nil, fmt.Errorf("storage.GetAll: failed to scan movie %w", err)
		}
		gameNames = append(gameNames, game)
	}

	return &GameList{List: gameNames}, nil
}

// // MarkAsWatched mark movie as watched
func (s *Queries) MarkAsWatched(ctx context.Context, movie string) error {
	query := `update movies set watched=true where name=@movie`
	if _, err := s.db.Exec(ctx, query, pgx.NamedArgs{"movie": movie}); err != nil {
		return fmt.Errorf("storage.MarkAsWatched: failed to remove movie %w", err)
	}
	return nil
}
