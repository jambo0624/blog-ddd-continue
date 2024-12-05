package repository

import (
	tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
)

type TagRepository interface {
	Save(tag *tagEntity.Tag) error
	FindByID(id uint) (*tagEntity.Tag, error)
	FindAll() ([]*tagEntity.Tag, error)
	Update(tag *tagEntity.Tag) error
	Delete(id uint) error
}
