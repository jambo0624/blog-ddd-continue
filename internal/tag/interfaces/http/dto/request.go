package dto

import (
	"github.com/jambo0624/blog/internal/shared/domain/constants"
	"github.com/jambo0624/blog/internal/shared/domain/query"
)

type CreateTagRequest struct {
	Name  string `json:"name" binding:"required"`
	Color string `json:"color" binding:"required"`
}

type UpdateTagRequest struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

func (r CreateTagRequest) Validate() error {
	if r.Name == "" {
		return query.ErrNameRequired
	}
	if len(r.Name) > constants.MaxNameLength {
		return query.ErrNameTooLong
	}

	if r.Color == "" {
		return query.ErrColorRequired
	}
	
	return nil
}

func (r UpdateTagRequest) Validate() error {
	if r.Name != "" && len(r.Name) > constants.MaxNameLength {
		return query.ErrNameTooLong
	}
	return nil
} 