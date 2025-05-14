package mock

import (
	"context"

	"github.com/erlendromo/forsete-atr/src/database"
)

type MockQuerier[T any] struct{}

func NewMockQuerier[T any](db database.Database) *MockQuerier[T] {
	return &MockQuerier[T]{}
}

func (q *MockQuerier[T]) QueryRowx(ctx context.Context, query string, args ...interface{}) (*T, error) {
	var result T
	return &result, nil
}

func (q *MockQuerier[T]) Queryx(ctx context.Context, query string, args ...interface{}) ([]*T, error) {
	return make([]*T, 0), nil
}

func (q *MockQuerier[T]) Executex(ctx context.Context, query string, args ...interface{}) error {
	return nil
}
