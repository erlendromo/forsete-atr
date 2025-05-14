package postgresql

import (
	"fmt"

	"github.com/erlendromo/forsete-atr/src/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type PostgreSQLDatabase struct {
	db *sqlx.DB
}

func NewPostgreSQLDatabase() *PostgreSQLDatabase {
	dbConfig := config.GetConfig().DBConfig()
	dataSourceName := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.DB_HOST,
		dbConfig.DB_PORT,
		dbConfig.DB_USER,
		dbConfig.DB_PASS,
		dbConfig.DB_NAME,
	)

	psql := &PostgreSQLDatabase{
		db: sqlx.MustConnect("postgres", dataSourceName),
	}

	if err := psql.MigrateUp(); err != nil {
		fmt.Println("Unable to migrate database-schema:\n" + err.Error())
	}

	return psql
}

func (p *PostgreSQLDatabase) Database() *sqlx.DB {
	return p.db
}

func (p *PostgreSQLDatabase) MigrateUp() error {
	fmt.Println("Database migrating up...")

	m, err := p.getMigrate()
	if err != nil {
		return fmt.Errorf("getting migration: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("running migrate up: %w", err)
	}

	return nil
}

func (p *PostgreSQLDatabase) MigrateDown() error {
	fmt.Println("Database migrating down...")

	m, err := p.getMigrate()
	if err != nil {
		return fmt.Errorf("getting migration: %w", err)
	}

	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("running migrate down: %w", err)
	}

	return nil
}

func (p *PostgreSQLDatabase) getMigrate() (*migrate.Migrate, error) {
	instance, err := postgres.WithInstance(p.db.DB, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance("file:///assets/migrations", "postgres", instance)
	if err != nil {
		return nil, err
	}

	return m, nil
}
