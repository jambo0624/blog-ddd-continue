package http_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/mock"

	categoryService "github.com/jambo0624/blog/internal/category/application/service"
	"github.com/jambo0624/blog/internal/category/domain/entity"
	categoryHandler "github.com/jambo0624/blog/internal/category/interfaces/http"
	"github.com/jambo0624/blog/tests/testutil"
	"github.com/jambo0624/blog/tests/testutil/factory"
	mockCategory "github.com/jambo0624/blog/tests/testutil/mock/category"
)

func setupTest(t *testing.T) (*testutil.HTTPTester, *mockCategory.MockCategoryRepository) {
	t.Helper()

	mockRepo := new(mockCategory.MockCategoryRepository)
	service := categoryService.NewCategoryService(mockRepo)
	handler := categoryHandler.NewCategoryHandler(service)
	router := categoryHandler.NewCategoryRouter(handler)

	tester := testutil.NewHTTPTester(t, router.Register)

	return tester, mockRepo
}

func TestCategoryHandler_Create(t *testing.T) {
	tester, mockRepo := setupTest(t)
	factory := factory.NewCategoryFactory()

	req := factory.BuildCreateRequest()
	expectedCategory := factory.BuildEntity(
		factory.WithName(req.Name),
		factory.WithSlug(req.Slug),
	)

	mockRepo.On("Save", mock.MatchedBy(func(c *entity.Category) bool {
		return c.Name == expectedCategory.Name && c.Slug == expectedCategory.Slug
	})).Return(nil)

	tester.
		WithJSONBody(req).
		Post("/api/categories").
		SeeStatus(http.StatusCreated)
}

func TestCategoryHandler_GetByID(t *testing.T) {
	tester, mockRepo := setupTest(t)
	factory := factory.NewCategoryFactory()
	category := factory.BuildEntity()

	mockRepo.On("FindByID", category.ID, mock.Anything).Return(category, nil)

	tester.
		Get("/api/categories/3", nil).
		SeeStatus(http.StatusOK)
}

func TestCategoryHandler_List(t *testing.T) {
	tester, mockRepo := setupTest(t)
	factory := factory.NewCategoryFactory()
	categories := factory.BuildList(2)

	mockRepo.On("FindAll", mock.AnythingOfType("*query.CategoryQuery")).
		Return(categories, int64(len(categories)), nil)

	tester.
		Get("/api/categories", nil).
		SeeStatus(http.StatusOK)
}

func TestCategoryHandler_Update(t *testing.T) {
	tester, mockRepo := setupTest(t)
	factory := factory.NewCategoryFactory()

	existingCategory := factory.BuildEntity()
	req := factory.BuildUpdateRequest()

	mockRepo.On("FindByID", existingCategory.ID, mock.Anything).Return(existingCategory, nil)
	mockRepo.On("Update", mock.MatchedBy(func(c *entity.Category) bool {
		return c.ID == existingCategory.ID && c.Name == req.Name && c.Slug == req.Slug
	})).Return(nil)

	tester.
		WithJSONBody(req).
		Put("/api/categories/3").
		SeeStatus(http.StatusOK)
}

func TestCategoryHandler_Delete(t *testing.T) {
	tester, mockRepo := setupTest(t)

	mockRepo.On("Delete", uint(1)).Return(nil)

	tester.
		Delete("/api/categories/1").
		SeeStatus(http.StatusNoContent)
}
