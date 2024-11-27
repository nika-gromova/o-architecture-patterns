package base

import (
	"errors"
)

// Errors
var (
	ErrGetProperty      = errors.New("failed to get property")
	ErrSetProperty      = errors.New("failed to set property")
	ErrCommandExecution = errors.New("error occurred during command execution")
)
