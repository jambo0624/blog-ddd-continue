package dto

type CreateArticleRequest struct {
	Title      string `json:"title" binding:"required,max=255"`
	Content    string `json:"content" binding:"required"`
	CategoryID uint   `json:"category_id" binding:"required"`
	TagIDs     []uint `json:"tag_ids" binding:"omitempty"`
}

func (r CreateArticleRequest) Validate() error {
	return nil
}
