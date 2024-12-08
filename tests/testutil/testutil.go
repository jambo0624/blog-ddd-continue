package testutil

import (
	"fmt"
	"os"
	"sync"
	"testing"

	"gorm.io/gorm"

	"github.com/jambo0624/blog/internal/shared/infrastructure/config"
	"github.com/jambo0624/blog/internal/shared/infrastructure/errors"
	"github.com/jambo0624/blog/internal/shared/infrastructure/persistence"
	"github.com/jambo0624/blog/tests/testutil/fixtures"
)

var (
	testDB   *gorm.DB
	testData *fixtures.TestData
	once     sync.Once
)

type TestDB struct {
	*gorm.DB
	Data *fixtures.TestData
}

// initTestDB initializes test database connection and loads fixtures once.
func initTestDB() (*gorm.DB, *fixtures.TestData, error) {
	var err error

	once.Do(func() {
		os.Setenv("GO_ENV", "test")

		var cfg *config.Config
		cfg, err = config.LoadConfig()
		if err != nil {
			return
		}

		testDB, err = persistence.InitDB(cfg)
		if err != nil {
			return
		}

		// Clean database and load fixtures only once during initialization
		cleanDB(testDB)
		testData, err = fixtures.LoadFixtures(testDB)
		if err != nil {
			return
		}
	})

	if err != nil {
		return nil, nil, errors.ErrFailedToInitializeDB
	}

	if testDB == nil {
		return nil, nil, errors.ErrDBNotInitialized
	}

	return testDB, testData, nil
}

// SetupTestDB sets up the test database with transaction.
func SetupTestDB(t *testing.T) (*TestDB, func()) {
	t.Helper()

	db, data, err := initTestDB()
	if err != nil {
		t.Fatalf("Failed to initialize test database: %v", err)
	}

	if db == nil {
		t.Fatal("Database connection is nil")
	}

	// Begin transaction
	tx := db.Begin()
	if tx.Error != nil {
		t.Fatalf("Failed to begin transaction: %v", tx.Error)
	}

	// Return cleanup function that rolls back transaction
	cleanup := func() {
		tx.Rollback()
	}

	return &TestDB{DB: tx, Data: data}, cleanup
}

// cleanDB cleans the database.
func cleanDB(db *gorm.DB) {
	tables := []string{
		"article_tags",
		"articles",
		"categories",
		"tags",
	}
	for _, table := range tables {
		db.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", table))
	}
}
