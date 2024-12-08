package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jambo0624/blog/internal/article/domain/entity"
	"github.com/jambo0624/blog/internal/article/interfaces/http/dto"
	categoryEntity "github.com/jambo0624/blog/internal/category/domain/entity"
	"github.com/jambo0624/blog/internal/shared/domain/errors"
	tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
)

func TestNewArticle(t *testing.T) {
	validCategory, _ := categoryEntity.NewCategory("Test Category", "test-category")
	validTag, _ := tagEntity.NewTag("Test Tag", "#FF0000")

	tests := []struct {
		name        string
		category    *categoryEntity.Category
		title       string
		content     string
		tags        []tagEntity.Tag
		wantErr     bool
		expectedErr error
	}{
		{
			name:     "valid article",
			category: validCategory,
			title:    "Test Title",
			content:  "Test Content",
			tags:     []tagEntity.Tag{*validTag},
			wantErr:  false,
		},
		{
			name:        "nil category",
			category:    nil,
			title:       "Test Title",
			content:     "Test Content",
			wantErr:     true,
			expectedErr: errors.ErrCategoryRequired,
		},
		{
			name:        "empty title",
			category:    validCategory,
			title:       "",
			content:     "Test Content",
			wantErr:     true,
			expectedErr: errors.ErrTitleRequired,
		},
		{
			name:        "empty content",
			category:    validCategory,
			title:       "Test Title",
			content:     "",
			wantErr:     true,
			expectedErr: errors.ErrContentRequired,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			article, err := entity.NewArticle(tt.category, tt.title, tt.content, tt.tags)
			if tt.wantErr {
				assert.ErrorIs(t, err, tt.expectedErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.title, article.Title)
			assert.Equal(t, tt.content, article.Content)
			assert.Equal(t, tt.category.ID, article.CategoryID)
			if len(tt.tags) > 0 {
				assert.Equal(t, tt.tags[0].ID, article.Tags[0].ID)
			}
		})
	}
}

func TestArticle_AddTag(t *testing.T) {
	category, _ := categoryEntity.NewCategory("Test Category", "test-category")
	article, _ := entity.NewArticle(category, "Test Title", "Test Content", nil)
	tag, _ := tagEntity.NewTag("Test Tag", "#FF0000")

	tests := []struct {
		name        string
		setupTags   func()
		tag         tagEntity.Tag
		wantErr     bool
		expectedErr error
	}{
		{
			name:    "add new tag",
			tag:     *tag,
			wantErr: false,
		},
		{
			name: "add duplicate tag",
			setupTags: func() {
				err := article.AddTag(*tag)
				require.Error(t, err)
			},
			tag:         *tag,
			wantErr:     true,
			expectedErr: errors.ErrTagAlreadyExists,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setupTags != nil {
				tt.setupTags()
			}

			err := article.AddTag(tt.tag)
			if tt.wantErr {
				assert.ErrorIs(t, err, tt.expectedErr)
				return
			}
			require.NoError(t, err)
			assert.Contains(t, article.Tags, tt.tag)
		})
	}
}

func TestArticle_Update(t *testing.T) {
	category, _ := categoryEntity.NewCategory("Original Category", "original-category")
	tag, _ := tagEntity.NewTag("Original Tag", "#000000")
	article, _ := entity.NewArticle(category, "Original Title", "Original Content", []tagEntity.Tag{*tag})

	newCategory, _ := categoryEntity.NewCategory("New Category", "new-category")
	newTag, _ := tagEntity.NewTag("New Tag", "#FFFFFF")

	req := &dto.UpdateArticleRequest{
		Title:      "Updated Title",
		Content:    "Updated Content",
		CategoryID: newCategory.ID,
		TagIDs:     []uint{newTag.ID},
	}

	article.Update(req, newCategory, []tagEntity.Tag{*newTag})

	assert.Equal(t, "Updated Title", article.Title)
	assert.Equal(t, "Updated Content", article.Content)
	assert.Equal(t, newCategory.ID, article.CategoryID)
	assert.Equal(t, newTag.ID, article.Tags[0].ID)
}
