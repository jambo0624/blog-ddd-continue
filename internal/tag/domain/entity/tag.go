package entity

import (
	"time"

	"github.com/jambo0624/blog/internal/shared/domain/validate"
	"github.com/jambo0624/blog/internal/tag/interfaces/http/dto"
)

type Tag struct {
	ID        uint       `gorm:"primary_key" json:"id" binding:"required"`
	Name      string     `gorm:"size:100;not null;unique" json:"name" binding:"required, max=100"`
	Color     string     `gorm:"size:50" json:"color" binding:"required, hexcolor"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}

// NewTag create new tag, all fields are required
func NewTag(name, color string) (*Tag, error) {
	if name == "" {
		return nil, validate.ErrNameRequired
	}
	if color == "" {
		return nil, validate.ErrColorRequired
	}

	return &Tag{
		Name:      name,
		Color:     color,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

// GetID get tag id, implement Entity interface
func (t Tag) GetID() uint {
	return t.ID
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
