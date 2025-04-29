package database

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Database interface {
	Database() *sqlx.DB
	MigrateUp() error
	MigrateDown() error
}

func QueryRowx[T any](ctx context.Context, db *sqlx.DB, query string, args ...interface{}) (*T, error) {
	var result T
	row := db.QueryRowxContext(ctx, query, args...)
	if err := row.StructScan(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func Queryx[T any](ctx context.Context, db *sqlx.DB, query string, args ...interface{}) ([]*T, error) {
	rows, err := db.QueryxContext(ctx, query, args...)
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

func ExecuteContext(ctx context.Context, db *sqlx.DB, query string, args ...interface{}) (int, error) {
	result, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	} else if rowsAffected == 0 {
		return 0, fmt.Errorf("no change")
	}

	return int(rowsAffected), nil
}
