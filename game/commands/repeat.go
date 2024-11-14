package commands

import "github.com/nika-gromova/o-architecture-patterns/game/base"

type RepeatCommand struct {
	Cmd base.Command
}

func (c *RepeatCommand) Execute() error {
	return c.Cmd.Execute()
}
