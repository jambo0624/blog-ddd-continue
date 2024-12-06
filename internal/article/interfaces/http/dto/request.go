package dto

import "github.com/jambo0624/blog/internal/shared/domain/query"

type CreateArticleRequest struct {
    CategoryID uint   `json:"category_id" binding:"required"`
    Title      string `json:"title" binding:"required"`
    Content    string `json:"content" binding:"required"`
    TagIDs     []uint `json:"tag_ids"`
}

type UpdateArticleRequest struct {
    CategoryID uint   `json:"category_id"`
    Title      string `json:"title"`
    Content    string `json:"content"`
    TagIDs     []uint `json:"tag_ids"`
} 


func (r CreateArticleRequest) Validate() error {
	if r.Title == "" {
		return query.ErrTitleRequired
	}
	return nil
}

func (r UpdateArticleRequest) Validate() error {
	if r.Title != "" && len(r.Title) > 100 {
		return query.ErrTitleTooLong
	}
	return nil
}