package repository

// DBRowsInterface defines the methods expected from sql.Rows
type DBRowsInterface interface {
	Scan(dest ...interface{}) error
	Next() bool
	Close() error
}

// MockDBRows is a mock implementation of DBRowsInterface
type MockDBRows struct {
	ScanFunc  func(dest ...interface{}) error
	NextFunc  func() bool
	CloseFunc func() error
}

// Scan implements the Scan method of DBRowsInterface
func (m *MockDBRows) Scan(dest ...interface{}) error {
	if m.ScanFunc != nil {
		return m.ScanFunc(dest...)
	}
	return nil
}

// Next implements the Next method of DBRowsInterface
func (m *MockDBRows) Next() bool {
	if m.NextFunc != nil {
		return m.NextFunc()
	}
	return false
}

// Close implements the Close method of DBRowsInterface
func (m *MockDBRows) Close() error {
	if m.CloseFunc != nil {
		return m.CloseFunc()
	}
	return nil
}
