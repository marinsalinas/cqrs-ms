package search

import (
	"context"

	"github.com/marinsalinas/cqrs-ms/schema"
)

//Repository ...
type Repository interface {
	Close()
	InsertMeow(ctx context.Context, meow schema.Meow) error
	SearchMeows(ctx context.Context, query string, skip uint64, take uint64) ([]schema.Meow, error)
}

var impl Repository

//SetRepository ...
func SetRepository(repository Repository) {
	impl = repository
}

//Close ...
func Close() {
	impl.Close()
}

//InsertMeow ...
func InsertMeow(ctx context.Context, meow schema.Meow) error {
	return impl.InsertMeow(ctx, meow)
}

//SearchMeows ...
func SearchMeows(ctx context.Context, query string, skip uint64, take uint64) ([]schema.Meow, error) {
	return impl.SearchMeows(ctx, query, skip, take)
}
