package repository

import (
	articleEntity "github.com/jambo0624/blog/internal/article/domain/entity"
	"github.com/jambo0624/blog/internal/article/domain/query"
)

type ArticleRepository interface {
	Save(article *articleEntity.Article) error
	FindByID(id uint) (*articleEntity.Article, error)
	Update(article *articleEntity.Article) error
	Delete(id uint) error
	FindAll(q *query.ArticleQuery) ([]*articleEntity.Article, error)
}
