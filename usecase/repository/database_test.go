package repository

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

const (
	testDBName = "test_database.sqlite"
)

func setupTestDB(t *testing.T) (*sql.DB, func()) {
	t.Helper()

	// Open the test database
	db, err := sql.Open("sqlite3", testDBName)
	if err != nil {
		t.Fatalf("Error opening test database: %v", err)
	}

	// Run migration on the test database
	RunMigration(db)

	// Return a cleanup function to close the test database
	cleanup := func() {
		db.Close()
		os.Remove(testDBName)
	}

	return db, cleanup
}

func TestNewDatabase(t *testing.T) {
	// Setup test database and cleanup function
	_, cleanup := setupTestDB(t)
	defer cleanup()

	// Call NewDatabase
	_, err := NewDatabase()
	if err != nil {
		t.Fatalf("Error creating new database: %v", err)
	}

	// You can add more assertions or checks as needed
}

func TestRunMigration(t *testing.T) {
	// Setup test database and cleanup function
	db, cleanup := setupTestDB(t)
	defer cleanup()

	// Run migration on the test database
	RunMigration(db)
}
