package entity

import (
	"time"

	"github.com/jambo0624/blog/internal/article/interfaces/http/dto"
	categoryEntity "github.com/jambo0624/blog/internal/category/domain/entity"
	"github.com/jambo0624/blog/internal/shared/domain/errors"
	tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
)

type Article struct {
	ID         uint                    `binding:"required"                        gorm:"primaryKey"         json:"id"`
	CategoryID uint                    `binding:"required"                        gorm:"not null"           json:"categoryId"`
	Category   categoryEntity.Category `gorm:"foreignKey:CategoryID"              json:"category"`
	Title      string                  `binding:"required"                        gorm:"size:255;not null"  json:"title"`
	Content    string                  `binding:"required"                        gorm:"type:text;not null" json:"content"`
	Tags       []tagEntity.Tag         `gorm:"many2many:article_tags"             json:"tags"`
	CreatedAt  time.Time               `gorm:"not null;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt  time.Time               `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeletedAt  *time.Time              `gorm:"index"                              json:"deletedAt"`
}

func NewArticle(category *categoryEntity.Category, title, content string, tags []tagEntity.Tag) (*Article, error) {
	if category == nil {
		return nil, errors.ErrCategoryRequired
	}
	if title == "" {
		return nil, errors.ErrTitleRequired
	}
	if content == "" {
		return nil, errors.ErrContentRequired
	}

	return &Article{
		CategoryID: category.ID,
		Category:   *category,
		Title:      title,
		Content:    content,
		Tags:       tags,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, nil
}

func (a *Article) AddTag(tag tagEntity.Tag) error {
	for _, existingTag := range a.Tags {
		if existingTag.ID == tag.ID {
			return errors.ErrTagAlreadyExists
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

// GetID get article id, implement Entity interface.
func (a Article) GetID() uint {
	return a.ID
}

func (a *Article) GetFieldValue(field string) string {
	switch field {
	case "Title":
		return a.Title
	case "Content":
		return a.Content
	default:
		return ""
	}
}
