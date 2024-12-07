package dto

type CreateArticleRequest struct {
	Title      string `json:"title" binding:"required,max=255"`
	Content    string `json:"content" binding:"required"`
	CategoryID uint   `json:"category_id" binding:"required"`
	TagIDs     []uint `json:"tag_ids" binding:"omitempty"`
}

type UpdateArticleRequest struct {
	Title      string `json:"title" binding:"omitempty,max=255"`
	Content    string `json:"content" binding:"omitempty"`
	CategoryID uint   `json:"category_id" binding:"omitempty"`
	TagIDs     []uint `json:"tag_ids" binding:"omitempty"`
}

func (r CreateArticleRequest) Validate() error {
	// Business rules validation
	return nil
}

func (r UpdateArticleRequest) Validate() error {
	// Business rules validation
	return nil
}
