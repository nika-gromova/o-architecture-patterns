package errors

import (
	"github.com/nika-gromova/o-architecture-patterns/game/base"
	"github.com/nika-gromova/o-architecture-patterns/game/commands"
)

type LogHandler struct {
	Queue chan base.Command
}

func (h *LogHandler) Handle(cmd base.Command, err error) base.Command {
	retry := &commands.QueueCommand{
		Queue: h.Queue,
		Cmd: &commands.LogCommand{
			Err: err,
			Cmd: cmd,
		},
	}
	return retry
}

type RepeatHandler struct {
	Queue chan base.Command
}

func (h *RepeatHandler) Handle(cmd base.Command, _ error) base.Command {
	retry := &commands.QueueCommand{
		Queue: h.Queue,
		Cmd: &commands.RepeatCommand{
			Cmd: cmd,
		},
	}
	return retry
}

type CounterHandler struct {
	Queue   chan base.Command
	Counter map[string]int
}

func (h *CounterHandler) Handle(cmd base.Command, err error) base.Command {
	defer func() {
		h.Counter[base.GetVarType(cmd)]++
	}()
	if h.Counter[base.GetVarType(cmd)] == 0 {
		return &commands.RepeatCommand{
			Cmd: cmd,
		}
	}
	return &commands.LogCommand{
		Err: err,
		Cmd: cmd,
	}
}

type DoubleHandler struct {
	Queue   chan base.Command
	Counter map[string]int
}

func (h *DoubleHandler) Handle(cmd base.Command, _ error) base.Command {
	defer func() {
		h.Counter[base.GetVarType(cmd)]++
	}()
	if h.Counter[base.GetVarType(cmd)] >= 2 {
		return &commands.RetryCommand{
			Cmd: cmd,
		}
	}
	return &commands.RepeatCommand{
		Cmd: cmd,
	}
}
