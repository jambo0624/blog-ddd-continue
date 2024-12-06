package repository

import (
	"github.com/jambo0624/blog/internal/shared/domain/repository"
	articleEntity "github.com/jambo0624/blog/internal/article/domain/entity"
	articleQuery "github.com/jambo0624/blog/internal/article/domain/query"
)

type ArticleRepository interface {
	repository.BaseRepository[articleEntity.Article, *articleQuery.ArticleQuery]
}
