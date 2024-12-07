package dto

type CreateTagRequest struct {
	Name  string `json:"name" binding:"required,max=50"`
	Color string `json:"color" binding:"required,hexcolor"`
}

func (r CreateTagRequest) Validate() error {
	return nil
} 