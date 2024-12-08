package repository

import (
	articleEntity "github.com/jambo0624/blog/internal/article/domain/entity"
	articleQuery "github.com/jambo0624/blog/internal/article/domain/query"
	"github.com/jambo0624/blog/internal/shared/domain/repository"
)

type ArticleRepository interface {
	repository.BaseRepository[articleEntity.Article, *articleQuery.ArticleQuery]
}
