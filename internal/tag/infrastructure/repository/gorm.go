package persistence

import (
	"gorm.io/gorm"

	persistence "github.com/jambo0624/blog/internal/shared/infrastructure/persistence"
	tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
	tagQuery "github.com/jambo0624/blog/internal/tag/domain/query"
	tagRepository "github.com/jambo0624/blog/internal/tag/domain/repository"
)

type GormTagRepository struct {
	*persistence.BaseGormRepository[tagEntity.Tag, *tagQuery.TagQuery]
}

func NewGormTagRepository(db *gorm.DB) tagRepository.TagRepository {
	return &GormTagRepository{
		BaseGormRepository: persistence.NewBaseGormRepository[tagEntity.Tag, *tagQuery.TagQuery](db),
	}
}
