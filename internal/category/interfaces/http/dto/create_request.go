package dto

import (
	"github.com/gin-gonic/gin"
	
	"github.com/jambo0624/blog/internal/shared/domain/constants"
	"github.com/jambo0624/blog/internal/shared/domain/validate"
	"github.com/jambo0624/blog/internal/shared/interfaces/http/dto"
)

type CreateCategoryRequest struct {
	dto.BaseRequest
	Name string `json:"name" binding:"required"`
	Slug string `json:"slug" binding:"required"`
}

func (r CreateCategoryRequest) Bind(c *gin.Context) error {
	return r.BaseRequest.Bind(c)
}
func (r CreateCategoryRequest) Validate() error {
	// Business rules validation
	if len(r.Name) > constants.MaxNameLength {
		return validate.ErrNameTooLong
	}

	if len(r.Slug) > constants.MaxSlugLength {
		return validate.ErrSlugTooLong
	}
	return nil
}