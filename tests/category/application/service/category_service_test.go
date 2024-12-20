package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"github.com/jambo0624/blog/internal/category/application/service"
	"github.com/jambo0624/blog/internal/category/domain/query"
	factory "github.com/jambo0624/blog/tests/testutil/factory"
	mockCategory "github.com/jambo0624/blog/tests/testutil/mock/category"
)

func setupTest(t *testing.T) (
	*mockCategory.MockCategoryRepository,
	*service.CategoryService,
	*factory.CategoryFactory,
) {
	t.Helper()

	mockRepo := new(mockCategory.MockCategoryRepository)
	categoryService := service.NewCategoryService(mockRepo)
	factory := factory.NewCategoryFactory()

	return mockRepo, categoryService, factory
}

func TestCategoryService_Create(t *testing.T) {
	mockRepo, categoryService, factory := setupTest(t)

	req := factory.BuildCreateRequest()

	mockRepo.On("Save", mock.AnythingOfType("*entity.Category")).Return(nil)

	category, err := categoryService.Create(req)

	require.NoError(t, err)
	require.NotNil(t, category)
	assert.Equal(t, req.Name, category.Name)
	assert.Equal(t, req.Slug, category.Slug)
}

func TestCategoryService_FindAll(t *testing.T) {
	mockRepo, categoryService, factory := setupTest(t)

	expectedCategories := factory.BuildList(2)

	q := query.NewCategoryQuery()
	mockRepo.On("FindAll", q).Return(expectedCategories, int64(2), nil)

	found, total, err := categoryService.FindAll(q)

	require.NoError(t, err)
	assert.Equal(t, int64(len(expectedCategories)), total)
	assert.Equal(t, expectedCategories, found)
}

func TestCategoryService_Update(t *testing.T) {
	mockRepo, categoryService, factory := setupTest(t)

	expectedCategory := factory.BuildEntity()

	mockRepo.On("FindByID", uint(1), mock.Anything).Return(expectedCategory, nil)
	mockRepo.On("Update", mock.AnythingOfType("*entity.Category")).Return(nil)

	req := factory.BuildUpdateRequest()

	updated, err := categoryService.Update(1, req)

	require.NoError(t, err)
	assert.Equal(t, expectedCategory.Name, updated.Name)
	assert.Equal(t, expectedCategory.Slug, updated.Slug)
}

func TestCategoryService_Delete(t *testing.T) {
	mockRepo, categoryService, _ := setupTest(t)

	mockRepo.On("Delete", uint(1)).Return(nil)

	err := categoryService.Delete(1)

	require.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCategoryService_Create_ValidationError(t *testing.T) {
	mockRepo, categoryService, factory := setupTest(t)

	req := factory.BuildCreateRequest()
	req.Name = "" // invalid name

	category, err := categoryService.Create(req)
	require.Error(t, err)
	assert.Nil(t, category)
	mockRepo.AssertNotCalled(t, "Save")
}

func TestCategoryService_Update_NotFound(t *testing.T) {
	mockRepo, categoryService, factory := setupTest(t)

	mockRepo.On("FindByID", uint(999), mock.Anything).Return(nil, gorm.ErrRecordNotFound)

	req := factory.BuildUpdateRequest()
	category, err := categoryService.Update(999, req)

	require.Error(t, err)
	assert.Nil(t, category)
	mockRepo.AssertNotCalled(t, "Update")
}
