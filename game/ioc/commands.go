package ioc

import (
	"context"
	"errors"
)

type RegisterCommand struct {
	scope       context.Context
	key         string
	constructor Constructor
}

func (c *RegisterCommand) Execute(_ context.Context) error {
	dependencies, ok := c.scope.Value(ScopeDependenciesKey).(*Dependencies)
	if !ok || dependencies == nil {
		return errors.New("no dependencies found for context")
	}
	(*dependencies)[c.key] = c.constructor
	return nil
}
