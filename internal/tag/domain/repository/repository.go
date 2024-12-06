package repository

import (
	tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
	"github.com/jambo0624/blog/internal/tag/domain/query"
)

type TagRepository interface {
	Save(tag *tagEntity.Tag) error
	FindByID(id uint) (*tagEntity.Tag, error)
	FindAll(query *query.TagQuery) ([]*tagEntity.Tag, error)
	Update(tag *tagEntity.Tag) error
	Delete(id uint) error
}
