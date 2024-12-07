package dto

import (
	"github.com/gin-gonic/gin"

	"github.com/jambo0624/blog/internal/shared/domain/constants"
	"github.com/jambo0624/blog/internal/shared/domain/validate"
	"github.com/jambo0624/blog/internal/shared/interfaces/http/dto"
)

type UpdateCategoryRequest struct {
	dto.BaseRequest
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (r UpdateCategoryRequest) Bind(c *gin.Context) error {
	return r.BaseRequest.Bind(c)
}

func (r UpdateCategoryRequest) Validate() error {
	if r.Name != "" && len(r.Name) > constants.MaxNameLength {
		return validate.ErrNameTooLong
	}
	if r.Slug != "" && len(r.Slug) > constants.MaxSlugLength {
		return validate.ErrSlugTooLong
	}
	return nil
}