package errors

import (
	"fmt"
	"testing"

	"github.com/nika-gromova/o-architecture-patterns/game/base"
	"github.com/nika-gromova/o-architecture-patterns/game/commands"
	"github.com/nika-gromova/o-architecture-patterns/game/move"
	"github.com/nika-gromova/o-architecture-patterns/game/rotate"
	"github.com/nika-gromova/o-architecture-patterns/game/tests/stubs"
	"github.com/stretchr/testify/require"
)

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
	testHandler.Register("ErrorCommand", stubs.TestError.Error(), doubleHandler.Handle)
	testHandler.Register("RetryCommand", stubs.TestError.Error(), logHandler.Handle)

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
		cmd := &stubs.ErrorCommand{}
		err := stubs.TestError
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

func TestHandler_Handle2(t *testing.T) {
	var (
		queue         = make(chan base.Command)
		moveCommand   = &move.MoveCommand{}
		rotateCommand = &rotate.RotateCommand{}
		testHandler   = NewHandler()
	)

	testHandler.Register("MoveCommand", base.ErrGetProperty.Error(), func(command base.Command, err error) base.Command {
		return &commands.QueueCommand{
			Queue: queue,
			Cmd: &commands.LogCommand{
				Err: base.ErrSetProperty,
				Cmd: moveCommand,
			},
		}
	})

	testHandler.Register("MoveCommand", base.ErrSetProperty.Error(), func(command base.Command, _ error) base.Command {
		return &commands.QueueCommand{
			Queue: queue,
			Cmd: &commands.RepeatCommand{
				Cmd: command,
			},
		}
	})

	testHandler.Register("RotateCommand", base.ErrGetProperty.Error(), func(command base.Command, _ error) base.Command {
		return &commands.RepeatCommand{
			Cmd: command,
		}
	})
	testHandler.Register("RepeatCommand", base.ErrGetProperty.Error(), func(command base.Command, err error) base.Command {
		return &commands.LogCommand{
			Cmd: command,
			Err: err,
		}
	})

	testHandler.Register("ErrorCommand", stubs.TestError.Error(), func(command base.Command, _ error) base.Command {
		return &commands.RetryCommand{
			Cmd: command,
		}
	})
	testHandler.Register("RetryCommand", stubs.TestError.Error(), func(command base.Command, _ error) base.Command {
		return &commands.RepeatCommand{
			Cmd: command,
		}
	})
	testHandler.Register("RepeatCommand", stubs.TestError.Error(), func(command base.Command, err error) base.Command {
		return &commands.LogCommand{
			Cmd: command,
			Err: err,
		}
	})

	t.Run("возвращает команду, которая ставит в очередь логирующую команду", func(t *testing.T) {
		cmd := moveCommand
		err := base.ErrGetProperty
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
		err := base.ErrSetProperty
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
			Cmd: expected1,
		}

		got1 := testHandler.Handle(cmd, err)

		require.Equal(t, expected1, got1)

		got2 := testHandler.Handle(got1, err)

		require.Equal(t, expected2, got2)
	})

	t.Run("на первые 2 вызова возвращается повторяющая команда, на третий - ретрай команда", func(t *testing.T) {
		cmd := &stubs.ErrorCommand{}
		err := stubs.TestError
		expected1 := &commands.RetryCommand{
			Cmd: cmd,
		}
		expected2 := &commands.RepeatCommand{
			Cmd: expected1,
		}
		expected3 := &commands.LogCommand{
			Err: err,
			Cmd: expected2,
		}

		got1 := testHandler.Handle(cmd, err)

		require.Equal(t, expected1, got1)

		got2 := testHandler.Handle(got1, err)

		require.Equal(t, expected2, got2)

		got3 := testHandler.Handle(got2, err)

		require.Equal(t, expected3, got3)
	})
}
