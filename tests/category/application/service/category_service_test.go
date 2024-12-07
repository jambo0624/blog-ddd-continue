package service_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/jambo0624/blog/internal/category/application/service"
	"github.com/jambo0624/blog/internal/category/domain/entity"
	"github.com/jambo0624/blog/internal/category/domain/query"
	"github.com/jambo0624/blog/internal/category/interfaces/http/dto"
	mockCategory "github.com/jambo0624/blog/tests/testutil/mock/category"
)

func TestCategoryService_Create(t *testing.T) {
	mockRepo := new(mockCategory.MockCategoryRepository)
	categoryService := service.NewCategoryService(mockRepo)

	req := &dto.CreateCategoryRequest{
		Name: "Test Category",
		Slug: "test-category",
	}

	mockRepo.On("Save", mock.AnythingOfType("*entity.Category")).Return(nil)

	category, err := categoryService.Create(req)

	assert.NoError(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, req.Name, category.Name)
	assert.Equal(t, req.Slug, category.Slug)
}

func TestCategoryService_FindAll(t *testing.T) {
	mockRepo := new(mockCategory.MockCategoryRepository)
	categoryService := service.NewCategoryService(mockRepo)

	categories := []*entity.Category{
		{ID: 1, Name: "Category 1", Slug: "category-1"},
		{ID: 2, Name: "Category 2", Slug: "category-2"},
	}

	q := query.NewCategoryQuery()
	mockRepo.On("FindAll", q).Return(categories, int64(2), nil)

	found, total, err := categoryService.FindAll(q)

	assert.NoError(t, err)
	assert.Equal(t, int64(2), total)
	assert.Len(t, found, 2)
	assert.Equal(t, "Category 1", found[0].Name)
}

func TestCategoryService_Update(t *testing.T) {
	mockRepo := new(mockCategory.MockCategoryRepository)
	categoryService := service.NewCategoryService(mockRepo)

	category := &entity.Category{ID: 1, Name: "Old Name", Slug: "old-slug"}

	mockRepo.On("FindByID", uint(1), mock.Anything).Return(category, nil)
	mockRepo.On("Update", mock.AnythingOfType("*entity.Category")).Return(nil)

	req := &dto.UpdateCategoryRequest{
		Name: "New Name",
		Slug: "new-slug",
	}

	updated, err := categoryService.Update(1, req)

	assert.NoError(t, err)
	assert.Equal(t, "New Name", updated.Name)
	assert.Equal(t, "new-slug", updated.Slug)
}

func TestCategoryService_Delete(t *testing.T) {
	mockRepo := new(mockCategory.MockCategoryRepository)
	categoryService := service.NewCategoryService(mockRepo)

	mockRepo.On("Delete", uint(1)).Return(nil)

	err := categoryService.Delete(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}