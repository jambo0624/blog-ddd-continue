package persistence

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    config "github.com/jambo0624/blog/internal/shared/infrastructure/config"
    articleEntity "github.com/jambo0624/blog/internal/article/domain/entity"
    categoryEntity "github.com/jambo0624/blog/internal/category/domain/entity"
    tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
)

func InitDB(cfg *config.Config) (*gorm.DB, error) {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
        cfg.Database.Host,
        cfg.Database.User,
        cfg.Database.Password,
        cfg.Database.DBName,
        cfg.Database.Port,
        cfg.Database.SSLMode,
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %w", err)
    }

    // 自动迁移
    err = db.AutoMigrate(
        &articleEntity.Article{},
        &categoryEntity.Category{},
        &tagEntity.Tag{},
    )
    if err != nil {
        return nil, fmt.Errorf("failed to auto migrate: %w", err)
    }

    return db, nil
} 