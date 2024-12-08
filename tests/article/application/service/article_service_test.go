package service_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"

	"github.com/jambo0624/blog/internal/article/application/service"
	"github.com/jambo0624/blog/internal/article/domain/query"
	mockArticle "github.com/jambo0624/blog/tests/testutil/mock/article"
	mockCategory "github.com/jambo0624/blog/tests/testutil/mock/category"
	mockTag "github.com/jambo0624/blog/tests/testutil/mock/tag"
	factory "github.com/jambo0624/blog/tests/testutil/factory"
)

func setupTest(t *testing.T) (
	*mockArticle.MockArticleRepository,
	*service.ArticleService,
	*factory.ArticleFactory,
	*mockCategory.MockCategoryRepository,
	*mockTag.MockTagRepository,
) {
	t.Helper()
	
	mockArticleRepo := new(mockArticle.MockArticleRepository)
	mockCategoryRepo := new(mockCategory.MockCategoryRepository)
	mockTagRepo := new(mockTag.MockTagRepository)
	articleService := service.NewArticleService(mockArticleRepo, mockCategoryRepo, mockTagRepo)
	articleFactory := factory.NewArticleFactory(factory.NewCategoryFactory(), factory.NewTagFactory())

	return mockArticleRepo, articleService, articleFactory, mockCategoryRepo, mockTagRepo
}

func TestArticleService_Create(t *testing.T) {
	mockArticleRepo, articleService, articleFactory, mockCategoryRepo, mockTagRepo := setupTest(t)

	req, category, tag := articleFactory.BuildCreateRequest()

	// Setup expectations
	mockCategoryRepo.On("FindByID", mock.AnythingOfType("uint"), []string(nil)).Return(category, nil)
	mockTagRepo.On("FindByID", mock.AnythingOfType("uint"), []string(nil)).Return(tag, nil)
	mockArticleRepo.On("Save", mock.AnythingOfType("*entity.Article")).Return(nil)

	article, err := articleService.Create(req)

	assert.NoError(t, err)
	assert.NotNil(t, article)
	assert.Equal(t, req.Title, article.Title)
	assert.Equal(t, req.Content, article.Content)
	assert.Equal(t, category.ID, article.CategoryID)
	assert.Len(t, article.Tags, 2)
	assert.Equal(t, tag.ID, article.Tags[0].ID)
}

func TestArticleService_FindAll(t *testing.T) {
	mockArticleRepo, articleService, articleFactory, _, _ := setupTest(t)

	articles := articleFactory.BuildList(2)

	q := query.NewArticleQuery()
	mockArticleRepo.On("FindAll", q).Return(articles, int64(2), nil)

	found, total, err := articleService.FindAll(q)

	assert.NoError(t, err)
	assert.Equal(t, int64(2), total)
	assert.Len(t, found, 2)
	assert.Equal(t, articles[0].Title, found[0].Title)
}

func TestArticleService_Update(t *testing.T) {
	mockArticleRepo, articleService, articleFactory, mockCategoryRepo, mockTagRepo := setupTest(t)

	article, _, _ := articleFactory.BuildEntity()
	req, category, tag := articleFactory.BuildUpdateRequest()

	mockArticleRepo.On("FindByID", mock.AnythingOfType("uint"), []string(nil)).Return(article, nil)
	mockCategoryRepo.On("FindByID", mock.AnythingOfType("uint"), []string(nil)).Return(category, nil)
	mockTagRepo.On("FindByID", mock.AnythingOfType("uint"), []string(nil)).Return(tag, nil)
	mockArticleRepo.On("Update", mock.AnythingOfType("*entity.Article")).Return(nil)

	updated, err := articleService.Update(article.ID, req)

	assert.NoError(t, err)
	assert.Equal(t, req.Title, updated.Title)
	assert.Equal(t, req.Content, updated.Content)
	assert.Equal(t, category.ID, updated.CategoryID)
	assert.Len(t, updated.Tags, 2)
	assert.Equal(t, tag.ID, updated.Tags[0].ID)
}

func TestArticleService_Delete(t *testing.T) {
	mockArticleRepo, articleService, _, _, _ := setupTest(t)

	mockArticleRepo.On("Delete", mock.AnythingOfType("uint")).Return(nil)

	err := articleService.Delete(1)

	assert.NoError(t, err)
	mockArticleRepo.AssertExpectations(t)
}

func TestArticleService_Create_ValidationError(t *testing.T) {
	mockArticleRepo, articleService, articleFactory, mockCategoryRepo, mockTagRepo := setupTest(t)

	req, category, tag := articleFactory.BuildCreateRequest()
	req.Title = "" // invalid title

	// Service still tries to find category and tag
	// So we need to set expectations
	mockCategoryRepo.On("FindByID", mock.AnythingOfType("uint"), []string(nil)).Return(category, nil)
	mockTagRepo.On("FindByID", mock.AnythingOfType("uint"), []string(nil)).Return(tag, nil)

	article, err := articleService.Create(req)

	assert.Error(t, err)
	assert.Nil(t, article)
	mockArticleRepo.AssertNotCalled(t, "Save")
}

func TestArticleService_Update_NotFound(t *testing.T) {
	mockArticleRepo, articleService, articleFactory, mockCategoryRepo, mockTagRepo := setupTest(t)

	mockArticleRepo.On("FindByID", uint(999), mock.Anything).Return(nil, gorm.ErrRecordNotFound)

	req, _, _ := articleFactory.BuildUpdateRequest()
	article, err := articleService.Update(999, req)

	assert.Error(t, err)
	assert.Nil(t, article)
	mockArticleRepo.AssertNotCalled(t, "Update")
	mockCategoryRepo.AssertNotCalled(t, "FindByID")
	mockTagRepo.AssertNotCalled(t, "FindByID")
}