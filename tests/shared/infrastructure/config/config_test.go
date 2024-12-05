package config_test

import (
    "testing"
    "os"
    "github.com/stretchr/testify/assert"
    "github.com/jambo0624/blog/internal/shared/infrastructure/config"
)

func TestLoadConfig(t *testing.T) {
    os.Setenv("GO_ENV", "test")
    
    cfg, err := config.LoadConfig()
    
    assert.NoError(t, err)
    assert.NotNil(t, cfg)
    assert.NotEmpty(t, cfg.Database.URL)
    assert.NotEmpty(t, cfg.Server.Port)
}

func TestParseDatabaseURL(t *testing.T) {
    dbURL := "postgres://user:pass@localhost:5432/testdb"
    
    config := config.ParseDatabaseURL(dbURL)
    
    assert.Equal(t, "localhost", config.Host)
    assert.Equal(t, "5432", config.Port)
    assert.Equal(t, "user", config.User)
    assert.Equal(t, "pass", config.Password)
    assert.Equal(t, "testdb", config.DBName)
} 