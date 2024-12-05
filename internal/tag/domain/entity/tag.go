package entity

import (
	"time"
)

type Tag struct {
	ID        uint       `gorm:"primary_key"`
	Name      string     `gorm:"size:100;not null;unique"`
	Color     string     `gorm:"size:50"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `gorm:"index"`
}

func NewTag(name, color string) *Tag {
	return &Tag{
		Name:      name,
		Color:     color,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
