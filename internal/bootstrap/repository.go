package bootstrap

import (
	"gorm.io/gorm"

	articleRepository "github.com/jambo0624/blog/internal/article/domain/repository"
	categoryRepository "github.com/jambo0624/blog/internal/category/domain/repository"
	tagRepository "github.com/jambo0624/blog/internal/tag/domain/repository"

	articlePersistence "github.com/jambo0624/blog/internal/article/infrastructure/persistence"
	categoryPersistence "github.com/jambo0624/blog/internal/category/infrastructure/persistence"
	tagPersistence "github.com/jambo0624/blog/internal/tag/infrastructure/persistence"
)

type Repositories struct {
	Article  articleRepository.ArticleRepository
	Category categoryRepository.CategoryRepository
	Tag      tagRepository.TagRepository
}

func SetupRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Article:  articlePersistence.NewGormArticleRepository(db),
		Category: categoryPersistence.NewGormCategoryRepository(db),
		Tag:      tagPersistence.NewGormTagRepository(db),
	}
} 