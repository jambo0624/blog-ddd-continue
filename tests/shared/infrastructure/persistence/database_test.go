package persistence_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/jambo0624/blog/internal/shared/infrastructure/config"
	"github.com/jambo0624/blog/internal/shared/infrastructure/persistence"
)

func TestInitDB(t *testing.T) {
	cfg, err := config.LoadConfig()
	assert.NoError(t, err)

	db, err := persistence.InitDB(cfg)
	assert.NoError(t, err)
	assert.NotNil(t, db)

	sqlDB, err := db.DB()
	assert.NoError(t, err)
	assert.NoError(t, sqlDB.Ping())
} 