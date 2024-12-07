package dto

type CreateTagRequest struct {
	Name  string `json:"name" binding:"required,max=50"`
	Color string `json:"color" binding:"required,hexcolor"`
}

type UpdateTagRequest struct {
	Name  string `json:"name" binding:"omitempty,max=100"`
	Color string `json:"color" binding:"omitempty,hexcolor"`
}

func (r CreateTagRequest) Validate() error {
	// Business rules validation
	return nil
}

func (r UpdateTagRequest) Validate() error {
	// Business rules validation
	return nil
}
