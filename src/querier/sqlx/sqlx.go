package sqlx

import (
	"context"
	"fmt"

	"github.com/erlendromo/forsete-atr/src/database"
)

type SqlxQuerier[T any] struct {
	db database.Database
}

func NewSqlxQuerier[T any](db database.Database) *SqlxQuerier[T] {
	return &SqlxQuerier[T]{
		db: db,
	}
}

func (q *SqlxQuerier[T]) QueryRowx(ctx context.Context, query string, args ...interface{}) (*T, error) {
	var result T
	row := q.db.Database().QueryRowxContext(ctx, query, args...)
	if err := row.StructScan(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (q *SqlxQuerier[T]) Queryx(ctx context.Context, query string, args ...interface{}) ([]*T, error) {
	rows, err := q.db.Database().QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*T
	for rows.Next() {
		var item T
		if err := rows.StructScan(&item); err != nil {
			return nil, err
		}
		results = append(results, &item)
	}

	return results, nil
}

func (q *SqlxQuerier[T]) Executex(ctx context.Context, query string, args ...interface{}) error {
	result, err := q.db.Database().ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("\nExecuted query on %d rows.\n", rowsAffected)

	return nil
}
