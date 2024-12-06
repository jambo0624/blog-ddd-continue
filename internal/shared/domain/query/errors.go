package query

import "errors"

var (
    ErrInvalidLimit = errors.New("invalid limit")
    ErrInvalidOffset = errors.New("invalid offset")
    ErrNameTooLong = errors.New("name too long")
    ErrTitleTooLong = errors.New("title too long")
    ErrContentTooLong = errors.New("content too long")
    ErrSlugTooLong = errors.New("slug too long")
) 