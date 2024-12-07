package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response standard response structure
type Response struct {
	Code    int         `json:"code"`    // Business status code
	Message string      `json:"message"` // Response message
	Data    interface{} `json:"data"`    // Response data
	Meta    Meta        `json:"meta,omitempty"` // Metadata (pagination, etc.)
}

// Predefined status codes
const (
	CodeSuccess          = 0
	CodeInvalidParams    = 400001
	CodeUnauthorized     = 401001
	CodeForbidden        = 403001
	CodeNotFound         = 404001
	CodeInternalError    = 500001
	CodeValidationFailed = 422001
)

// Predefined messages
var messages = map[int]string{
	CodeSuccess:          "success",
	CodeInvalidParams:    "invalid parameters",
	CodeUnauthorized:     "unauthorized",
	CodeForbidden:        "forbidden",
	CodeNotFound:         "resource not found",
	CodeInternalError:    "internal server error",
	CodeValidationFailed: "validation failed",
}

// Success successful response
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    CodeSuccess,
		Message: messages[CodeSuccess],
		Data:    data,
	})
}

// SuccessWithMeta successful response with metadata
func SuccessWithMeta(c *gin.Context, data interface{}, meta Meta) {
	c.JSON(http.StatusOK, Response{
		Code:    CodeSuccess,
		Message: messages[CodeSuccess],
		Data:    data,
		Meta:    meta,
	})
}

// Created successful response
func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, Response{
		Code:    CodeSuccess,
		Message: messages[CodeSuccess],
		Data:    data,
	})
}

// NoContent no content response
func NoContent(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
}

// Error error response
func Error(c *gin.Context, httpCode, code int, msg string) {
	message := msg
	if message == "" {
		message = messages[code]
	}
	c.JSON(httpCode, Response{
		Code:    code,
		Message: message,
	})
}

// ValidationError validation error response
func ValidationError(c *gin.Context, err error) {
	Error(c, http.StatusUnprocessableEntity, CodeValidationFailed, err.Error())
}

// BadRequest bad request response
func BadRequest(c *gin.Context, err error) {
	Error(c, http.StatusBadRequest, CodeInvalidParams, err.Error())
}

// NotFound resource not found response
func NotFound(c *gin.Context) {
	Error(c, http.StatusNotFound, CodeNotFound, "")
}

// InternalError internal server error response
func InternalError(c *gin.Context, err error) {
	Error(c, http.StatusInternalServerError, CodeInternalError, err.Error())
} 
