package dto

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