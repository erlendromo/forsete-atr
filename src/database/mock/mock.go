package mock

import (
	"github.com/jmoiron/sqlx"
)

type MockDatabase struct{}

func NewMockDatabase() *MockDatabase {
	return &MockDatabase{}
}

func (m *MockDatabase) DB() *sqlx.DB {
	return &sqlx.DB{}
}

func (m *MockDatabase) MigrateUp() error {
	return nil
}

func (m *MockDatabase) MigrateDown() error {
	return nil
}
