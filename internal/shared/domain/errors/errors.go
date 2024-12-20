package errors

import "errors"

// Validation errors.
var (
	// ID.
	ErrInvalidIDFormat = errors.New("invalid id format")

	// Title.
	ErrTitleRequired = errors.New("title is required")
	ErrTitleTooLong  = errors.New("title too long")

	// Category.
	ErrCategoryRequired = errors.New("category is required")

	// Content.
	ErrContentRequired = errors.New("content is required")
	ErrContentTooLong  = errors.New("content too long")

	// Name.
	ErrNameRequired = errors.New("name is required")
	ErrNameTooLong  = errors.New("name too long")

	// Tag.
	ErrTagAlreadyExists = errors.New("tag already exists")

	// Slug.
	ErrSlugRequired = errors.New("slug is required")
	ErrSlugTooLong  = errors.New("slug too long")

	// Color.
	ErrColorRequired = errors.New("color is required")

	// Limit.
	ErrInvalidLimit  = errors.New("invalid limit")
	ErrInvalidOffset = errors.New("invalid offset")

	// OrderBy.
	ErrInvalidOrderByField = errors.New("invalid order by field")
)
