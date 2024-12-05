package bootstrap

import (
	config "github.com/jambo0624/blog/internal/shared/infrastructure/config"
	persistence "github.com/jambo0624/blog/internal/shared/infrastructure/persistence"
	"gorm.io/gorm"
)

func SetupDB(cfg *config.Config) (*gorm.DB, error) {
	return persistence.InitDB(cfg)
} 