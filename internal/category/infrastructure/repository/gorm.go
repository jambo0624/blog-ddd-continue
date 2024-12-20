package persistence

import (
	"gorm.io/gorm"

	categoryEntity "github.com/jambo0624/blog/internal/category/domain/entity"
	categoryQuery "github.com/jambo0624/blog/internal/category/domain/query"
	categoryRepository "github.com/jambo0624/blog/internal/category/domain/repository"
	persistence "github.com/jambo0624/blog/internal/shared/infrastructure/persistence"
)

type GormCategoryRepository struct {
	*persistence.BaseGormRepository[categoryEntity.Category, *categoryQuery.CategoryQuery]
}

func NewGormCategoryRepository(db *gorm.DB) categoryRepository.CategoryRepository {
	return &GormCategoryRepository{
		BaseGormRepository: persistence.NewBaseGormRepository[categoryEntity.Category, *categoryQuery.CategoryQuery](db),
	}
}
