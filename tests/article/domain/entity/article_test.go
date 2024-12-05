package entity_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/jambo0624/blog/internal/article/domain/entity"
)

func TestNewArticle(t *testing.T) {
	categoryID := uint(1)
	title := "Test Article"
	content := "Test Content"

	article := entity.NewArticle(categoryID, title, content)

	assert.Equal(t, categoryID, article.CategoryID)
	assert.Equal(t, title, article.Title)
	assert.Equal(t, content, article.Content)
	assert.NotZero(t, article.CreatedAt)
} 