package entity

import (
	"time"

	categoryEntity "github.com/jambo0624/blog/internal/category/domain/entity"
	tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
	"github.com/jambo0624/blog/internal/shared/domain/validate"
	"github.com/jambo0624/blog/internal/article/interfaces/http/dto"
)

type Article struct {
	ID         uint `gorm:"primaryKey" json:"id" binding:"required"`
	CategoryID uint `gorm:"not null" json:"category_id" binding:"required"`
	Category   categoryEntity.Category `gorm:"foreignKey:CategoryID" json:"category"`
	Title      string `gorm:"size:255;not null" json:"title" binding:"required"`
	Content    string `gorm:"type:text;not null" json:"content" binding:"required"`
	Tags       []tagEntity.Tag `gorm:"many2many:article_tags" json:"tags"`
	CreatedAt  time.Time	`gorm:"not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time	`gorm:"not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt  *time.Time	`gorm:"index" json:"deleted_at"`
}

func NewArticle(category *categoryEntity.Category, title, content string, tags []tagEntity.Tag) (*Article, error) {
	if category == nil {
		return nil, validate.ErrCategoryRequired
	}
	if title == "" {
		return nil, validate.ErrTitleRequired
	}
	if content == "" {
		return nil, validate.ErrContentRequired
	}

	return &Article{
		CategoryID: category.ID,
		Category:   *category,
		Title:     title,
		Content:   content,
		Tags:      tags,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (a *Article) AddTag(tag tagEntity.Tag) error {
	for _, existingTag := range a.Tags {
		if existingTag.ID == tag.ID {
			return validate.ErrTagAlreadyExists
		}
	}
	a.Tags = append(a.Tags, tag)
	return nil
}

func (a *Article) Update(req *dto.UpdateArticleRequest, category *categoryEntity.Category, tags []tagEntity.Tag) {
	if category != nil {
		a.CategoryID = category.ID
		a.Category = *category
	}
	if req.Title != "" {
		a.Title = req.Title
	}
	if req.Content != "" {
		a.Content = req.Content
	}
	if tags != nil {
		a.Tags = tags
	}
	a.UpdatedAt = time.Now()
}

// GetID get article id, implement Entity interface
func (a Article) GetID() uint {
	return a.ID
}
