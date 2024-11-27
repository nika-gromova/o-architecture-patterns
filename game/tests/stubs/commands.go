package stubs

import "errors"

var TestError = errors.New("test error")

type ErrorCommand struct{}

func (c *ErrorCommand) Execute() error {
	return TestError
}

type NoErrorCommand struct{}

func (c *NoErrorCommand) Execute() error {
	return nil
}
