package repository

import (
	categoryEntity "github.com/jambo0624/blog/internal/category/domain/entity"
)

type CategoryRepository interface {
	Save(category *categoryEntity.Category) error
	FindByID(id uint) (*categoryEntity.Category, error)
	FindAll() ([]*categoryEntity.Category, error)
	Update(category *categoryEntity.Category) error
	Delete(id uint) error
}
