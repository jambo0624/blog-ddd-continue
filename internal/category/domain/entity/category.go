package entity

import (
	"time"

	"github.com/jambo0624/blog/internal/category/interfaces/http/dto"
	"github.com/jambo0624/blog/internal/shared/domain/errors"
)

type Category struct {
	ID        uint       `binding:"required"               gorm:"primary_key"              json:"id"`
	Name      string     `binding:"required"               gorm:"size:100;not null"        json:"name"`
	Slug      string     `binding:"required"               gorm:"size:100;not null;unique" json:"slug"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"index"                     json:"deletedAt"`
}

// NewCategory create new category, all fields are required.
func NewCategory(name, slug string) (*Category, error) {
	if name == "" {
		return nil, errors.ErrNameRequired
	}
	if slug == "" {
		return nil, errors.ErrSlugRequired
	}

	return &Category{
		Name:      name,
		Slug:      slug,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

// Update update category, only update provided fields.
func (c *Category) Update(req *dto.UpdateCategoryRequest) {
	if req.Name != "" {
		c.Name = req.Name
	}
	if req.Slug != "" {
		c.Slug = req.Slug
	}
	c.UpdatedAt = time.Now()
}

// GetID get category id, implement Entity interface.
func (c Category) GetID() uint {
	return c.ID
}
