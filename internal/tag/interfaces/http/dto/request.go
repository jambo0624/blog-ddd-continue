package dto

type CreateTagRequest struct {
    Name  string `json:"name" binding:"required"`
    Color string `json:"color" binding:"required"`
}

type UpdateTagRequest struct {
    Name  string `json:"name"`
    Color string `json:"color"`
} 