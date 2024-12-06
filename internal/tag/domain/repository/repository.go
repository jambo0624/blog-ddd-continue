package repository

import (
	tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
	tagQuery "github.com/jambo0624/blog/internal/tag/domain/query"
	"github.com/jambo0624/blog/internal/shared/domain/repository"
)

type TagRepository interface {
	repository.BaseRepository[tagEntity.Tag, *tagQuery.TagQuery]
}
