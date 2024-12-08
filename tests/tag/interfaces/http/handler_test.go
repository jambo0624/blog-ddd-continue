package http_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/mock"

	tagService "github.com/jambo0624/blog/internal/tag/application/service"
	"github.com/jambo0624/blog/internal/tag/domain/entity"
	tagHandler "github.com/jambo0624/blog/internal/tag/interfaces/http"
	"github.com/jambo0624/blog/tests/testutil"
	"github.com/jambo0624/blog/tests/testutil/factory"
	mockTag "github.com/jambo0624/blog/tests/testutil/mock/tag"
)

func setupTest(t *testing.T) (*testutil.HTTPTester, *mockTag.MockTagRepository) {
	t.Helper()

	mockRepo := new(mockTag.MockTagRepository)
	service := tagService.NewTagService(mockRepo)
	handler := tagHandler.NewTagHandler(service)
	router := tagHandler.NewTagRouter(handler)

	tester := testutil.NewHTTPTester(t, router.Register)

	return tester, mockRepo
}

func TestTagHandler_Create(t *testing.T) {
	tester, mockRepo := setupTest(t)
	factory := factory.NewTagFactory()

	req := factory.BuildCreateRequest()
	expectedTag := factory.BuildEntity(
		factory.WithName(req.Name),
		factory.WithColor(req.Color),
	)

	mockRepo.On("Save", mock.MatchedBy(func(t *entity.Tag) bool {
		return t.Name == expectedTag.Name && t.Color == expectedTag.Color
	})).Return(nil)

	tester.
		WithJSONBody(req).
		Post("/api/tags").
		SeeStatus(http.StatusCreated)
}

func TestTagHandler_GetByID(t *testing.T) {
	tester, mockRepo := setupTest(t)
	factory := factory.NewTagFactory()
	tag := factory.BuildEntity()

	mockRepo.On("FindByID", tag.ID, mock.Anything).Return(tag, nil)

	tester.
		Get("/api/tags/3", nil).
		SeeStatus(http.StatusOK)
}

func TestTagHandler_List(t *testing.T) {
	tester, mockRepo := setupTest(t)
	factory := factory.NewTagFactory()
	tags := factory.BuildList(2)

	mockRepo.On("FindAll", mock.AnythingOfType("*query.TagQuery")).
		Return(tags, int64(len(tags)), nil)

	tester.
		Get("/api/tags", nil).
		SeeStatus(http.StatusOK)
}

func TestTagHandler_Update(t *testing.T) {
	tester, mockRepo := setupTest(t)
	factory := factory.NewTagFactory()

	existingTag := factory.BuildEntity()
	req := factory.BuildUpdateRequest()

	mockRepo.On("FindByID", existingTag.ID, mock.Anything).Return(existingTag, nil)
	mockRepo.On("Update", mock.MatchedBy(func(t *entity.Tag) bool {
		return t.ID == existingTag.ID && t.Name == req.Name && t.Color == req.Color
	})).Return(nil)

	tester.
		WithJSONBody(req).
		Put("/api/tags/3").
		SeeStatus(http.StatusOK)
}

func TestTagHandler_Delete(t *testing.T) {
	tester, mockRepo := setupTest(t)

	mockRepo.On("Delete", uint(1)).Return(nil)

	tester.
		Delete("/api/tags/1").
		SeeStatus(http.StatusNoContent)
}
