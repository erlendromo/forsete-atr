package database

import (
	"github.com/jmoiron/sqlx"
)

type Database interface {
	DB() *sqlx.DB
	MigrateUp() error
	MigrateDown() error
}
