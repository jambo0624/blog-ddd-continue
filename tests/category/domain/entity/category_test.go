package entity_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/jambo0624/blog/internal/category/domain/entity"
)

func TestNewCategory(t *testing.T) {
    name := "Test Category"
    slug := "test-category"

    category := entity.NewCategory(name, slug)

    assert.Equal(t, name, category.Name)
    assert.Equal(t, slug, category.Slug)
    assert.NotZero(t, category.CreatedAt)
} 