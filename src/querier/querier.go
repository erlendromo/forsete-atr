package querier

import (
	"context"
)

type Querier[T any] interface {
	QueryRowx(ctx context.Context, query string, args ...interface{}) (*T, error)
	Queryx(ctx context.Context, query string, args ...interface{}) ([]*T, error)
	Executex(ctx context.Context, query string, args ...interface{}) error
}
