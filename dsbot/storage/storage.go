package storage

import (
	"context"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgconn"
)

type Movie struct {
	Names []string
}
type Games struct {
	Names []string
}

type Storage interface {
	Add(ctx context.Context, m string) error
	MarkedAsWatched(ctx context.Context, m string) error
	GetAll(ctx context.Context) (*Movie, error)
}

type DBTX interface {
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

type Queries struct {
	db DBTX
}

func (q *Queries) WithTx(tx pgx.Tx) *Queries {
	return &Queries{
		db: tx,
	}
}
