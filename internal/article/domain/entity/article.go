package entity

import (
	"time"

	categoryEntity "github.com/jambo0624/blog/internal/category/domain/entity"
	tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
)

type Article struct {
	ID         uint
	CategoryID uint
	Category   categoryEntity.Category
	Title      string
	Content    string
	Tags       []tagEntity.Tag
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

func NewArticle(categoryID uint, title, content string) *Article {
	return &Article{
		CategoryID: categoryID,
		Title:      title,
		Content:    content,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

func (a *Article) AddTag(tag tagEntity.Tag) {
	a.Tags = append(a.Tags, tag)
}

func (a *Article) UpdateContent(title, content string) {
	a.Title = title
	a.Content = content
	a.UpdatedAt = time.Now()
}
