package queue

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/nika-gromova/o-architecture-patterns/game/base"
	"github.com/nika-gromova/o-architecture-patterns/game/commands"
	"github.com/nika-gromova/o-architecture-patterns/game/errors"
	"github.com/nika-gromova/o-architecture-patterns/game/move"
	"github.com/nika-gromova/o-architecture-patterns/game/rotate"
	"github.com/nika-gromova/o-architecture-patterns/game/tests/mocks"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

type ErrorCommand struct{}

func (c *ErrorCommand) Execute() error {
	return fmt.Errorf("err")
}

func TestQueue_Process(t *testing.T) {
	var (
		queue        = make(chan base.Command, 10)
		ctrl         = gomock.NewController(t)
		fakeMoving   = mocks.NewMockMovingObject(ctrl)
		fakeRotating = mocks.NewMockRotatingObject(ctrl)
		moveCommand  = &move.MoveCommand{
			Obj: fakeMoving,
		}
		rotateCommand = &rotate.RotateCommand{
			Obj: fakeRotating,
		}
		testHandler = errors.NewHandler()
	)
	// контекст с отменой, чтобы выйти из цикла
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// моки для возврата ошибок
	fakeMoving.EXPECT().GetLocation().Return(base.Vector{}, false)
	fakeRotating.EXPECT().GetAngle().DoAndReturn(func() (any, any) {
		cancel() // отмена контекста, чтобы выйти после обработки последней команды
		return base.Angle{}, false
	})

	// логирование ошибки
	testHandler.Register("MoveCommand", base.ErrGetProperty.Error(), func(command base.Command, err error) base.Command {
		return &commands.LogCommand{
			Cmd: moveCommand,
			Err: base.ErrGetProperty,
		}
	})
	// ошибка при обработке ошибки
	testHandler.Register("RotateCommand", base.ErrGetProperty.Error(), func(command base.Command, err error) base.Command {
		return &ErrorCommand{}
	})

	// инициализация обработчика очереди
	testQ := New(ctx, testHandler, queue)

	// добавление в очередь команд
	queue <- moveCommand
	queue <- rotateCommand

	// буфер логов, чтобы проверить содержимое после теста
	var buf bytes.Buffer
	logrus.SetOutput(&buf)
	defer func() {
		logrus.SetOutput(os.Stderr)
	}()

	// обработка очереди
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		_ = testQ.Process()
	}()
	wg.Wait()

	// проверка логов
	require.Contains(t, buf.String(), "failed to execute")
	require.Contains(t, buf.String(), "failed to handler error")
}

func TestQueue_ProcessRepeat(t *testing.T) {
	var (
		queue       = make(chan base.Command, 10)
		ctrl        = gomock.NewController(t)
		fakeMoving  = mocks.NewMockMovingObject(ctrl)
		moveCommand = &move.MoveCommand{
			Obj: fakeMoving,
		}
		logHandler = errors.LogHandler{
			Queue: queue,
		}
		testHandler = errors.NewHandler()
	)

	fakeMoving.EXPECT().GetLocation().Return(base.Vector{}, false)

	// добавление в очереди логирующей задачи
	testHandler.Register("MoveCommand", base.ErrGetProperty.Error(), logHandler.Handle)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	testQ := New(ctx, testHandler, queue)
	queue <- moveCommand

	var buf bytes.Buffer
	logrus.SetOutput(&buf)
	defer func() {
		logrus.SetOutput(os.Stderr)
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		_ = testQ.Process()
	}()
	wg.Wait()

	require.Contains(t, buf.String(), "failed to execute")
}
