package dto

type CreateTagRequest struct {
	Name  string `binding:"required,max=50"   json:"name"`
	Color string `binding:"required,hexcolor" json:"color"`
}

type UpdateTagRequest struct {
	Name  string `binding:"omitempty,max=100"  json:"name"`
	Color string `binding:"omitempty,hexcolor" json:"color"`
}

func (r CreateTagRequest) Validate() error {
	// Business rules validation
	return nil
}

func (r UpdateTagRequest) Validate() error {
	// Business rules validation
	return nil
}
