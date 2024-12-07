package dto

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required,max=100"`
	Slug string `json:"slug" binding:"required,max=100"`
}

func (r CreateCategoryRequest) Validate() error {
	return nil
}
