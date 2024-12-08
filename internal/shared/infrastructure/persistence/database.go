package persistence

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	articleEntity "github.com/jambo0624/blog/internal/article/domain/entity"
	categoryEntity "github.com/jambo0624/blog/internal/category/domain/entity"
	config "github.com/jambo0624/blog/internal/shared/infrastructure/config"
	tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
)

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.Database.URL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// only auto migrate in non-production environment
	if cfg.Environment != "production" {
		err = db.AutoMigrate(
			&articleEntity.Article{},
			&categoryEntity.Category{},
			&tagEntity.Tag{},
		)
		if err != nil {
			return nil, fmt.Errorf("failed to auto migrate: %w", err)
		}
	}

	return db, nil
}
