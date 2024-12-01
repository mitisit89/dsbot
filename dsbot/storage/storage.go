package storage

import "context"

type Movie struct {
	Names []string
}
type Storage interface {
	Add(ctx context.Context, m *Movie) error
	MarkedAsWatched(ctx context.Context, m *Movie) error
	GetAll(ctx context.Context) (*Movie, error)
}
