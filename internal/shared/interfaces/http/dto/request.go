package dto

import "github.com/gin-gonic/gin"

// RequestDTO interface for all DTOs
type RequestDTO interface {
	Bind(*gin.Context) error // Handle binding and basic validation
	Validate() error         // Handle business rules validation
}

// BaseRequest provides common binding functionality
type BaseRequest struct{}

func (r *BaseRequest) Bind(c *gin.Context) error {
	if err := c.ShouldBindJSON(r); err != nil {
		return err
	}
	return nil
}
