package entity

import (
	"fmt"
	"time"

	"github.com/jambo0624/blog/internal/tag/interfaces/http/dto"
)

type Tag struct {
	ID        uint       `gorm:"primary_key"`
	Name      string     `gorm:"size:100;not null;unique"`
	Color     string     `gorm:"size:50"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `gorm:"index"`
}

// NewTag create new tag, all fields are required
func NewTag(name, color string) (*Tag, error) {
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}
	if color == "" {
		return nil, fmt.Errorf("color is required")
	}

	return &Tag{
		Name:      name,
		Color:     color,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

// Update update tag, only update provided fields
func (t *Tag) Update(req *dto.UpdateTagRequest) {
	if req.Name != "" {
		t.Name = req.Name
	}
	if req.Color != "" {
		t.Color = req.Color
	}
	t.UpdatedAt = time.Now()
}
