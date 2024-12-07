package dto

import (
	"github.com/gin-gonic/gin"
	
	"github.com/jambo0624/blog/internal/shared/domain/constants"
	"github.com/jambo0624/blog/internal/shared/domain/validate"
	"github.com/jambo0624/blog/internal/shared/interfaces/http/dto"
)

type UpdateTagRequest struct {
	dto.BaseRequest
	Name  string `json:"name"`
	Color string `json:"color"`
}

func (r UpdateTagRequest) Bind(c *gin.Context) error {
	return r.BaseRequest.Bind(c)
}

func (r UpdateTagRequest) Validate() error {
	if r.Name != "" && len(r.Name) > constants.MaxNameLength {
		return validate.ErrNameTooLong
	}
	return nil
} 