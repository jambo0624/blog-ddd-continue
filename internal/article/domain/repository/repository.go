package repository

import (
	articleEntity "github.com/jambo0624/blog/internal/article/domain/entity"
	categoryEntity "github.com/jambo0624/blog/internal/category/domain/entity"
	tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
)

type ArticleRepository interface {
	Save(article *articleEntity.Article) error
	FindByID(id uint) (*articleEntity.Article, error)
	Update(article *articleEntity.Article) error
	Delete(id uint) error
	FindAll() ([]*articleEntity.Article, error)
}

type CategoryRepository interface {
	Save(category *categoryEntity.Category) error
	FindByID(id uint) (*categoryEntity.Category, error)
	FindBySlug(slug string) (*categoryEntity.Category, error)
	Update(category *categoryEntity.Category) error
	Delete(id uint) error
}

type TagRepository interface {
	Save(tag *tagEntity.Tag) error
	FindByID(id uint) (*tagEntity.Tag, error)
	FindByName(name string) (*tagEntity.Tag, error)
	Update(tag *tagEntity.Tag) error
	Delete(id uint) error
}
