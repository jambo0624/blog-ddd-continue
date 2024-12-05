package dto

type TagResponse struct {
    ID        uint   `json:"id"`
    Name      string `json:"name"`
    Color     string `json:"color"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
} 