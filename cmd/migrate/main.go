package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	articleEntity "github.com/jambo0624/blog/internal/article/domain/entity"
	categoryEntity "github.com/jambo0624/blog/internal/category/domain/entity"
	config "github.com/jambo0624/blog/internal/shared/infrastructure/config"
	tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
)

func main() {
	// load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// connect to database
	db, err := gorm.Open(postgres.Open(cfg.Database.URL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// execute migrations
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
