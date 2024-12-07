package dto

import (
	"github.com/gin-gonic/gin"

	"github.com/jambo0624/blog/internal/shared/domain/constants"
	"github.com/jambo0624/blog/internal/shared/domain/validate"
	"github.com/jambo0624/blog/internal/shared/interfaces/http/dto"
)

type UpdateArticleRequest struct {
	dto.BaseRequest
	Title      string `json:"title"`
	Content    string `json:"content"`
	CategoryID uint   `json:"category_id"`
	TagIDs     []uint `json:"tag_ids"`
}

func (r UpdateArticleRequest) Bind(c *gin.Context) error {
	return r.BaseRequest.Bind(c)
}

func (r UpdateArticleRequest) Validate() error {
	// Business rules validation
	if len(r.Title) > constants.MaxTitleLength {
		return validate.ErrTitleTooLong
	}
	if len(r.Content) > constants.MaxContentLength {
		return validate.ErrContentTooLong
	}
	return nil
}
