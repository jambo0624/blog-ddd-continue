package repository

import (
	"github.com/jambo0624/blog/internal/shared/domain/repository"
	tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
	tagQuery "github.com/jambo0624/blog/internal/tag/domain/query"
)

type TagRepository interface {
	repository.BaseRepository[tagEntity.Tag, *tagQuery.TagQuery]
}
