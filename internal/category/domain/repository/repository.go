package repository

import (
	categoryEntity "github.com/jambo0624/blog/internal/category/domain/entity"
	"github.com/jambo0624/blog/internal/category/domain/query"
)

type CategoryRepository interface {
	Save(category *categoryEntity.Category) error
	FindByID(id uint) (*categoryEntity.Category, error)
	Update(category *categoryEntity.Category) error
	Delete(id uint) error
	FindAll(q *query.CategoryQuery) ([]*categoryEntity.Category, error)
}
