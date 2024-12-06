package dto

import "github.com/jambo0624/blog/internal/shared/domain/query"

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
	Slug string `json:"slug" binding:"required"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (r CreateCategoryRequest) Validate() error {
	if r.Name == "" {
		return query.ErrNameRequired
	}
	return nil
}

func (r UpdateCategoryRequest) Validate() error {
	if r.Name != "" && len(r.Name) > 100 {
		return query.ErrNameTooLong
	}
	return nil
}