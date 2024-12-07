package dto

// RequestDTO interface for all DTOs
type RequestDTO interface {
	Validate() error // Handle business rules validation
}
