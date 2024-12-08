package http_test

import (
	"testing"
	"net/http"
	"github.com/stretchr/testify/mock"

	articleHandler "github.com/jambo0624/blog/internal/article/interfaces/http"
	articleService "github.com/jambo0624/blog/internal/article/application/service"
	mockArticle "github.com/jambo0624/blog/tests/testutil/mock/article"
	mockCategory "github.com/jambo0624/blog/tests/testutil/mock/category"
	mockTag "github.com/jambo0624/blog/tests/testutil/mock/tag"
	"github.com/jambo0624/blog/internal/article/interfaces/http/dto"
	"github.com/jambo0624/blog/tests/testutil/factory"
	"github.com/jambo0624/blog/tests/testutil"
)

func setupTest(t *testing.T) (*testutil.HttpTester, *mockArticle.MockArticleRepository, *mockCategory.MockCategoryRepository, *mockTag.MockTagRepository) {
	mockArticleRepo := new(mockArticle.MockArticleRepository)
	mockCategoryRepo := new(mockCategory.MockCategoryRepository)
	mockTagRepo := new(mockTag.MockTagRepository)

	service := articleService.NewArticleService(mockArticleRepo, mockCategoryRepo, mockTagRepo)
	handler := articleHandler.NewArticleHandler(service)
	router := articleHandler.NewArticleRouter(handler)

	actor := testutil.NewHttpTester(t, router)

	return actor, mockArticleRepo, mockCategoryRepo, mockTagRepo
}

func TestArticleHandler_Create(t *testing.T) {
	actor, mockArticleRepo, mockCategoryRepo, mockTagRepo := setupTest(t)

	categoryFactory := factory.NewCategoryFactory()
	tagFactory := factory.NewTagFactory()
	articleFactory := factory.NewArticleFactory(categoryFactory, tagFactory)

	category := categoryFactory.BuildEntity()
	tag := tagFactory.BuildEntity()
	req := articleFactory.BuildCreateRequest(func(r *dto.CreateArticleRequest) {
		r.CategoryID = category.ID
		r.TagIDs = []uint{tag.ID}
	})

	mockCategoryRepo.On("FindByID", category.ID, mock.Anything).Return(category, nil)
	mockTagRepo.On("FindByID", tag.ID, mock.Anything).Return(tag, nil)
	mockArticleRepo.On("Save", mock.AnythingOfType("*entity.Article")).Return(nil)

	actor.
		WithJSONBody(req).
		Post("/api/articles").
		SeeStatus(http.StatusCreated)
}

func TestArticleHandler_GetByID(t *testing.T) {
	actor, mockArticleRepo, _, _ := setupTest(t)
	articleFactory := factory.NewArticleFactory(factory.NewCategoryFactory(), factory.NewTagFactory())
	article := articleFactory.BuildEntity()

	mockArticleRepo.On("FindByID", article.ID, mock.Anything).Return(article, nil)

	actor.
		Get("/api/articles/3", nil).
		SeeStatus(http.StatusOK)
}

func TestArticleHandler_List(t *testing.T) {
	actor, mockArticleRepo, _, _ := setupTest(t)
	articleFactory := factory.NewArticleFactory(factory.NewCategoryFactory(), factory.NewTagFactory())
	articles := articleFactory.BuildList(2)

	mockArticleRepo.On("FindAll", mock.AnythingOfType("*query.ArticleQuery")).
		Return(articles, int64(len(articles)), nil)

	actor.
		Get("/api/articles", nil).
		SeeStatus(http.StatusOK)
}

func TestArticleHandler_Update(t *testing.T) {
	actor, mockArticleRepo, mockCategoryRepo, mockTagRepo := setupTest(t)

	categoryFactory := factory.NewCategoryFactory()
	tagFactory := factory.NewTagFactory()
	articleFactory := factory.NewArticleFactory(categoryFactory, tagFactory)

	category := categoryFactory.BuildEntity()
	tag := tagFactory.BuildEntity()
	article := articleFactory.BuildEntity()

	req := articleFactory.BuildUpdateRequest(func(r *dto.UpdateArticleRequest) {
		r.CategoryID = category.ID
		r.TagIDs = []uint{tag.ID}
	})

	mockArticleRepo.On("FindByID", article.ID, mock.Anything).Return(article, nil)
	mockCategoryRepo.On("FindByID", category.ID, mock.Anything).Return(category, nil)
	mockTagRepo.On("FindByID", tag.ID, mock.Anything).Return(tag, nil)
	mockArticleRepo.On("Update", mock.AnythingOfType("*entity.Article")).Return(nil)

	actor.
		WithJSONBody(req).
		Put("/api/articles/3").
		SeeStatus(http.StatusOK)
}

func TestArticleHandler_Delete(t *testing.T) {
	actor, mockArticleRepo, _, _ := setupTest(t)

	mockArticleRepo.On("Delete", uint(1)).Return(nil)

	actor.
		Delete("/api/articles/1").
		SeeStatus(http.StatusNoContent)
}