package database

import "github.com/jmoiron/sqlx"

type Database interface {
	Database() *sqlx.DB
	MigrateUp() error
	MigrateDown() error
}
