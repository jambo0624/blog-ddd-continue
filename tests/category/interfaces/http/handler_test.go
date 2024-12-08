package http_test

import (
	"testing"
	"net/http"
	"github.com/stretchr/testify/mock"

	categoryHandler "github.com/jambo0624/blog/internal/category/interfaces/http"
	categoryService "github.com/jambo0624/blog/internal/category/application/service"
	mockCategory "github.com/jambo0624/blog/tests/testutil/mock/category"
	"github.com/jambo0624/blog/internal/category/domain/entity"
	"github.com/jambo0624/blog/tests/testutil/factory"
	"github.com/jambo0624/blog/tests/testutil"
)

func setupTest(t *testing.T) (*testutil.HttpTester, *mockCategory.MockCategoryRepository) {
	mockRepo := new(mockCategory.MockCategoryRepository)
	service := categoryService.NewCategoryService(mockRepo)
	handler := categoryHandler.NewCategoryHandler(service)
	router := categoryHandler.NewCategoryRouter(handler)

	actor := testutil.NewHttpTester(t, router)

	return actor, mockRepo
}

func TestCategoryHandler_Create(t *testing.T) {
	actor, mockRepo := setupTest(t)
	factory := factory.NewCategoryFactory()

	req := factory.BuildCreateRequest()
	expectedCategory := factory.BuildEntity(
		factory.WithName(req.Name),
		factory.WithSlug(req.Slug),
	)

	mockRepo.On("Save", mock.MatchedBy(func(c *entity.Category) bool {
		return c.Name == expectedCategory.Name && c.Slug == expectedCategory.Slug
	})).Return(nil)

	actor.
		WithJSONBody(req).
		Post("/api/categories").
		SeeStatus(http.StatusCreated)
}

func TestCategoryHandler_GetByID(t *testing.T) {
	actor, mockRepo := setupTest(t)
	factory := factory.NewCategoryFactory()
	category := factory.BuildEntity()

	mockRepo.On("FindByID", category.ID, mock.Anything).Return(category, nil)

	actor.
		Get("/api/categories/3", nil).
		SeeStatus(http.StatusOK)
}

func TestCategoryHandler_List(t *testing.T) {
	actor, mockRepo := setupTest(t)
	factory := factory.NewCategoryFactory()
	categories := factory.BuildList(2)

	mockRepo.On("FindAll", mock.AnythingOfType("*query.CategoryQuery")).
		Return(categories, int64(len(categories)), nil)

	actor.
		Get("/api/categories", nil).
		SeeStatus(http.StatusOK)
}

func TestCategoryHandler_Update(t *testing.T) {
	actor, mockRepo := setupTest(t)
	factory := factory.NewCategoryFactory()

	existingCategory := factory.BuildEntity()
	req := factory.BuildUpdateRequest()

	mockRepo.On("FindByID", existingCategory.ID, mock.Anything).Return(existingCategory, nil)
	mockRepo.On("Update", mock.MatchedBy(func(c *entity.Category) bool {
		return c.ID == existingCategory.ID && c.Name == req.Name && c.Slug == req.Slug
	})).Return(nil)

	actor.
		WithJSONBody(req).
		Put("/api/categories/3").
		SeeStatus(http.StatusOK)
}

func TestCategoryHandler_Delete(t *testing.T) {
	actor, mockRepo := setupTest(t)

	mockRepo.On("Delete", uint(1)).Return(nil)

	actor.
		Delete("/api/categories/1").
		SeeStatus(http.StatusNoContent)
} 