package errors

import "errors"

// Repository errors.
var (
	ErrInvalidQueryType     = errors.New("query does not implement QueryFilter interface")
	ErrFailedToInitializeDB = errors.New("failed to initialize database")
	ErrDBNotInitialized     = errors.New("database connection not initialized")
)

// Config errors.
var (
	ErrFailedToReadConfig = errors.New("failed to read config")
)
