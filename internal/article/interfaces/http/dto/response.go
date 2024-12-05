package dto

type ArticleResponse struct {
    ID         uint   `json:"id"`
    CategoryID uint   `json:"category_id"`
    Title      string `json:"title"`
    Content    string `json:"content"`
    TagIDs     []uint `json:"tag_ids"`
    CreatedAt  string `json:"created_at"`
    UpdatedAt  string `json:"updated_at"`
} 