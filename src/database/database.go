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

// Using generics,
// query a row in the database and return the data of the selected generic type,
// assuming no errors.
// Otherwise the error is returned.
func QueryRowx[T any](ctx context.Context, db *sqlx.DB, query string, args ...interface{}) (*T, error) {
	var result T
	row := db.QueryRowxContext(ctx, query, args...)
	if err := row.StructScan(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// Using generics,
// query multiple rows in the database and return the data in a list of the selected generic type,
// assuming no errors.
// Otherwise the error is returned.
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

// Returns the affected rows after the query is run, assuming no errors.
// Otherwise the error is returned.
func ExecuteContext(ctx context.Context, db *sqlx.DB, query string, args ...interface{}) error {
	result, err := db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("\nUpdated/Deleted %d rows.\n", rowsAffected)

	return nil
}
