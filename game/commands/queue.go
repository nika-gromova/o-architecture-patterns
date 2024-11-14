package commands

import "github.com/nika-gromova/o-architecture-patterns/game/base"

type QueueCommand struct {
	Queue chan base.Command
	Cmd   base.Command
}

func (h *QueueCommand) Execute() error {
	h.Queue <- h.Cmd
	return nil
}
