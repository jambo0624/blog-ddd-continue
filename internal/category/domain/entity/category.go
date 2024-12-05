package entity

import "time"

type Category struct {
    ID        uint       `gorm:"primary_key"`
    Name      string     `gorm:"size:100;not null"`
    Slug      string     `gorm:"size:100;not null;unique"`
    CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
    UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
    DeletedAt *time.Time `gorm:"index"`
}

func NewCategory(name, slug string) *Category {
    return &Category{
        Name:      name,
        Slug:      slug,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }
} 