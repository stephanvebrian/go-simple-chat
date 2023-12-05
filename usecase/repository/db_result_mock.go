package repository

import (
	"github.com/stretchr/testify/mock"
)

// MockResult is a mock implementation of the Result interface
type MockResult struct {
	mock.Mock
}

// LastInsertId mocks the LastInsertId method of Result interface
func (m *MockResult) LastInsertId() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}

// RowsAffected mocks the RowsAffected method of Result interface
func (m *MockResult) RowsAffected() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}
