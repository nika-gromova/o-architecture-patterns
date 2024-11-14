package errors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/nika-gromova/o-architecture-patterns/game/base"
	"github.com/nika-gromova/o-architecture-patterns/game/commands"
	"github.com/nika-gromova/o-architecture-patterns/game/move"
	"github.com/nika-gromova/o-architecture-patterns/game/rotate"
	"github.com/stretchr/testify/require"
)

var testError = errors.New("test error")

type ErrorCommand struct{}

func (c *ErrorCommand) Execute() error {
	return testError
}

func TestHandler_Handle(t *testing.T) {
	var (
		queue      = make(chan base.Command)
		logHandler = &LogHandler{
			Queue: queue,
		}
		repeatHandler = &RepeatHandler{
			Queue: queue,
		}
		counterHandler = &CounterHandler{
			Queue:   queue,
			Counter: make(map[string]int),
		}
		doubleHandler = &DoubleHandler{
			Queue:   queue,
			Counter: make(map[string]int),
		}
		moveCommand   = &move.MoveCommand{}
		rotateCommand = &rotate.RotateCommand{}
		testHandler   = NewHandler()
	)

	testHandler.Register("MoveCommand", base.ErrGetProperty.Error(), repeatHandler.Handle)
	testHandler.Register("MoveCommand", base.ErrSetProperty.Error(), logHandler.Handle)
	testHandler.Register("RotateCommand", base.ErrGetProperty.Error(), counterHandler.Handle)
	testHandler.Register("ErrorCommand", testError.Error(), doubleHandler.Handle)
	testHandler.Register("RetryCommand", testError.Error(), logHandler.Handle)

	t.Run("возвращает команду, которая ставит в очередь логирующую команду", func(t *testing.T) {
		cmd := moveCommand
		err := base.ErrSetProperty
		expected := &commands.QueueCommand{
			Queue: queue,
			Cmd: &commands.LogCommand{
				Err: base.ErrSetProperty,
				Cmd: moveCommand,
			},
		}

		got := testHandler.Handle(cmd, err)

		require.Equal(t, expected, got)
	})

	t.Run("возвращает команду, которая ставит в очередь повторяющую команду", func(t *testing.T) {
		cmd := moveCommand
		err := base.ErrGetProperty
		expected := &commands.QueueCommand{
			Queue: queue,
			Cmd: &commands.RepeatCommand{
				Cmd: moveCommand,
			},
		}

		got := testHandler.Handle(cmd, err)

		require.Equal(t, expected, got)
	})

	t.Run("возвращает дефолтную команду, если обработчик не был найден", func(t *testing.T) {
		cmd := rotateCommand
		err := fmt.Errorf("unknown error")
		expected := &commands.LogCommand{
			Err: err,
			Cmd: rotateCommand,
		}

		got := testHandler.Handle(cmd, err)

		require.Equal(t, expected, got)
	})

	t.Run("на первый вызов команда добавляется повторяющая команда, на второй - логирующая", func(t *testing.T) {
		cmd := rotateCommand
		err := base.ErrGetProperty
		expected1 := &commands.RepeatCommand{
			Cmd: rotateCommand,
		}
		expected2 := &commands.LogCommand{
			Err: err,
			Cmd: rotateCommand,
		}

		got1 := testHandler.Handle(cmd, err)

		require.Equal(t, expected1, got1)

		got2 := testHandler.Handle(cmd, err)

		require.Equal(t, expected2, got2)
	})

	t.Run("на первые 2 вызова возвращается повторяющая команда, на третий - ретрай команда", func(t *testing.T) {
		cmd := &ErrorCommand{}
		err := testError
		expected1 := &commands.RepeatCommand{
			Cmd: cmd,
		}
		expected2 := &commands.RepeatCommand{
			Cmd: cmd,
		}
		expected3 := &commands.RetryCommand{
			Cmd: cmd,
		}
		expected4 := &commands.QueueCommand{
			Queue: queue,
			Cmd: &commands.LogCommand{
				Err: err,
				Cmd: expected3,
			},
		}

		got1 := testHandler.Handle(cmd, err)

		require.Equal(t, expected1, got1)

		got2 := testHandler.Handle(cmd, err)

		require.Equal(t, expected2, got2)

		got3 := testHandler.Handle(cmd, err)

		require.Equal(t, expected3, got3)

		// вызываем ретрай команду - должны получить логирующую
		got4 := testHandler.Handle(got3, got3.Execute())

		require.Equal(t, expected4, got4)

	})
}
