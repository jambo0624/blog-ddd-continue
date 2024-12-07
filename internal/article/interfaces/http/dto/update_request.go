package dto

type UpdateArticleRequest struct {
	Title      string `json:"title" binding:"omitempty,max=255"`
	Content    string `json:"content" binding:"omitempty"`
	CategoryID uint   `json:"category_id" binding:"omitempty"`
	TagIDs     []uint `json:"tag_ids" binding:"omitempty"`
}

func (r UpdateArticleRequest) Validate() error {
	return nil
}
