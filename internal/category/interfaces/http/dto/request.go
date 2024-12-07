package dto

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required,max=100"`
	Slug string `json:"slug" binding:"required,max=100"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name" binding:"omitempty,max=100"`
	Slug string `json:"slug" binding:"omitempty,max=100"`
}

func (r CreateCategoryRequest) Validate() error {
	// Business rules validation
	return nil
}

func (r UpdateCategoryRequest) Validate() error {
	// Business rules validation
	return nil
}
