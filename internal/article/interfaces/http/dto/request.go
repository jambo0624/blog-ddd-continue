package dto

type CreateArticleRequest struct {
	Title      string `binding:"required,max=255" json:"title"`
	Content    string `binding:"required"         json:"content"`
	CategoryID uint   `binding:"required"         json:"categoryId"`
	TagIDs     []uint `binding:"omitempty"        json:"tagIds"`
}

type UpdateArticleRequest struct {
	Title      string `binding:"omitempty,max=255" json:"title"`
	Content    string `binding:"omitempty"         json:"content"`
	CategoryID uint   `binding:"omitempty"         json:"categoryId"`
	TagIDs     []uint `binding:"omitempty"         json:"tagIds"`
}

func (r CreateArticleRequest) Validate() error {
	// Business rules validation
	return nil
}

func (r UpdateArticleRequest) Validate() error {
	// Business rules validation
	return nil
}
