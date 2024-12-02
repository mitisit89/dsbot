package storage

import "context"

type Movie struct {
	Names []string
}

type Storage interface {
	Add(ctx context.Context, m string) error
	MarkedAsWatched(ctx context.Context, m string) error
	GetAll(ctx context.Context) (*Movie, error)
}
