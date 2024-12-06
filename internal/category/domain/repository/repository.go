package repository

import (
	"github.com/jambo0624/blog/internal/shared/domain/repository"
	categoryEntity "github.com/jambo0624/blog/internal/category/domain/entity"
	categoryQuery "github.com/jambo0624/blog/internal/category/domain/query"
)

type CategoryRepository interface {
	repository.BaseRepository[categoryEntity.Category, *categoryQuery.CategoryQuery]
}
