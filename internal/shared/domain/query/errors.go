package query

import "errors"

var (
	ErrInvalidIDFormat = errors.New("invalid id format")
	ErrInvalidLimit = errors.New("invalid limit")
	ErrInvalidOffset = errors.New("invalid offset")
	ErrNameRequired = errors.New("name is required")
	ErrNameTooLong = errors.New("name too long")
	ErrTitleRequired = errors.New("title is required")
	ErrTitleTooLong = errors.New("title too long")
	ErrContentTooLong = errors.New("content too long")
	ErrSlugTooLong = errors.New("slug too long")
	ErrInvalidOrderByField = errors.New("invalid order by field")
) 