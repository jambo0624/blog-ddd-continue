package http_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/mock"

	articleService "github.com/jambo0624/blog/internal/article/application/service"
	articleHandler "github.com/jambo0624/blog/internal/article/interfaces/http"
	"github.com/jambo0624/blog/tests/testutil"
	"github.com/jambo0624/blog/tests/testutil/factory"
	mockArticle "github.com/jambo0624/blog/tests/testutil/mock/article"
	mockCategory "github.com/jambo0624/blog/tests/testutil/mock/category"
	mockTag "github.com/jambo0624/blog/tests/testutil/mock/tag"
)

func setupTest(t *testing.T) (
	*testutil.HTTPTester,
	*mockArticle.MockArticleRepository,
	*mockCategory.MockCategoryRepository,
	*mockTag.MockTagRepository,
) {
	t.Helper()
	mockArticleRepo := new(mockArticle.MockArticleRepository)
	mockCategoryRepo := new(mockCategory.MockCategoryRepository)
	mockTagRepo := new(mockTag.MockTagRepository)

	service := articleService.NewArticleService(mockArticleRepo, mockCategoryRepo, mockTagRepo)
	handler := articleHandler.NewArticleHandler(service)
	router := articleHandler.NewArticleRouter(handler)

	tester := testutil.NewHTTPTester(t, router.Register)

	return tester, mockArticleRepo, mockCategoryRepo, mockTagRepo
}

func TestArticleHandler_Create(t *testing.T) {
	tester, mockArticleRepo, mockCategoryRepo, mockTagRepo := setupTest(t)

	categoryFactory := factory.NewCategoryFactory()
	tagFactory := factory.NewTagFactory()
	articleFactory := factory.NewArticleFactory(categoryFactory, tagFactory)

	req, category, tag := articleFactory.BuildCreateRequest()

	mockCategoryRepo.On("FindByID", mock.AnythingOfType("uint"), []string(nil)).Return(category, nil)
	mockTagRepo.On("FindByID", mock.AnythingOfType("uint"), []string(nil)).Return(tag, nil)
	mockArticleRepo.On("Save", mock.AnythingOfType("*entity.Article")).Return(nil)

	tester.
		WithJSONBody(req).
		Post("/api/articles").
		SeeStatus(http.StatusCreated)
}

func TestArticleHandler_GetByID(t *testing.T) {
	tester, mockArticleRepo, _, _ := setupTest(t)
	articleFactory := factory.NewArticleFactory(factory.NewCategoryFactory(), factory.NewTagFactory())
	article, _, _ := articleFactory.BuildEntity()

	mockArticleRepo.On("FindByID", article.ID, mock.Anything).Return(article, nil)

	tester.
		Get(fmt.Sprintf("/api/articles/%d", article.ID), nil).
		SeeStatus(http.StatusOK)
}

func TestArticleHandler_List(t *testing.T) {
	tester, mockArticleRepo, _, _ := setupTest(t)
	articleFactory := factory.NewArticleFactory(factory.NewCategoryFactory(), factory.NewTagFactory())
	articles := articleFactory.BuildList(2)

	mockArticleRepo.On("FindAll", mock.AnythingOfType("*query.ArticleQuery")).
		Return(articles, int64(len(articles)), nil)

	tester.
		Get("/api/articles", nil).
		SeeStatus(http.StatusOK)
}

func TestArticleHandler_Update(t *testing.T) {
	tester, mockArticleRepo, mockCategoryRepo, mockTagRepo := setupTest(t)

	categoryFactory := factory.NewCategoryFactory()
	tagFactory := factory.NewTagFactory()
	articleFactory := factory.NewArticleFactory(categoryFactory, tagFactory)

	article, _, _ := articleFactory.BuildEntity()

	req, category, tag := articleFactory.BuildUpdateRequest()

	mockArticleRepo.On("FindByID", mock.AnythingOfType("uint"), []string(nil)).Return(article, nil)
	mockCategoryRepo.On("FindByID", mock.AnythingOfType("uint"), []string(nil)).Return(category, nil)
	mockTagRepo.On("FindByID", mock.AnythingOfType("uint"), []string(nil)).Return(tag, nil)
	mockArticleRepo.On("Update", mock.AnythingOfType("*entity.Article")).Return(nil)

	tester.
		WithJSONBody(req).
		Put(fmt.Sprintf("/api/articles/%d", article.ID)).
		SeeStatus(http.StatusOK)
}

func TestArticleHandler_Delete(t *testing.T) {
	tester, mockArticleRepo, _, _ := setupTest(t)

	mockArticleRepo.On("Delete", uint(1)).Return(nil)

	tester.
		Delete("/api/articles/1").
		SeeStatus(http.StatusNoContent)
}
