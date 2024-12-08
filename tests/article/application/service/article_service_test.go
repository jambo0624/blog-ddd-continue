package service_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/jambo0624/blog/internal/article/application/service"
	"github.com/jambo0624/blog/internal/article/domain/query"
	articleEntity "github.com/jambo0624/blog/internal/article/domain/entity"
	categoryEntity "github.com/jambo0624/blog/internal/category/domain/entity"
	tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
	"github.com/jambo0624/blog/internal/article/interfaces/http/dto"
	mockArticle "github.com/jambo0624/blog/tests/testutil/mock/article"
	mockCategory "github.com/jambo0624/blog/tests/testutil/mock/category"
	mockTag "github.com/jambo0624/blog/tests/testutil/mock/tag"
	factory "github.com/jambo0624/blog/tests/testutil/factory"
)

func TestArticleService_Create(t *testing.T) {
	mockArticleRepo := new(mockArticle.MockArticleRepository)
	mockCategoryRepo := new(mockCategory.MockCategoryRepository)
	mockTagRepo := new(mockTag.MockTagRepository)

	articleService := service.NewArticleService(mockArticleRepo, mockCategoryRepo, mockTagRepo)

	category := &categoryEntity.Category{ID: 1, Name: "Test Category"}
	tag := &tagEntity.Tag{ID: 1, Name: "Test Tag"}

	// Setup expectations
	mockCategoryRepo.On("FindByID", uint(1), mock.Anything).Return(category, nil)
	mockTagRepo.On("FindByID", uint(1), mock.Anything).Return(tag, nil)
	mockArticleRepo.On("Save", mock.AnythingOfType("*entity.Article")).Return(nil)

	categoryFactory := factory.NewCategoryFactory()
	tagFactory := factory.NewTagFactory()
	articleFactory := factory.NewArticleFactory(categoryFactory, tagFactory)
	req := articleFactory.BuildCreateRequest(
		func(req *dto.CreateArticleRequest) {
			req.CategoryID = category.ID
			req.TagIDs = []uint{tag.ID}
		},
	)

	article, err := articleService.Create(req)

	assert.NoError(t, err)
	assert.NotNil(t, article)
	assert.Equal(t, req.Title, article.Title)
	assert.Equal(t, req.Content, article.Content)
	assert.Equal(t, category.ID, article.CategoryID)
	assert.Len(t, article.Tags, 1)
	assert.Equal(t, tag.ID, article.Tags[0].ID)
}

func TestArticleService_FindAll(t *testing.T) {
	mockArticleRepo := new(mockArticle.MockArticleRepository)
	mockCategoryRepo := new(mockCategory.MockCategoryRepository)
	mockTagRepo := new(mockTag.MockTagRepository)

	articleService := service.NewArticleService(mockArticleRepo, mockCategoryRepo, mockTagRepo)

	categoryFactory := factory.NewCategoryFactory()
	tagFactory := factory.NewTagFactory()
	articleFactory := factory.NewArticleFactory(categoryFactory, tagFactory)
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
	mockArticleRepo := new(mockArticle.MockArticleRepository)
	mockCategoryRepo := new(mockCategory.MockCategoryRepository)
	mockTagRepo := new(mockTag.MockTagRepository)

	articleService := service.NewArticleService(mockArticleRepo, mockCategoryRepo, mockTagRepo)

	categoryFactory := factory.NewCategoryFactory()
	tagFactory := factory.NewTagFactory()
	articleFactory := factory.NewArticleFactory(categoryFactory, tagFactory)
	article := articleFactory.BuildEntity(
		func(a *articleEntity.Article) {
			a.ID = 1
		},
	)
	category := categoryFactory.BuildEntity(
		func(c *categoryEntity.Category) {
			c.ID = 2
		},
	)
	tag := tagFactory.BuildEntity(
		func(t *tagEntity.Tag) {
			t.ID = 2
		},
	)

	mockArticleRepo.On("FindByID", uint(1), mock.Anything).Return(article, nil)
	mockCategoryRepo.On("FindByID", uint(2), mock.Anything).Return(category, nil)
	mockTagRepo.On("FindByID", uint(2), mock.Anything).Return(tag, nil)
	mockArticleRepo.On("Update", mock.AnythingOfType("*entity.Article")).Return(nil)

	req := articleFactory.BuildUpdateRequest(
		func(req *dto.UpdateArticleRequest) {
			req.CategoryID = category.ID
			req.TagIDs = []uint{tag.ID}
		},
	)

	updated, err := articleService.Update(1, req)

	assert.NoError(t, err)
	assert.Equal(t, req.Title, updated.Title)
	assert.Equal(t, req.Content, updated.Content)
	assert.Equal(t, category.ID, updated.CategoryID)
	assert.Len(t, updated.Tags, 1)
	assert.Equal(t, tag.ID, updated.Tags[0].ID)
}

func TestArticleService_Delete(t *testing.T) {
	mockArticleRepo := new(mockArticle.MockArticleRepository)
	mockCategoryRepo := new(mockCategory.MockCategoryRepository)
	mockTagRepo := new(mockTag.MockTagRepository)

	articleService := service.NewArticleService(mockArticleRepo, mockCategoryRepo, mockTagRepo)

	mockArticleRepo.On("Delete", uint(1)).Return(nil)

	err := articleService.Delete(1)

	assert.NoError(t, err)
	mockArticleRepo.AssertExpectations(t)
}