package dto

type CreateCategoryRequest struct {
	Name string `binding:"required,max=100" json:"name"`
	Slug string `binding:"required,max=100" json:"slug"`
}

type UpdateCategoryRequest struct {
	Name string `binding:"omitempty,max=100" json:"name"`
	Slug string `binding:"omitempty,max=100" json:"slug"`
}

func (r CreateCategoryRequest) Validate() error {
	// Business rules validation
	return nil
}

func (r UpdateCategoryRequest) Validate() error {
	// Business rules validation
	return nil
}
