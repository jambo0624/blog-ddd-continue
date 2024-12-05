package entity_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/jambo0624/blog/internal/tag/domain/entity"
)

func TestNewTag(t *testing.T) {
    name := "Test Tag"
    color := "#FF0000"

    tag := entity.NewTag(name, color)

    assert.Equal(t, name, tag.Name)
    assert.Equal(t, color, tag.Color)
    assert.NotZero(t, tag.CreatedAt)
} 