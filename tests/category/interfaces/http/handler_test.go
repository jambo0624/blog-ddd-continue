package http_test

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    categoryHttp "github.com/jambo0624/blog/internal/category/interfaces/http"
    categoryService "github.com/jambo0624/blog/internal/category/application/service"
		categoryMock "github.com/jambo0624/blog/tests/testutil/mock/category"
)	


func TestCategoryHandler_CreateCategory(t *testing.T) {
	mockRepo := new(categoryMock.MockCategoryRepository)
	mockService := categoryService.NewCategoryService(mockRepo)
	handler := categoryHttp.NewCategoryHandler(mockService)

	body := `{"name": "Test Category"}`
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	jsonData, _ := json.Marshal(body)
	c.Request = httptest.NewRequest("POST", "/", bytes.NewBuffer(jsonData))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.CreateCategory(c)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestCategoryHandler_GetCategory(t *testing.T) {
    mockCategoryRepo := new(categoryMock.MockCategoryRepository)
		categoryService := categoryService.NewCategoryService(mockCategoryRepo)
		handler := categoryHttp.NewCategoryHandler(categoryService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "1"}}
	handler.GetCategory(c)

	assert.Equal(t, http.StatusOK, w.Code)
} 

func TestCategoryHandler_UpdateCategory(t *testing.T) {
	mockRepo := new(categoryMock.MockCategoryRepository)
	categoryService := categoryService.NewCategoryService(mockRepo)
	handler := categoryHttp.NewCategoryHandler(categoryService)

	body := `{"name":"Updated Category"}`
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("PUT", "/", bytes.NewBuffer([]byte(body)))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = []gin.Param{{Key: "id", Value: "1"}}

	handler.UpdateCategory(c)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCategoryHandler_DeleteCategory(t *testing.T) {
	mockRepo := new(categoryMock.MockCategoryRepository)
	categoryService := categoryService.NewCategoryService(mockRepo)
	handler := categoryHttp.NewCategoryHandler(categoryService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "1"}}

	handler.DeleteCategory(c)

	assert.Equal(t, http.StatusNoContent, w.Code)
} 	