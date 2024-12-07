package dto

type UpdateCategoryRequest struct {
	Name string `json:"name" binding:"omitempty,max=100"`
	Slug string `json:"slug" binding:"omitempty,max=100"`
}

func (r UpdateCategoryRequest) Validate() error {
	return nil
}
