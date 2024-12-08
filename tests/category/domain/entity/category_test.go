package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/jambo0624/blog/internal/category/domain/entity"
	"github.com/jambo0624/blog/internal/category/interfaces/http/dto"
	"github.com/jambo0624/blog/internal/shared/domain/validate"
)

func TestNewCategory(t *testing.T) {
	tests := []struct {
		name        string
		categoryName string
		slug        string
		wantErr     bool
		expectedErr error
	}{
		{
			name:        "valid category",
			categoryName: "Test Category",
			slug:        "test-category",
			wantErr:     false,
		},
		{
			name:        "empty name",
			categoryName: "",
			slug:        "test-category",
			wantErr:     true,
			expectedErr: validate.ErrNameRequired,
		},
		{
			name:        "empty slug",
			categoryName: "Test Category",
			slug:        "",
			wantErr:     true,
			expectedErr: validate.ErrSlugRequired,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			category, err := entity.NewCategory(tt.categoryName, tt.slug)
			if tt.wantErr {
				assert.ErrorIs(t, err, tt.expectedErr)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.categoryName, category.Name)
			assert.Equal(t, tt.slug, category.Slug)
		})
	}
}

func TestCategory_Update(t *testing.T) {
	category, _ := entity.NewCategory("Original", "original")
	req := &dto.UpdateCategoryRequest{
		Name: "Updated",
		Slug: "updated",
	}

	category.Update(req)

	assert.Equal(t, "Updated", category.Name)
	assert.Equal(t, "updated", category.Slug)
} 