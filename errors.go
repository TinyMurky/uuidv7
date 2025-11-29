package uuidv7

import "errors"

var (
	// ErrInvalidUUIDFormat represent invalid UUID (xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx) format
	ErrInvalidUUIDFormat = errors.New("invalid UUID format")
)
