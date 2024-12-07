package dto

type UpdateTagRequest struct {
	Name  string `json:"name" binding:"omitempty,max=100"`
	Color string `json:"color" binding:"omitempty,hexcolor"`
}

func (r UpdateTagRequest) Validate() error {
	return nil
} 