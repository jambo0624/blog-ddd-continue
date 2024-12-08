package entity

import (
	"time"

	"github.com/jambo0624/blog/internal/shared/domain/errors"
	"github.com/jambo0624/blog/internal/tag/interfaces/http/dto"
)

type Tag struct {
	ID        uint       `binding:"required"               gorm:"primary_key"              json:"id"`
	Name      string     `binding:"required, max=100"      gorm:"size:100;not null;unique" json:"name"`
	Color     string     `binding:"required, hexcolor"     gorm:"size:50"                  json:"color"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"index"                     json:"deletedAt"`
}

// NewTag create new tag, all fields are required.
func NewTag(name, color string) (*Tag, error) {
	if name == "" {
		return nil, errors.ErrNameRequired
	}
	if color == "" {
		return nil, errors.ErrColorRequired
	}

	return &Tag{
		Name:      name,
		Color:     color,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

// GetID get tag id, implement Entity interface.
func (t Tag) GetID() uint {
	return t.ID
}

// Update update tag, only update provided fields.
func (t *Tag) Update(req *dto.UpdateTagRequest) {
	if req.Name != "" {
		t.Name = req.Name
	}
	if req.Color != "" {
		t.Color = req.Color
	}
	t.UpdatedAt = time.Now()
}
