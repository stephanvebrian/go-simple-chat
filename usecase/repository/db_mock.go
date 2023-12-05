package repository

import (
	"database/sql"

	"github.com/stretchr/testify/mock"
)

// MockDB is a mock implementation of DBInterface
type MockDB struct {
	mock.Mock
}

// Exec is a mocked implementation of sql.DB.Exec
func (m *MockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	// ArgumentsExpected is used to assert that the method was called with the expected arguments.
	// For simplicity, we're not handling arguments here. You can extend it based on your needs.
	args = append([]interface{}{query}, args...)
	result := m.Called(args...)
	return result.Get(0).(sql.Result), result.Error(1)
}

// Query is a mocked implementation of sql.DB.Query
func (m *MockDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	// ArgumentsExpected is used to assert that the method was called with the expected arguments.
	// For simplicity, we're not handling arguments here. You can extend it based on your needs.
	args = append([]interface{}{query}, args...)
	result := m.Called(args...)
	return result.Get(0).(*sql.Rows), result.Error(1)
}
