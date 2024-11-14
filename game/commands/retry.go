package commands

import "github.com/nika-gromova/o-architecture-patterns/game/base"

type RetryCommand struct {
	Cmd base.Command
}

func (c *RetryCommand) Execute() error {
	return c.Cmd.Execute()
}
