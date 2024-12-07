package entity

import (
	"time"

	"github.com/jambo0624/blog/internal/category/interfaces/http/dto"
	"github.com/jambo0624/blog/internal/shared/domain/validate"
)

type Category struct {
	ID        uint       `gorm:"primary_key" json:"id" binding:"required"`
	Name      string     `gorm:"size:100;not null" json:"name" binding:"required"`
	Slug      string     `gorm:"size:100;not null;unique" json:"slug" binding:"required"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}

// NewCategory create new category, all fields are required
func NewCategory(name, slug string) (*Category, error) {
	if name == "" {
		return nil, validate.ErrNameRequired
	}
	if slug == "" {
		return nil, validate.ErrSlugRequired
	}

	return &Category{
		Name:      name,
		Slug:      slug,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

// Update update category, only update provided fields
func (c *Category) Update(req *dto.UpdateCategoryRequest) {
	if req.Name != "" {
		c.Name = req.Name
	}
	if req.Slug != "" {
		c.Slug = req.Slug
	}
	c.UpdatedAt = time.Now()
}

// GetID get category id, implement Entity interface
func (c Category) GetID() uint {
	return c.ID
}
