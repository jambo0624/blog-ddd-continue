package dto

import (
	"github.com/gin-gonic/gin"
	
	"github.com/jambo0624/blog/internal/shared/domain/constants"
	"github.com/jambo0624/blog/internal/shared/domain/validate"
	"github.com/jambo0624/blog/internal/shared/interfaces/http/dto"
)

type CreateTagRequest struct {
	dto.BaseRequest
	Name  string `json:"name" binding:"required"`
	Color string `json:"color" binding:"required"`
}

func (r CreateTagRequest) Bind(c *gin.Context) error {
	return r.BaseRequest.Bind(c)
}

func (r CreateTagRequest) Validate() error {
	if len(r.Name) > constants.MaxNameLength {
		return validate.ErrNameTooLong
	}
	return nil
} 