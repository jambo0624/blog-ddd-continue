package persistence

import (
	"gorm.io/gorm"

	persistence "github.com/jambo0624/blog/internal/shared/infrastructure/persistence"
	articleEntity "github.com/jambo0624/blog/internal/article/domain/entity"
	articleQuery "github.com/jambo0624/blog/internal/article/domain/query"
	articleRepository "github.com/jambo0624/blog/internal/article/domain/repository"
)

type GormArticleRepository struct {
	*persistence.BaseGormRepository[articleEntity.Article, *articleQuery.ArticleQuery]
}

func NewGormArticleRepository(db *gorm.DB) articleRepository.ArticleRepository {
	return &GormArticleRepository{
		BaseGormRepository: persistence.NewBaseGormRepository[articleEntity.Article, *articleQuery.ArticleQuery](db),
	}
}
