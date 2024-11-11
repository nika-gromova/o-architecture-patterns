package base

import "fmt"

// Errors
var (
	ErrGetProperty = fmt.Errorf("failed to get property")
	ErrSetProperty = fmt.Errorf("failed to set property")
)
