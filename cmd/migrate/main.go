package main

import (
    "log"
    "fmt"

    config "github.com/jambo0624/blog/internal/shared/infrastructure/config"
    articleEntity "github.com/jambo0624/blog/internal/article/domain/entity"
    categoryEntity "github.com/jambo0624/blog/internal/category/domain/entity"
    tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
)

func main() {
    // 加载配置
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    // 连接数据库
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
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // 执行迁移
    err = db.AutoMigrate(
        &articleEntity.Article{},
        &categoryEntity.Category{},
        &tagEntity.Tag{},
    )
    if err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }

    log.Println("Migration completed successfully")
} 