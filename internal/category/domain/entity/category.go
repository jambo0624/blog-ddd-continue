package entity

import (
    "fmt"
    "time"
    "github.com/jambo0624/blog/internal/category/interfaces/http/dto"
)

type Category struct {
    ID        uint       `gorm:"primary_key"`
    Name      string     `gorm:"size:100;not null"`
    Slug      string     `gorm:"size:100;not null;unique"`
    CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
    UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
    DeletedAt *time.Time `gorm:"index"`
}

// NewCategory create new category, all fields are required
func NewCategory(name, slug string) (*Category, error) {
    if name == "" {
        return nil, fmt.Errorf("name is required")
    }
    if slug == "" {
        return nil, fmt.Errorf("slug is required")
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