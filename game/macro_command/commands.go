package macro_command

import (
	"fmt"

	"github.com/nika-gromova/o-architecture-patterns/game/base"
)

type MacroCommand struct {
	Commands []base.Command
}

func (c *MacroCommand) Execute() error {
	for _, cmd := range c.Commands {
		if err := cmd.Execute(); err != nil {
			return fmt.Errorf("%w: %w", base.ErrCommandExecution, err)
		}
	}
	return nil
}
