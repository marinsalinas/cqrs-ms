package db

import (
	"context"

	"github.com/marinsalinas/cqrs-ms/schema"
)

//Repository ...
type Repository interface {
	Close()
	InsertMeow(ctx context.Context, meow schema.Meow) error
	ListMeows(ctx context.Context, skip uint64, take uint64) ([]schema.Meow, error)
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

//ListMeows ...
func ListMeows(ctx context.Context, skip uint64, take uint64) ([]schema.Meow, error) {
	return impl.ListMeows(ctx, skip, take)
}
