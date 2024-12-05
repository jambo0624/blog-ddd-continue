package testutil

import (
    "os"
    "testing"
    "gorm.io/gorm"
    "github.com/jambo0624/blog/internal/shared/infrastructure/config"
    "github.com/jambo0624/blog/internal/shared/infrastructure/persistence"
)

var TestDB *gorm.DB

func SetupTestDB(t *testing.T) func() {
    // 确保使用测试环境
    os.Setenv("GO_ENV", "test")
    
    cfg, err := config.LoadConfig()
    if err != nil {
        t.Fatalf("Failed to load test config: %v", err)
    }

    db, err := persistence.InitDB(cfg)
    if err != nil {
        t.Fatalf("Failed to initialize test database: %v", err)
    }

    TestDB = db

    return func() {
        sqlDB, err := db.DB()
        if err != nil {
            t.Errorf("Failed to get database instance: %v", err)
            return
        }
        sqlDB.Close()
    }
} 