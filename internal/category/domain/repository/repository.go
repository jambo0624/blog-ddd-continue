package repository

import (
	categoryEntity "github.com/jambo0624/blog/internal/category/domain/entity"
	categoryQuery "github.com/jambo0624/blog/internal/category/domain/query"
	"github.com/jambo0624/blog/internal/shared/domain/repository"
)

type CategoryRepository interface {
	repository.BaseRepository[categoryEntity.Category, *categoryQuery.CategoryQuery]
}
